package reflect_test

import (
	"reflect"
	"testing"
	"time"
)

type NotNilTime struct {
	time.Time
}

type TestType struct {
	Name string
}

func TestTypePath(t *testing.T) {
	typ := reflect.TypeOf(TestType{})
	t.Log(typ.PkgPath(), typ.Name())
}
