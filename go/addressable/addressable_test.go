package addressable

import (
	"reflect"
	"testing"
)

func TestMapAddressable(t *testing.T) {
	m := map[int]int{1: 1}
	mValue := reflect.ValueOf(&m).Elem()
	vValue := mValue.MapIndex(reflect.ValueOf(1))
	t.Log(vValue.Interface())
	t.Logf("can address %v", vValue.CanAddr())
}

func TestStructAddressable(t *testing.T) {
	m := struct {
		ID int
	}{1}
	mValue := reflect.ValueOf(&m).Elem()
	vValue := mValue.Field(0)
	t.Log(vValue.Interface())
	t.Logf("can address %v", vValue.CanAddr())
}
