package hash

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"testing"
)

func TestSha1(t *testing.T) {
	mac := hmac.New(sha1.New, []byte("abcd"))
	_, err := mac.Write([]byte("abcd"))
	if err != nil {
		panic(err)
	}

	t.Log(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}
