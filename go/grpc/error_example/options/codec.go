package options

import (
	"bytes"
	"encoding/gob"
	"google.golang.org/grpc/codes"
	"reflect"
)

type CodeError struct {
	Code string
	Msg  string
}

func (e CodeError) Error() string {
	return e.Code + " " + e.Msg
}

type UnknownError struct {
	Type string
	Str  string
}

func (e UnknownError) Error() string {
	return e.Str
}

type CodecHandlers struct{}

func (CodecHandlers) Marshal(v interface{}) ([]byte, error) {
	buff := bytes.Buffer{}
	encoder := gob.NewEncoder(&buff)
	if e, ok := v.(CodeError); ok {
		return buff.Bytes(), encoder.Encode(e)
	} else {
		val := reflect.ValueOf(v)
		err := UnknownError{
			Type: val.Type().String(),
			Str:  val.String(),
		}

		return buff.Bytes(), encoder.Encode(err)
	}

}

func (CodecHandlers) Unmarshal(data []byte, v interface{}) error {
	//decoder := gob.NewDecoder(bytes.NewReader(data))
	//err := decoder.Decode(&v)
	//if err != nil {
	//	fmt.Println("decode err", err)
	//	return err
	//}
	//fmt.Printf("%T %v\n", v, v)
	return nil
}

func (CodecHandlers) String() string {
	return "codec handlers"
}

const TooLong codes.Code = 1024
