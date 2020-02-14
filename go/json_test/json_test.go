package json_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

type TestJsonData struct {
	Num json.Number
	I64 int64
}

type TestEmbedMap map[string]string

type TestEmbedMapData struct {
	TestEmbedMap
	Data string
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

func TestEmbedMapEncode(t *testing.T) {
	data := TestEmbedMapData{
		TestEmbedMap: TestEmbedMap{"test": "22"},
		Data:         "123",
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

func TestJsonEncodeDecode(t *testing.T) {
	type Data struct {
		Msg string `json:"msg"`
	}
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	decoder := json.NewDecoder(&buff)
	for i := 0; i < 100; i++ {
		d := Data{Msg: "val" + strconv.Itoa(i)}
		err := encoder.Encode(d)
		require.NoError(t, err)
		err = encoder.Encode(d)
		require.NoError(t, err)
		{

			var dd Data
			err = decoder.Decode(&dd)
			require.NoError(t, err)
		}
		{
			var dd Data
			err = decoder.Decode(&dd)
			require.NoError(t, err)
			t.Log(dd)
			require.Equal(t, d, dd)

		}
	}

}
