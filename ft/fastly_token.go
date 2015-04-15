package ft

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"strings"
	"time"
)

func GenerateToken(secret string, seconds_valid int, encoding *base64.Encoding) (token string, err error) {
	return GenerateTokenForURL("", secret, seconds_valid, encoding)
}

func GenerateTokenForURL(url string, secret string, seconds_valid int, encoding *base64.Encoding) (token string, err error) {
	var secret_bytes []byte
	if secret_bytes, err = base64.StdEncoding.DecodeString(secret); err != nil {
		return "", err
	}

	var secret_buffer bytes.Buffer
	secret_buffer.Write(secret_bytes)
	secret_buffer.WriteString(url)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint64(time.Now().Unix()/int64(seconds_valid)))
	mac := hmac.New(sha256.New, secret_buffer.Bytes())
	mac.Write(buf.Bytes())
	token = strings.TrimSpace(encoding.EncodeToString(mac.Sum(nil)))

	return
}
