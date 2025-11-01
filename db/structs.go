package db

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"
	"time"
)

func toJsonString(v any, indent string) string {
	j, err := json.MarshalIndent(v, "", indent)
	if err != nil {
		return err.Error()
	}
	return string(j)
}

// 0 H-Magazines
type ResultDataHMagazines struct {
	Todo struct{}
}

func (rd ResultDataHMagazines) String() string            { return "[TODO]" }
func (rd ResultDataHMagazines) Json(indent string) string { return toJsonString(rd, indent) }

// 2 H-Game CG
type ResultDataHGameCg struct {
	Title    string `json:"title"`
	Company  string `json:"company"`
	GetchuId string `json:"getchu_id"`
}

func (rd ResultDataHGameCg) String() string {
	return fmt.Sprintf(
		`%s
Company: %s
GetchuId: %s`,
		rd.Title,
		rd.Company,
		rd.GetchuId,
	)
}
func (rd ResultDataHGameCg) Json(indent string) string { return toJsonString(rd, indent) }

// 3 DoujinshiDB
type ResultDataDoujinshiDb struct {
	Todo struct{}
}

func (rd ResultDataDoujinshiDb) String() string            { return "[TODO]" }
func (rd ResultDataDoujinshiDb) Json(indent string) string { return toJsonString(rd, indent) }

// 5 pixiv Images
type ResultDataPixiv struct {
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	PixivId    int      `json:"pixiv_id"`
	MemberName string   `json:"member_name"`
	MemberId   int      `json:"member_id"`
}

func (rd ResultDataPixiv) String() string {
	return fmt.Sprintf(
		`%s
pixiv.net/i/%d
%s: pixiv.net/u/%d`,
		rd.Title,
		rd.PixivId,
		rd.MemberName, rd.MemberId,
	)
}
func (rd ResultDataPixiv) Json(indent string) string { return toJsonString(rd, indent) }

// 8 Nico Nico Seiga
type ResultDataSeiga struct {
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	SeigaId    int      `json:"seiga_id"`
	MemberName string   `json:"member_name"`
	MemberId   int      `json:"member_id"`
}

func (rd ResultDataSeiga) String() string {
	return fmt.Sprintf(
		`%s
seiga.nicovideo.jp/seiga/im%d
MemberName: %s
MemberId: %d`,
		rd.Title,
		rd.SeigaId,
		rd.MemberName,
		rd.MemberId,
	)
}
func (rd ResultDataSeiga) Json(indent string) string { return toJsonString(rd, indent) }

// 9 Danbooru
type ResultDataDanbooru struct {
	ExtUrls    []string `json:"ext_urls"`
	DanbooruId int      `json:"danbooru_id"` // "https://danbooru.donmai.us/posts/{.DanbooruId}"
	GelbooruId int      `json:"gelbooru_id"` // "https://gelbooru.com/index.php?page=post&s=view&id={.GelbooruId}"
	Creator    string   `json:"creator"`     // 作者 // "earosoligt"
	Material   string   `json:"material"`    // 作品 // "blue archive"
	Characters string   `json:"characters"`  // 角色 // "miyako (blue archive)"
	Source     string   `json:"source"`      // url // twitter | pixiv | lofter
}

func (rd ResultDataDanbooru) String() string {
	return fmt.Sprintf(
		`%s
%s
%s
danbooru.donmai.us/posts/%d
gelbooru.com/index.php?page=post&s=view&id=%d
%s`,
		rd.Characters,
		rd.Material,
		rd.Creator,
		rd.DanbooruId,
		rd.GelbooruId,
		rd.Source,
	)
}
func (rd ResultDataDanbooru) Json(indent string) string { return toJsonString(rd, indent) }

// 10 drawr Images
type ResultDataDrawr struct {
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	DrawrId    int      `json:"drawr_id"`
	MemberName string   `json:"member_name"`
	MemberId   int      `json:"member_id"`
}

