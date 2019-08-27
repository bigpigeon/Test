package error_demo

import (
	"fmt"
	"strings"
	"sync"
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

func TestPanic(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err", r)
		}
	}()
	go func() {
		//defer wg.Done()
		a := []int{1, 2, 4}
		fmt.Println(a[4])
	}()
	wg.Wait()
	t.Log("done!!")
}

type ErrSlice []string

func (e ErrSlice) Error() string {
	return strings.Join(e, "\n")
}
func TestSliceError(t *testing.T) {

	err := func() error {
		var errSlice ErrSlice
		return errSlice
	}()
	if err == nil {
		t.Log("nil err")
	} else {
		t.Log("not nil err", err)
	}
}
