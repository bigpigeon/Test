package copy

import (
	"bytes"
	"encoding/gob"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type TestCopyData struct {
	SVal  string
	Embed *struct {
		IVal int
		SVal string
	}
}

// reflect.Copy is not deep copy
func TestReflectCopy(t *testing.T) {
	val := []TestCopyData{{
		SVal: "copy data",
		Embed: &struct {
			IVal int
			SVal string
		}{
			IVal: 1,
			SVal: "embed copy data",
		},
	}}
	newVal := make([]TestCopyData, len(val))
	reflect.Copy(reflect.ValueOf(newVal), reflect.ValueOf(val))
	t.Logf("old embed val ptr %p new embed ptr val %p", val[0].Embed, newVal[0].Embed)
	t.Logf("old embed val %v new embed val %v", *val[0].Embed, *newVal[0].Embed)
	assert.True(t, newVal[0].Embed != val[0].Embed)
}

func TestGobCopy(t *testing.T) {
	val := []TestCopyData{{
		SVal: "copy data",
		Embed: &struct {
			IVal int
			SVal string
		}{
			IVal: 1,
			SVal: "embed copy data",
		},
	}}
	var newVal []TestCopyData
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(val)
	require.NoError(t, err)
	err = gob.NewDecoder(&buff).Decode(&newVal)
	require.NoError(t, err)
	t.Logf("old embed val ptr %p new embed ptr val %p", val[0].Embed, newVal[0].Embed)
	t.Logf("old embed val %v new embed val %v", *val[0].Embed, *newVal[0].Embed)

	assert.True(t, newVal[0].Embed != val[0].Embed)
	assert.Equal(t, *newVal[0].Embed, *val[0].Embed)
}

func TestSliceCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 4)
	copy(b[:2], a[:2])
	copy(b[2:], a[4:])
	t.Log(b)
	copy(b[2:], a[6:])
}