func (rd ResultDataDrawr) String() string {
	return fmt.Sprintf(
		`%s
drawr.net/show.php?id=%d
MemberName: %s
MemberId: %d`,
		rd.Title,
		rd.DrawrId,
		rd.MemberName,
		rd.MemberId,
	)
}
func (rd ResultDataDrawr) Json(indent string) string { return toJsonString(rd, indent) }

// 11 Nijie Images
type ResultDataNijie struct {
	Todo struct{}
}

func (rd ResultDataNijie) String() string            { return "[TODO]" }
func (rd ResultDataNijie) Json(indent string) string { return toJsonString(rd, indent) }

// 12 Yande.re
type ResultDataYandere struct {
	ExtUrls    []string `json:"ext_urls"`
	YandereId  int      `json:"yandere_id"` // "https://yande.re/post/show/{.YandereId}"
	Creator    string   `json:"creator"`    // 作者 // "momoko (momopoco)"
	Material   string   `json:"material"`   // 作品 // "tokidoki bosotto roshia-go de dereru tonari no arya-san"
	Characters string   `json:"characters"` // 角色 // "alisa nikolaevna kujou"
	Source     string   `json:"source"`     // url // twitter | pixiv | lofter
}

func (rd ResultDataYandere) String() string {
	return fmt.Sprintf(
		`%s
%s
%s
yande.re/post/show/%d
%s`,
		rd.Characters,
		rd.Material,
		rd.Creator,
		rd.YandereId,
		rd.Source,
	)
}
func (rd ResultDataYandere) Json(indent string) string { return toJsonString(rd, indent) }

// 15 Shutterstock
type ResultDataShutterstock struct {
	Todo struct{}
}

func (rd ResultDataShutterstock) String() string            { return "[TODO]" }
func (rd ResultDataShutterstock) Json(indent string) string { return toJsonString(rd, indent) }

// 16 FAKKU
type ResultDataFakku struct {
	// ALL
	ExtUrls []string `json:"ext_urls"`
	Source  string   `json:"source"`
	Creator string   `json:"creator"`
}

func (rd ResultDataFakku) String() string {
	return fmt.Sprintf(
		`Source: %s
Creator: %s
%s`,
		rd.Source,
		rd.Creator,
		strings.Join(rd.ExtUrls, "\n"),
	)
}
func (rd ResultDataFakku) Json(indent string) string { return toJsonString(rd, indent) }

// 18|38
type ResultDataDoujin struct {
	Source  string   `json:"source"`
	Creator []string `json:"creator"`
	EngName string   `json:"eng_name"`
	JpName  string   `json:"jp_name"`
}

func (rd ResultDataDoujin) String() string {
	name := rd.JpName
	if name == "" {
		name = rd.EngName
	}
	return fmt.Sprintf(
		`%s
%s
Creator: %s`,
		rd.Source,
		name,
		strings.Join(rd.Creator, ", "),
	)
}
func (rd ResultDataDoujin) Json(indent string) string { return toJsonString(rd, indent) }

// 18 H-Misc (nH)
type ResultDataNHentai = ResultDataDoujin

// 19 2D-Market
type ResultDataMarket2d struct {
	Todo struct{}
}

func (rd ResultDataMarket2d) String() string            { return "[TODO]" }
func (rd ResultDataMarket2d) Json(indent string) string { return toJsonString(rd, indent) }

// 20 MediBang
type ResultDataMediBang struct {
	Todo struct{}
}

func (rd ResultDataMediBang) String() string            { return "[TODO]" }
func (rd ResultDataMediBang) Json(indent string) string { return toJsonString(rd, indent) }

// 21 Anime
type ResultDataAnime struct {
	ExtUrls   []string `json:"ext_urls"`
	Source    string   `json:"source"`     // 作品
	AnidbAid  int      `json:"anidb_aid"`  // "https://anidb.net/anime/{.AnidbAid}"
	AnilistId int      `json:"anilist_id"` // "https://anilist.co/anime/{.AnilistId}"
	MalId     int      `json:"mal_id"`     // "https://myanimelist.net/anime/{.MalId}"
	Part      string   `json:"part"`
	Year      string   `json:"year"`
	EstTime   string   `json:"est_time"`
}

