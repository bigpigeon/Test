package json_test

import (
	"encoding/json"
	"math"
	"testing"
)

type TestJsonData struct {
	Num json.Number
	I64 int64
}

func TestEncode(t *testing.T) {
	data := TestJsonData{
		Num: "8446744073709551616",
		I64: 8446744073709551616,
	}
	encodeData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(encodeData))
	//decode
	{
		var decodeData TestJsonData
		err := json.Unmarshal(encodeData, &decodeData)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		t.Logf("%#v\n", decodeData)
	}
}

func TestEncodeNil(t *testing.T) {
	data := math.NaN()
	_, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
