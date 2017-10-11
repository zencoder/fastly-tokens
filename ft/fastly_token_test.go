package ft

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("http://token.fastly.com/token"); err != nil {
		t.Error("Error reported when retrieving token from Fastly service", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	token := GenerateToken("Fastly Token Test", 60*time.Second, base64.StdEncoding)

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

	token := GenerateTokenForURL("http://www.example.com/index.html", "0bgZZu4uzL1K2My1842DjuAvkJnE8j9s", time.Now(), base64.StdEncoding)

	if token == string(body) {
		t.Error("Expected token mismatch between Fastly service token and URL-specific token", err)
	}
}

func TestGenerateTokenForWildcardURL(t *testing.T) {
	var urlToSign = "*://www.example.com/*"
	var expiryTime = time.Unix(1507727103, 0)

	token := GenerateTokenForWildcardURL(
		urlToSign,
		"WZmGbDWYGVG2/FyXLYO2dnaRIh4g2pH61k/YdJsk3Bo=",
		expiryTime,
		base64.StdEncoding,
	)

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		t.Errorf("Unexpected error while converting token from Base 64: %s", err.Error())
	}

	tokenParts := strings.Split(string(decodedToken), "_")
	if len(tokenParts) != 3 {
		t.Errorf("Expected 3 parts to token (Expiry / Signature / Signed URL), but got %d", len(tokenParts))
	}

	var expectedTimestamp = "59de16ff"
	if tokenParts[0] != expectedTimestamp {
		t.Errorf("Expiry timestamp was wrong, expected 1507727211 in Hex (%s) but got %s", expectedTimestamp, tokenParts[0])
	}

	var expectedSignature = "409867a8411de63a8659b2ed396910668b158215924374de2658b20bb52b6917"
	if tokenParts[1] != expectedSignature {
		t.Errorf("SHA-256 of secret+expiry was wrong. Expected %s, but got %s", expectedSignature, tokenParts[1])
	}

	b, err := hex.DecodeString(tokenParts[2])
	if err != nil {
		t.Errorf("Unexpected error while converting signed URL from Hex: %s", err.Error())
	}
	if string(b) != urlToSign {
		t.Errorf("Expected token to include signed URL of %q, but got %q", urlToSign, string(b))
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateToken("Fastly Token Test", 60*time.Second, base64.StdEncoding)
	}
}