func (rd ResultDataAnime) String() string {
	return fmt.Sprintf(
		`%s
anidb.net/anime/%d
anilist.co/anime/%d
myanimelist.net/anime/%d
Part: %s  Year: %s  Est: %s`,
		rd.Source,
		rd.AnidbAid,
		rd.AnilistId,
		rd.MalId,
		rd.Part, rd.Year, rd.EstTime,
	)
}
func (rd ResultDataAnime) Json(indent string) string { return toJsonString(rd, indent) }

// 22 H-Anime
type ResultDataHAnime struct {
	Todo struct{}
}

func (rd ResultDataHAnime) String() string            { return "[TODO]" }
func (rd ResultDataHAnime) Json(indent string) string { return toJsonString(rd, indent) }

// 23 Movies
type ResultDataMovies struct {
	ExtUrls []string `json:"ext_urls"`
	Source  string   `json:"source"`
	ImdbId  string   `json:"imdb_id"` // "https://www.imdb.com/title/{.ImdbId}"
	Part    string   `json:"part"`
	Year    string   `json:"year"`
	EstTime string   `json:"est_time"`
}

func (rd ResultDataMovies) String() string {
	return fmt.Sprintf(
		`%s
imdb.com/title/%s
Part: %s  Year: %s  Est: %s`,
		rd.Source,
		rd.ImdbId,
		rd.Part,
		rd.Year,
		rd.EstTime,
	)
}
func (rd ResultDataMovies) Json(indent string) string { return toJsonString(rd, indent) }

// 24 Shows
type ResultDataShows struct {
	Todo struct{}
}

func (rd ResultDataShows) String() string            { return "[TODO]" }
func (rd ResultDataShows) Json(indent string) string { return toJsonString(rd, indent) }

// 25 Gelbooru
type ResultDataGelbooru struct {
	ExtUrls    []string `json:"ext_urls"`
	GelbooruId int      `json:"gelbooru_id"`
	Creator    string   `json:"creator"`    // ""
	Material   string   `json:"material"`   // ""
	Characters string   `json:"characters"` // ""
	Source     string   `json:"source"`
}

func (rd ResultDataGelbooru) String() string {
	return fmt.Sprintf(
		`%s
%s
%s
gelbooru.com/index.php?page=post&s=view&id=%d
%s`,
		rd.Characters,
		rd.Material,
		rd.Creator,
		rd.GelbooruId,
		rd.Source,
	)
}
func (rd ResultDataGelbooru) Json(indent string) string { return toJsonString(rd, indent) }

// 26 Konachan
type ResultDataKonachan struct {
	Todo struct{}
}

func (rd ResultDataKonachan) String() string            { return "[TODO]" }
func (rd ResultDataKonachan) Json(indent string) string { return toJsonString(rd, indent) }

// 27 Sankaku Channel
type ResultDataSankaku struct {
	Todo struct{}
}

func (rd ResultDataSankaku) String() string            { return "[TODO]" }
func (rd ResultDataSankaku) Json(indent string) string { return toJsonString(rd, indent) }

// 28 Anime-Pictures.net
type ResultDataAnimePictures struct {
	Todo struct{}
}

func (rd ResultDataAnimePictures) String() string            { return "[TODO]" }
func (rd ResultDataAnimePictures) Json(indent string) string { return toJsonString(rd, indent) }

// 29 e621.net
type ResultDataE621 struct {
	Todo struct{}
}

func (rd ResultDataE621) String() string            { return "[TODO]" }
func (rd ResultDataE621) Json(indent string) string { return toJsonString(rd, indent) }

// 30 Idol Complex
type ResultDataIdolComplex struct {
	Todo struct{}
}

func (rd ResultDataIdolComplex) String() string            { return "[TODO]" }
func (rd ResultDataIdolComplex) Json(indent string) string { return toJsonString(rd, indent) }

// 31 bcy.net Illust
type ResultDataBcyIllust struct {
	ExtUrls      []string `json:"ext_urls"`
	Title        string   `json:"title"`
	BcyId        int      `json:"bcy_id"`
	MemberName   string   `json:"member_name"`
	MemberId     int      `json:"member_id"`
	MemberLinkId int      `json:"member_link_id"` // "https://bcy.net/illust/detail/{.MemberLinkId}"
	BcyType      string   `json:"bcy_type"`       // "illust"
}

