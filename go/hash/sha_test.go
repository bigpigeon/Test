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

func TestConflict(t *testing.T) {
	d1 := sha1.Sum([]byte("0"))
	d2 := sha1.Sum([]byte("100"))
	t.Log(base64.StdEncoding.EncodeToString(d1[:]))
	t.Log(base64.StdEncoding.EncodeToString(d2[:]))
}
