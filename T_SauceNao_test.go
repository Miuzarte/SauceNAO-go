package SauceNao

import (
	"io"
	"net/http"
	"testing"

	fs "github.com/Miuzarte/FlareSolverr-go"
)

const testUrl = `https://saucenao.com/userdata/tmp/WJFAtBj5r.png`

var testFsClient = fs.NewClient("http://127.0.0.1:8191/v1")

func TestPost(t *testing.T) {
	imgResp, err := http.Get(testUrl)
	if err != nil {
		t.Fatal(err)
	}
	defer imgResp.Body.Close()
	if imgResp.StatusCode != http.StatusOK {
		t.Fatalf("failed to download test image: %s", imgResp.Status)
	}
	imgData, err := io.ReadAll(imgResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient("", "", 0, false, testFsClient)
	resp, err := client.Post(t.Context(), imgData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", *resp)
}

func TestGet(t *testing.T) {
	client := NewClient("", "", 0, false, testFsClient)
	resp, err := client.Get(t.Context(), testUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", *resp)
}
