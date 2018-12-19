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

func TestMapAssign(t *testing.T) {
	//{
	//	m1 := map[string]int{
	//		"1": 1,
	//		"2": 2,
	//	}
	//	m1Val := reflect.ValueOf(m1)
	//	m1Val.SetMapIndex(reflect.ValueOf("3"), reflect.ValueOf(1))
	//}
	//{
	//	m2 := map[int]string{
	//		1: "1",
	//	}
	//
	//	m2Val := reflect.ValueOf(m2)
	//	m2Val.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("3"))
	//}
	//{
	//	type MyType struct {
	//		ID int
	//	}
	//	m3 := map[MyType]string{
	//		MyType{1}: "1",
	//	}
	//
	//	mVal := reflect.ValueOf(m3)
	//	mVal.SetMapIndex(reflect.ValueOf(MyType{1}), reflect.ValueOf("3"))
	//}
	{
		m4 := map[interface{}]interface{}{
			2:   3,
			"1": 4,
		}
		mVal := reflect.ValueOf(m4)
		iVal := 2
		channel := make(chan int)
		mVal.SetMapIndex(reflect.ValueOf(struct{}{}), reflect.ValueOf("3"))
		mVal.SetMapIndex(reflect.ValueOf(&iVal), reflect.ValueOf("3"))
		mVal.SetMapIndex(reflect.ValueOf(channel), reflect.ValueOf("3"))
		mVal.SetMapIndex(reflect.ValueOf(0.0), reflect.ValueOf("3"))
		mVal.SetMapIndex(reflect.ValueOf(-0.0), reflect.ValueOf("3"))
		t.Log(m4)

	}
}

func TestMapList(t *testing.T) {
	a := make([]map[string]int, 2)
	a[1]["22"] = 3
	t.Log(a)
}
