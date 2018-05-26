package error_demo

import (
	"fmt"
	"testing"
)

type ErrMap map[int]string

func (e ErrMap) Error() string {
	s := ""
	for k, v := range e {
		s += fmt.Sprintf("%d %s;", k, v)
	}
	return s
}

func TestError(t *testing.T) {
	{
		e := ErrMap{}
		if e != nil {
			t.Log("err", e.Error())
		}
	}
	{
		var e ErrMap
		if e != nil {
			t.Log("err", e.Error())
		}
		e[2] = ""
		if e != nil {
			t.Log("err", e.Error())
		}
	}
}
