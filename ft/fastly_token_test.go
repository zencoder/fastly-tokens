package ft

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("http://token.fastly.com/token"); err != nil {
		t.Error("Error reported when retrieving token from Fastly service", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	token, err := GenerateToken("Fastly Token Test", 60, base64.StdEncoding)
	if err != nil {
		t.Error("Error while generating token", err)
	}

	if token != string(body) {
		t.Errorf("Expected token: %s, Actual Token: %s", body, token)
	}
}

func TestGenerateTokenForURL(t *testing.T) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("http://token.fastly.com/token"); err != nil {
		t.Error("Error reported when retrieving token from Fastly service", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	token, err := GenerateTokenForURL("http://www.example.com/index.html", "0bgZZu4uzL1K2My1842DjuAvkJnE8j9s", 60, base64.StdEncoding)
	if err != nil {
		t.Error("Error while generating token", err)
	}

	if token == string(body) {
		t.Error("Expected token mismatch between Fastly service token and URL-specific token", err)
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateToken("Fastly Token Test", 60, base64.StdEncoding)
	}
}
