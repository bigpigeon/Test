package type_convert

import (
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
