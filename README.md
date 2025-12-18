# SauceNAO-go

## 要求

- Go 1.25+
- 运行中的 FlareSolverr 服务端

## 安装

```bash
go get github.com/Miuzarte/SauceNAO-go
```

## 示例

```go
package main

import (
    "context"
    "fmt"
    "reflect"
    "strconv"
    fs "github.com/Miuzarte/FlareSolverr-go"
    sn "github.com/Miuzarte/SauceNAO-go"
)

const (
    // https://saucenao.com/user.php
    API_KEY = "abc123"
    // using `https://saucenao.com` when empty
    OVERRIDE_HOST = ""
    // number of results
    NUM_RESULT = 2
    // 0=show all
    // 1=hide expected explicit
    // 2=hide expected and suspected explicit
    // 3=hide all but expected safe
    HIDE = 0
)

func main() {
    fsClient := fs.NewClient("http://127.0.0.1:8191/v1")
    client := sn.NewClient(
        API_KEY,
        OVERRIDE_HOST,
        NUM_RESULT,
        HIDE,
        fsClient,
    )

    // can be:
    // url.(string): using .Get()
    // filepath.(string), data.([]byte), r.(io.Reader): using .Post()
    var img any
    resp, err := client.Search(context.Background(), img)
    if err != nil {
        panic(err)
    }

    for _, result := range resp.Results {
        similarity, _ := strconv.ParseFloat(header.Similarity, 64)

        hideThumbnail := similarity < 60.0 && result.Header.Hidden > 0
        if !hideThumbnail {
            fmt.Printf("thumbnail: %s", result.Header.Thumbnail)
        }

        var data interface{String() string} = result.DecodeData()
        v := reflect.ValueOf(data)
        if v.Kind() == reflect.Pointer && v.IsNil() {
            // the corresponding struct is not completed
            fmt.Printf("%s: %s\n", header.IndexName, resp.RawBody)
            continue
        }

        fmt.Printf(
            "SauceNAO (%s%%)\n%s\n%s\n",
            result.Header.Similarity,
            result.Header.IndexName,
            data,
        )
    }
}
```
