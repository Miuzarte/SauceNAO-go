package SauceNao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/Miuzarte/SauceNAO-go/db"

	fs "github.com/Miuzarte/FlareSolverr-go"
	"github.com/go-viper/mapstructure/v2"
)

const (
	API_HOST  = `https://saucenao.com`
	API_PATH  = `/search.php`
	USER_PATH = `/user.php` // 必定触发 cf
)

type Client struct {
	ApiKey             string
	Host               string
	NumRes             int
	Hide               bool
	FlareSolverrClient *fs.Client

	cache struct {
		userAgent string
		cookies   []*http.Cookie
	}
}

type HttpError struct {
	StatusCode int
	Url        string
	Body       string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error %d: %s, %s", e.StatusCode, e.Url, e.Body)
}

func NewClient(apiKey, overrideHost string, numRes int, hide bool, fsClient *fs.Client) *Client {
	host := overrideHost
	if host == "" {
		host = API_HOST
	} else {
		if !strings.HasPrefix(overrideHost, "http") {
			overrideHost = "https://" + overrideHost
		}
		host = strings.TrimRight(overrideHost, "/")
	}
	return &Client{
		ApiKey:             apiKey,
		Host:               host,
		NumRes:             numRes,
		Hide:               hide,
		FlareSolverrClient: fsClient,
	}
}

func (c *Client) Search(ctx context.Context, image any) (resp *Response, err error) {
	switch img := image.(type) {
	case string:
		if strings.HasPrefix(img, "http") {
			return c.Get(ctx, img)
		} else {
			// read local
			f, err := os.Open(img)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			return c.Search(ctx, f)
		}

	case []byte:
		return c.Post(ctx, img)
	case io.Reader:
		imgData, err := io.ReadAll(img)
		if err != nil {
			return nil, err
		}
		return c.Post(ctx, imgData)

	default:
		return nil, fmt.Errorf("unsupported image type: %T", image)
	}
}

func (c *Client) Post(ctx context.Context, imgData []byte) (*Response, error) {
	return c.do(ctx, func() (*http.Request, error) {
		return c.buildPostRequest(ctx, imgData)
	})
}

func (c *Client) Get(ctx context.Context, imgUrl string) (*Response, error) {
	return c.do(ctx, func() (*http.Request, error) {
		return c.buildGetRequest(ctx, imgUrl)
	})
}

