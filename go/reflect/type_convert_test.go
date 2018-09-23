package reflect_test

import (
	"reflect"
	"testing"
	"time"
)

type NotNilTime struct {
	time.Time
}

func TestConvert(t *testing.T) {
	a := NotNilTime{
		time.Now(),
	}
	aVal := reflect.ValueOf(a)

}
