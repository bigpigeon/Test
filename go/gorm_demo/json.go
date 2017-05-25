package gormdemo

import (
	"encoding/json"
)

func Encode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