func (c *Client) do(ctx context.Context, requestBuilder func() (*http.Request, error)) (*Response, error) {
	const bypassCfRetryTimes = 1
	numRetries := bypassCfRetryTimes
TRYAGAIN:
	req, err := requestBuilder()
	if err != nil {
		return nil, err
	}

	hResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer hResp.Body.Close()

	switch hResp.StatusCode {
	case http.StatusOK:
		break

	case http.StatusForbidden:
		// 尝试过 cf
		_, _, e := c.bypassCf(ctx)
		if e == nil {
			if numRetries > 0 {
				// 成功后重试一次
				numRetries--
				goto TRYAGAIN
			}
		}
		fallthrough
	default:
		const bodyTruncateLen = 1024
		body, _ := io.ReadAll(hResp.Body)
		if len(body) > bodyTruncateLen {
			body = body[:bodyTruncateLen]
		}
		return nil, &HttpError{
			StatusCode: hResp.StatusCode,
			Url:        req.URL.String(),
			Body:       string(body),
		}
	}

	body, err := io.ReadAll(hResp.Body)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	resp.RawBody = string(body)
	err = json.Unmarshal([]byte(body), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// fsGet 完成后缓存 user agent 和 cookies
func (c *Client) fsGet(ctx context.Context, url string) (string, error) {
	if c.FlareSolverrClient == nil {
		return "", fmt.Errorf("FlareSolverrClient is not set")
	}
	resp, err := c.FlareSolverrClient.Get(ctx, url, map[string]any{
		fs.PARAM_MAX_TIMEOUT: 60000,
	})
	if err != nil {
		return "", err
	}
	if resp.Solution.Status != http.StatusOK {
		return "", fmt.Errorf("flaresolverr failed %d: %s, %s",
			resp.Solution.Status, resp.Message, resp.Solution.Response)
	}

	// cache user agent and cookies
	c.cache.userAgent = resp.Solution.UserAgent
	c.cache.cookies = resp.Solution.Cookies.ToHttpCookies()
	return resp.Solution.Response, nil
}

// bypassCf 访问 user.php 获取 cf challenge 凭证
func (c *Client) bypassCf(ctx context.Context) (ua string, cookies []*http.Cookie, err error) {
	userUrl := c.Host + USER_PATH
	_, err = c.fsGet(ctx, userUrl)
	if err != nil {
		return "", nil, err
	}
	return c.cache.userAgent, c.cache.cookies, nil
}

// requestSetHeader 设置请求头
func (c *Client) requestSetHeader(req *http.Request) *http.Request {
	if c.cache.userAgent != "" {
		req.Header.Set("User-Agent", c.cache.userAgent)
	}
	for _, cookie := range c.cache.cookies {
		req.AddCookie(cookie)
	}
	return req
}

func (c *Client) buildPostRequest(ctx context.Context, imgData []byte) (*http.Request, error) {
	buf := bytes.Buffer{}
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", "image")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(imgData)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Host+API_PATH, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	query := req.URL.Query()
	query.Add("api_key", c.ApiKey)
	query.Add("output_type", "2")
	if c.NumRes > 0 {
		query.Add("numres", strconv.Itoa(c.NumRes))
	}
	query.Add("hide", strconv.FormatBool(c.Hide))
	req.URL.RawQuery = query.Encode()

	return c.requestSetHeader(req), nil
}

func (c *Client) buildGetRequest(ctx context.Context, imgUrl string) (*http.Request, error) {
	u, err := url.Parse(c.Host + API_PATH)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("url", imgUrl)
	query.Add("api_key", c.ApiKey)
	query.Add("output_type", "2")
	query.Add("numres", strconv.Itoa(c.NumRes))
	query.Add("hide", strconv.FormatBool(c.Hide))
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.requestSetHeader(req), nil
}

type Response struct {
	Header  ResponseHeader `json:"header"`
	Results []Result       `json:"results"`
	RawBody string         `json:"-"` // for debug
}

type ResponseHeader struct {
	UserId            string  `json:"user_id"`             // api key 所属
	AccountType       string  `json:"account_type"`        // "1"
	ShortLimit        string  `json:"short_limit"`         // "4"   // 30s
	LongLimit         string  `json:"long_limit"`          // "100" // 24h
	LongRemaining     int     `json:"long_remaining"`      // 85
	ShortRemaining    int     `json:"short_remaining"`     // 3
	Status            int     `json:"status"`              // 0
	ResultsRequested  int     `json:"results_requested"`   // 对应 numres 参数
	SearchDepth       string  `json:"search_depth"`        // "128"
	MinimumSimilarity float64 `json:"minimum_similarity"`  // 41.49
	QueryImageDisplay string  `json:"query_image_display"` // 提交的图片 url
	QueryImage        string  `json:"query_image"`         // 提交的图片文件名
	ResultsReturned   int     `json:"results_returned"`    // 最终返回的结果数量

	Index map[string]IndexEntry `json:"index"` // "0","2","44" ...
}

type IndexEntry struct {
	Status   int `json:"status"`
	ParentID int `json:"parent_id"`
	ID       int `json:"id"`
	Results  int `json:"results"`
}

type Result struct {
	Header ResultHeader   `json:"header"`
	Data   map[string]any `json:"data"` // delay decode using [mapstructure.Decode]
}

type ResultHeader struct {
	Similarity string     `json:"similarity"` // .2f
	Thumbnail  string     `json:"thumbnail"`
	IndexId    db.IndexId `json:"index_id"`
	IndexName  string     `json:"index_name"`
	Dupes      int        `json:"dupes"`
	Hidden     int        `json:"hidden"`
}

func decodeTo[T any](input map[string]any) *T {
	output := new(T)

	config := mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           output,
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	return output
}

func (r Result) DecodeData() (ret interface{ String() string }) {
	defer func() {
		if rec := recover(); rec != nil {
			ret = &db.ResultDataUnknown{
				Raw: r.Data,
				Err: fmt.Errorf("failed to decode data: %v", rec),
			}
		}
	}()
	data := r.Data
	switch r.Header.IndexId {
	case db.HMAGAZINES:
		return decodeTo[db.ResultDataHMagazines](data)
	case db.HGAMECG:
		return decodeTo[db.ResultDataHGameCg](data)
	case db.DOUJINSHIDB:
		return decodeTo[db.ResultDataDoujinshiDb](data)
	case db.PIXIV:
		return decodeTo[db.ResultDataPixiv](data)
	case db.SEIGA:
		return decodeTo[db.ResultDataSeiga](data)
	case db.DANBOORU:
		return decodeTo[db.ResultDataDanbooru](data)
	case db.DRAWR:
		return decodeTo[db.ResultDataDrawr](data)
	case db.NIJIE:
		return decodeTo[db.ResultDataNijie](data)
	case db.YANDERE:
		return decodeTo[db.ResultDataYandere](data)
	case db.SHUTTERSTOCK:
		return decodeTo[db.ResultDataShutterstock](data)
	case db.FAKKU:
		return decodeTo[db.ResultDataFakku](data)
	case db.NHENTAI:
		return decodeTo[db.ResultDataNHentai](data)
	case db.MARKET2D:
		return decodeTo[db.ResultDataMarket2d](data)
	case db.MEDIBANG:
		return decodeTo[db.ResultDataMediBang](data)
	case db.ANIME:
		return decodeTo[db.ResultDataAnime](data)
	case db.HANIME:
		return decodeTo[db.ResultDataHAnime](data)
	case db.MOVIES:
		return decodeTo[db.ResultDataMovies](data)
	case db.SHOWS:
		return decodeTo[db.ResultDataShows](data)
	case db.GELBOORU:
		return decodeTo[db.ResultDataGelbooru](data)
	case db.KONACHAN:
		return decodeTo[db.ResultDataKonachan](data)
	case db.SANKAKU:
		return decodeTo[db.ResultDataSankaku](data)
	case db.ANIMEPICTURES:
		return decodeTo[db.ResultDataAnimePictures](data)
	case db.E621:
		return decodeTo[db.ResultDataE621](data)
	case db.IDOLCOMPLEX:
		return decodeTo[db.ResultDataIdolComplex](data)
	case db.BCY_ILLUST:
		return decodeTo[db.ResultDataBcyIllust](data)
	case db.BCY_COSPLAY:
		return decodeTo[db.ResultDataBcyCosplay](data)
	case db.PORTALGRAPHICS:
		return decodeTo[db.ResultDataPortalGraphics](data)
	case db.DEVIANTART:
		return decodeTo[db.ResultDataDeviantArt](data)
	case db.PAWOO:
		return decodeTo[db.ResultDataPawoo](data)
	case db.MADOKAMI:
		return decodeTo[db.ResultDataMadokami](data)
	case db.MANGADEX:
		return decodeTo[db.ResultDataMangaDex](data)
	case db.EHENTAI:
		return decodeTo[db.ResultDataEHentai](data)
	case db.ARTSTATION:
		return decodeTo[db.ResultDataArtStation](data)
	case db.FURAFFINITY:
		return decodeTo[db.ResultDataFurAffinity](data)
	case db.TWITTER:
		return decodeTo[db.ResultDataTwitter](data)
	case db.FURRYNETWORK:
		return decodeTo[db.ResultDataFurryNetwork](data)
	case db.KEMONO:
		return decodeTo[db.ResultDataKemono](data)
	case db.SKEB:
		return decodeTo[db.ResultDataSkeb](data)
	default:
		return &db.ResultDataUnknown{Raw: data}
	}
}
