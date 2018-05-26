package map_add_demo

import (
	"testing"
)

type TestData struct {
	ID      int
	Data    string
	SubData []TestDataSub
}

type TestDataSub struct {
	MainId int
	Data   string
}

type MapAdd struct {
	Data map[int]*TestData
}

func (m *MapAdd) Map(data *TestData) {
	m.Data[data.ID] = data
}

func (m *MapAdd) Add(data TestDataSub) {
	m.Data[data.MainId].SubData = append(m.Data[data.MainId].SubData, data)
}

func TestMapReset(t *testing.T) {
	var a []*TestData
	for id, d := range []string{"a", "b", "c", "d", "e"} {
		a = append(a, &TestData{id, d, nil})
	}

	var b []TestDataSub
	for i := 0; i < 10; i++ {
		for _, d := range a {
			b = append(b, TestDataSub{d.ID, d.Data})
		}
	}

	m := MapAdd{map[int]*TestData{}}
	for _, d := range a {
		m.Map(d)
	}
	t.Log(m)
	for _, d := range b {
		m.Add(d)
	}
	for _, d := range a {
		t.Log(d)
	}
}
