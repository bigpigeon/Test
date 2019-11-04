package error_demo

import (
	"errors"
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

func errorDeferProcess() (int, error) {
	return 0, errors.New("has error")
}

func ErrorDeferProcess(t *testing.T) (v [20]byte, err error) {
	defer func() {
		if err != nil {
			t.Log(err)
		}
	}()

	n, err := errorDeferProcess()
	n = n
	if err != nil {
		return
	}
	return [20]byte{}, nil
}

//
//func ErrorDeferBlockProcess(t *testing.T) (err error) {
//	defer func() {
//		if err != nil {
//			t.Log(err)
//		}
//	}()
//	{
//		n, err := errorDeferProcess()
//		n = n
//		if err != nil {
//			t.Log(err)
//			return
//		}
//	}
//	return
//}

func TestDeferProcessError(t *testing.T) {
	ErrorDeferProcess(t)
	//ErrorDeferBlockProcess(t)
}

type HashKey [20]byte

func TestHashKeyCompare(t *testing.T) {
	key := HashKey{}
	if key != [20]byte{} {
		t.Log("no equal")
	} else {
		t.Log("no equal")
	}
}
