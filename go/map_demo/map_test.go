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

func TestMapGet(t *testing.T) {
	m := map[string]int{
		"1":    1,
		"2":    0,
		"案件编号": 0,
	}
	d, ok := m["2"]
	t.Log(d, ok)

	d, ok = m["案件编号"]
	t.Log(d, ok)
	t.Logf("%d %x", len("案件编号"), "案件编号")
	b := []byte{239, 187, 191}
	t.Log(string(b))

}
