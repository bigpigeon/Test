package type_convert

import (
	"io"
	"reflect"
	"testing"
	"time"
)

type CustomTime struct {
	*time.Time
}

func TestConvert(t *testing.T) {
	now := time.Now()
	ti := CustomTime{
		&now,
	}
	var v interface{} = ti
	_, ok := v.(*time.Time)
	t.Log(ok)
}

func TestName(t *testing.T) {
	//var i *int
	var reader io.Reader
	iTyp := reflect.TypeOf(reader)
	t.Log(iTyp.String())
}