func (rd ResultDataBcyIllust) String() string {
	return fmt.Sprintf(
		`%s
bcy.net/illust/detail/%d
MemberName: %s
MemberId: %d`,
		rd.Title,
		rd.MemberLinkId,
		rd.MemberName,
		rd.MemberId,
	)
}
func (rd ResultDataBcyIllust) Json(indent string) string { return toJsonString(rd, indent) }

// 32 bcy.net Cosplay
type ResultDataBcyCosplay struct {
	Todo struct{}
}

func (rd ResultDataBcyCosplay) String() string            { return "[TODO]" }
func (rd ResultDataBcyCosplay) Json(indent string) string { return toJsonString(rd, indent) }

// 33 PortalGraphics.net
type ResultDataPortalGraphics struct {
	Todo struct{}
}

func (rd ResultDataPortalGraphics) String() string            { return "[TODO]" }
func (rd ResultDataPortalGraphics) Json(indent string) string { return toJsonString(rd, indent) }

// 34 deviantArt
type ResultDataDeviantArt struct {
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	DaId       string   `json:"da_id"`
	AuthorName string   `json:"author_name"`
	AuthorUrl  string   `json:"author_url"`
}

func (rd ResultDataDeviantArt) String() string {
	return fmt.Sprintf(
		`%s
deviantart.com/view/%s
%s: %s`,
		rd.Title,
		rd.DaId,
		rd.AuthorName, rd.AuthorUrl,
	)
}
func (rd ResultDataDeviantArt) Json(indent string) string { return toJsonString(rd, indent) }

// 35 Pawoo.net
type ResultDataPawoo struct {
	Todo struct{}
}

func (rd ResultDataPawoo) String() string            { return "[TODO]" }
func (rd ResultDataPawoo) Json(indent string) string { return toJsonString(rd, indent) }

// 36 Madokami (Manga)
type ResultDataMadokami struct {
	Todo struct{}
}

func (rd ResultDataMadokami) String() string            { return "[TODO]" }
func (rd ResultDataMadokami) Json(indent string) string { return toJsonString(rd, indent) }

// 37 MangaDex
type ResultDataMangaDex struct {
	ExtUrls []string `json:"ext_urls"`
	Source  string   `json:"source"` // 作品
	MdId    string   `json:"md_id"`  // "https://mangadex.org/chapter/{.MdId}"
	MuId    int      `json:"mu_id"`  // "https://www.mangaupdates.com/series.html?id={.MuId}"
	MalId   int      `json:"mal_id"` // "https://myanimelist.net/manga/{.MalId}"
	Part    string   `json:"part"`
	Artist  string   `json:"artist"`
	Author  string   `json:"author"`
}

func (rd ResultDataMangaDex) String() string {
	return fmt.Sprintf(
		`%s%s
mangadex.org/chapter/%s
www.mangaupdates.com/series.html?id=%d
myanimelist.net/manga/%d
Artist: %s
Author: %s`,
		rd.Source, rd.Part,
		rd.MdId,
		rd.MuId,
		rd.MalId,
		rd.Artist,
		rd.Author,
	)
}
func (rd ResultDataMangaDex) Json(indent string) string { return toJsonString(rd, indent) }

// 38 H-Misc (eH)
type ResultDataEHentai = ResultDataDoujin

// 39 ArtStation
type ResultDataArtStation struct {
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	AsProject  string   `json:"as_project"`
	AuthorName string   `json:"author_name"`
	AuthorUrl  string   `json:"author_url"`
}

func (rd ResultDataArtStation) String() string {
	return fmt.Sprintf(
		`%s
artstation.com/artwork/%s
%s: %s`,
		rd.Title,
		rd.AsProject,
		rd.AuthorName, rd.AuthorUrl,
	)
}
func (rd ResultDataArtStation) Json(indent string) string { return toJsonString(rd, indent) }

