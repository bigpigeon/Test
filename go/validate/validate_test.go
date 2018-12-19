package validate

import (
	validator "gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestMaxMin(t *testing.T) {
	type TestData struct {
		Val  int   `validate:"max=10,min=1"`
		List []int `validate:"min=1"`
	}
	data := TestData{Val: 5}
	validate := validator.New()
	err := validate.Struct(data)
	t.Log(err)
}

func TestStartEnd(t *testing.T) {
	type TestData struct {
		Start int64
		End   int64 `validate:"gtefield=Start"`
	}
	data := TestData{Start: 5}
	validate := validator.New()
	err := validate.Struct(data)
	t.Log(err)
}
