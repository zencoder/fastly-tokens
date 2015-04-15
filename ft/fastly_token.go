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

func GenerateToken(secret string, secondsValid int, encoding *base64.Encoding) (token string, err error) {
	return GenerateTokenForURL("", secret, secondsValid, encoding)
}

func GenerateTokenForURL(url string, secret string, secondsValid int, encoding *base64.Encoding) (token string, err error) {
	var secretBytes []byte
	if secretBytes, err = base64.StdEncoding.DecodeString(secret); err != nil {
		return "", err
	}

	var secretBuf bytes.Buffer
	secretBuf.Write(secretBytes)
	secretBuf.WriteString(url)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint64(time.Now().Unix()/int64(secondsValid)))
	mac := hmac.New(sha256.New, secretBuf.Bytes())
	mac.Write(buf.Bytes())
	token = strings.TrimSpace(encoding.EncodeToString(mac.Sum(nil)))

	return
}
