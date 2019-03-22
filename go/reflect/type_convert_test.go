package reflect_test

import (
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/require"
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

func ToGrpcData(in, out interface{}) error {
	inVal := reflect.ValueOf(in)
	if inVal.Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	}
	inTyp := inVal.Type()

	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Ptr {
		return errors.New("out data must be point value")
	}

	outVal = outVal.Elem()
	outTyp := outVal.Type()

	strWrapperType := reflect.TypeOf(wrappers.StringValue{})
	// range all 'in' fields
	for i := 0; i < inVal.NumField(); i++ {
		fTyp := inTyp.Field(i)
		fVal := inVal.Field(i)
		if fTyp.Type.Kind() == reflect.Ptr {
			switch fTyp.Type.Elem().Kind() {
			case reflect.String: // only implement string in this test
				oFTyp, ok := outTyp.FieldByName(fTyp.Name)
				if ok == false {
					return errors.New("not match field in out value")
				}
				if oFTyp.Type.Elem() != strWrapperType {
					return errors.New("not match field in out value")
				}
				if fVal.IsNil() == false {
					outVal.FieldByName(fTyp.Name).Set(
						reflect.ValueOf(&wrappers.StringValue{
							Value: fVal.Elem().String(),
						}),
					)
				}
			}
		} else {
			outVal.FieldByName(fTyp.Name).Set(fVal)
		}
	}
	return nil
}

func TestGrpcConvert(t *testing.T) {
	type User struct {
		Name  string
		Value *string
	}
	type ServerUser struct {
		Name  string
		Value *wrappers.StringValue
	}
	usValue := "123"
	u1 := User{
		Name:  "tom",
		Value: &usValue,
	}
	u2 := ServerUser{}
	require.NoError(t, ToGrpcData(&u1, &u2))
	t.Log(u2)
}
