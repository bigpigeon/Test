package type_convert

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"testing"
	"time"
	"unsafe"
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

func TestTypeAddr(t *testing.T) {
	type S struct {
		A int32
		B int64
		C struct{}
	}
	s := S{A: 1, B: 2}
	fmt.Printf("%p %p %p %d\n", &s.A, &s.B, &s.C, unsafe.Sizeof(s))
	b := struct {
	}{}
	c := struct {
	}{}
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
}

func TestStruct(t *testing.T) {
	data := map[string]struct{}{}
	data["a"] = struct{}{}

	fmt.Printf("%p %p\n", data["a"], data["b"])
}

func TestMapStruct(t *testing.T) {
	m := make(map[int]struct{}, 10<<20)
	for i := 0; i < 10000; i++ {
		m[i] = struct{}{}
	}
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	fmt.Printf("stat %d %d\n", stat.Alloc, stat.TotalAlloc)
}

func TestMapBool(t *testing.T) {
	m := make(map[int]bool, 10<<20)
	for i := 0; i < 10000; i++ {
		m[i] = false
	}
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	fmt.Printf("stat %d %d\n", stat.Alloc, stat.TotalAlloc)
}
