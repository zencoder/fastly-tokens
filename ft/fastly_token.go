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
	var secret_bytes []byte
	if secret_bytes, err = base64.StdEncoding.DecodeString(secret); err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint64(time.Now().Unix()/int64(seconds_valid)))
	mac := hmac.New(sha256.New, secret_bytes)
	mac.Write(buf.Bytes())
	token = strings.TrimSpace(encoding.EncodeToString(mac.Sum(nil)))

	return
}
