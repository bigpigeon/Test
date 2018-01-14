package reflectdemo

import (
	"reflect"
	"testing"
)

func TestZeroValue(t *testing.T) {
	typeValue := []interface{}{
		true,
		1,
		1.0,
		1.0 + 1i,
		"",
		[]int{1, 2},
		[2]int{1, 2},
		map[int]int{1: 2},
		interface{}(2),
		func() int { return 2 },
		make(chan int, 2),
		struct{ Name string }{"code"},
	}
	for _, v := range typeValue {
		vValue := reflect.ValueOf(v)
		t.Logf("kind %s,zero value: %#v, nil value: %#v\n", vValue.Kind(), reflect.Zero(vValue.Type()), reflect.New(vValue.Type()).Elem())
	}
}
