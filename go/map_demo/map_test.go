package map_demo

import (
	"reflect"
	"testing"
)

func TestMapPoint(t *testing.T) {
	m := make(map[int]int, 0)
	m2 := m
	m3 := reflect.ValueOf(m)
	t.Logf("m1 %p", m)
	t.Logf("m2 %p", m2)
	t.Logf("m3 %p", m3.Interface())
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
	t.Logf("m1 %p", m)
	t.Logf("m2 %p", m2)
	t.Logf("m3 %p", m3.Interface())

}
