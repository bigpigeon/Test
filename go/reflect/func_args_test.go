package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type fooPrintData struct {
	Data string
}

func fooPrint(name fooPrintData) {
	fmt.Println("hello", name)
}

func TestGetFuncArgsInfo(t *testing.T) {
	ft := reflect.TypeOf(fooPrint)
	nameT := ft.In(0)
	t.Log(nameT.Field(0).Name, nameT.Field(0).Type)
	// use caller
}