// 40 FurAffinity
type ResultDataFurAffinity struct {
	Todo struct{}
}

func (rd ResultDataFurAffinity) String() string            { return "[TODO]" }
func (rd ResultDataFurAffinity) Json(indent string) string { return toJsonString(rd, indent) }

// 41 Twitter
type ResultDataTwitter struct {
	ExtUrls           []string `json:"ext_urls"`
	CreatedAt         string   `json:"created_at"` // "2019-07-18T16:09:17Z"
	TweetId           string   `json:"tweet_id"`   // https://x.com/i/web/status/{.TweetId}
	TwitterUserId     string   `json:"twitter_user_id"`
	TwitterUserHandle string   `json:"twitter_user_handle"` // https://x.com/{.TwitterUserHandle}
}

func (rd ResultDataTwitter) String() string {
	// 重新格式化时间
	formatTime := func(s string) string {
		const layout = "2006/01/02 15:04:05"
		if s == "" {
			return s
		}
		if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
			return t.In(time.Local).Format(layout)
		}
		if t, err := time.Parse(time.RFC3339, s); err == nil {
			return t.In(time.Local).Format(layout)
		}
		return s
	}
	return fmt.Sprintf(
		`%s
x.com/%s/status/%s
x.com/intent/user?user_id=%s`,
		formatTime(rd.CreatedAt),
		rd.TwitterUserHandle,
		rd.TweetId,
		rd.TwitterUserId,
	)
}
func (rd ResultDataTwitter) Json(indent string) string { return toJsonString(rd, indent) }

// 42 Furry Network
type ResultDataFurryNetwork struct {
	Todo struct{}
}

func (rd ResultDataFurryNetwork) String() string            { return "[TODO]" }
func (rd ResultDataFurryNetwork) Json(indent string) string { return toJsonString(rd, indent) }

// 43 Kemono
type ResultDataKemono struct {
	ExtUrls     []string `json:"ext_urls"`
	Published   string   `json:"published"` // "2020-09-25T01:34:37.000Z"
	Title       string   `json:"title"`
	Service     string   `json:"service"`      // "fanbox"
	ServiceName string   `json:"service_name"` // "pixiv FANBOX"
	Id          string   `json:"id"`
	UserId      string   `json:"user_id"` // "https://www.pixiv.net/fanbox/creator/{.UserId}/post/{.Id}"
	UserName    string   `json:"user_name"`
}

func (rd ResultDataKemono) String() string {
	return fmt.Sprintf(
		`%s
pixiv.net/fanbox/creator/%s/post/%s
%s: pixiv.net/fanbox/creator/%s`,
		rd.Title,
		rd.UserId, rd.Id,
		rd.UserName, rd.UserId,
	)
}
func (rd ResultDataKemono) Json(indent string) string { return toJsonString(rd, indent) }

// 44 Skeb
type ResultDataSkeb struct {
	ExtUrls     []string `json:"ext_urls"`
	Path        string   `json:"path"`         // "/@neko_satsuma/works/21"
	Creator     string   `json:"creator"`      // "@neko_satsuma"
	CreatorName string   `json:"creator_name"` // "\u306d\u3053\u3055\u3064\u307e"
	AuthorName  string   `json:"author_name"`  // null
	AuthorUrl   string   `json:"author_url"`   // "https://skeb.jp/@neko_satsuma"
}

func (rd ResultDataSkeb) String() string {
	return fmt.Sprintf(
		`skeb.jp%s
%s: skeb.jp/%s`,
		rd.Path,
		rd.CreatorName, rd.Creator,
	)
}
func (rd ResultDataSkeb) Json(indent string) string { return toJsonString(rd, indent) }

type ResultDataUnknown struct {
	Raw map[string]any
	Err error
}

func (rd ResultDataUnknown) String() string {
	if rd.Err != nil {
		rd.Raw = maps.Clone(rd.Raw) // 可能没必要
		rd.Raw["_decode_error"] = rd.Err.Error()
	}
	return rd.Json("  ")
}
func (rd ResultDataUnknown) Json(indent string) string { return toJsonString(rd.Raw, indent) }
