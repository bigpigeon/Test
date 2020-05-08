package validate

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
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

func TestRegistryTag(t *testing.T) {
	type TestData struct {
		Url string `validate:"url,telnet"`
	}
	data := TestData{Url: "https://baidu.com"}
	validate := validator.New()
	err := validate.RegisterValidationCtx("telnet", func(ctx context.Context, fl validator.FieldLevel) bool {
		name := fl.FieldName()
		fmt.Println(fl.Param())
		fmt.Println(name)
		return true
	})
	require.NoError(t, err)
	err = validate.Struct(data)
	t.Log(err)
}
