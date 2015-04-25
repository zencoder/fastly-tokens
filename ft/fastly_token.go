package ft

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func GenerateToken(secret string, secondsValid int, encoding *base64.Encoding) (token string, err error) {
	secretBytes := []byte(secret)

	var secretBuf bytes.Buffer
	secretBuf.Write(secretBytes)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint64(time.Now().Unix()/int64(secondsValid)))
	mac := hmac.New(sha256.New, secretBuf.Bytes())
	mac.Write(buf.Bytes())
	token = strings.TrimSpace(encoding.EncodeToString(mac.Sum(nil)))

	return
}

/* VCL for verifying URL specific tokens

# Header rewrite Token : 10
set req.http.X-Token = digest.base64_decode(urldecode(regsub(req.url, ".*token=([^&]+)(?:&|$).*", "\1")));

# Header rewrite Token Expiry : 11
set req.http.X-Token-Expiry = regsub(req.http.X-Token, "^([^_]+)_.*", "\1");

# Header rewrite Token Signature : 11
set req.http.X-Token-Signature = regsub(req.http.X-Token, "^[^_]+_(.*)", "\1");

# Header rewrite Expected Signature : 12
set req.http.X-Expected-Sig = regsub(digest.hmac_sha256("0bgZZu4uzL1K2My1842DjuAvkJnE8j9s", req.url.path req.http.X-Token-Expiry), "^0x", "");


# Request Condition: Token Auth Prio: 10
if( !req.http.Fastly-FF && !((req.http.X-Expected-Sig == req.http.X-Token-Signature) && time.is_after(time.hex_to_time(1, req.http.X-Token-Expiry), now)) ) {
	# ResponseObject: Token Authentication
	error 900 "Fastly Internal";
}

*/
func GenerateTokenForURL(filename string, secret string, expiration int, encoding *base64.Encoding) (token string, err error) {
	data := fmt.Sprintf("%s%x", filename, expiration)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	digest := mac.Sum(nil)

	hexDigest := hex.EncodeToString(digest)
	token = url.QueryEscape(strings.TrimSpace(encoding.EncodeToString([]byte(fmt.Sprintf("%x_%s", expiration, hexDigest)))))
	return
}
