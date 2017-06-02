package reflectdemo

import (
	"reflect"
	"testing"
)

func TestRelect(t *testing.T) {
	t.Run("get int object attributes", func(t *testing.T) {
		var a int = 20
		aVal := reflect.ValueOf(a)
		aType := reflect.TypeOf(a)

		t.Logf("variable a Value is %v", aVal.Interface())
		t.Logf("variable a Type is %s", aType.String())
	})
	t.Run("get string object attributes", func(t *testing.T) {
		var a string = "abcd1234"
		aVal := reflect.ValueOf(a)
		aType := reflect.TypeOf(a)
		t.Logf("variable a Value is %v", aVal.Interface())
		t.Logf("variable a Type is %s", aType.String())
		t.Logf("")
	})

	app := MacApplication{
		Application: Application{
			Name:        "sandbox tower defence",
			Description: "a rpg td game",
			packageData: []byte{},
		},
		AppleStore: "https://itunes.apple.com/us/app/example",
		Favorite:   0,
		Money: Money{
			MoneyType: MoneyTypeUS,
			Number:    0,
		},
	}
	var RecursionGetField func(string, reflect.Type, func(field *reflect.StructField))
	RecursionGetField = func(prefix string, structVal reflect.Type, f func(field *reflect.StructField)) {
		for i := 0; i < structVal.NumField(); i++ {
			field := structVal.Field(i)
			field.Name = prefix + "." + field.Name
			f(&field)
			if field.Type.Kind() == reflect.Struct {
				RecursionGetField(field.Name, field.Type, f)
			}
		}
	}
	var RecursionReadMethod func(reflect.Type, func(method *reflect.Method))
	RecursionReadMethod = func(structVal reflect.Type, f func(method *reflect.Method)) {
		for i := 0; i < structVal.NumMethod(); i++ {
			method := structVal.Method(i)
			f(&method)
		}
	}

	t.Run("get field type with struct", func(t *testing.T) {
		structVal := reflect.TypeOf(app)
		RecursionGetField("", structVal, func(field *reflect.StructField) {
			switch field.Type.Kind() {
			case reflect.Struct:
				// Anonymous表示一个匿名的字段
				if field.Anonymous == true {
					t.Logf("%s, this field is a embedded struct of %s ", field.Name, field.Type.Name())
				} else {
					t.Logf("%s, this field is a struct of %s ", field.Name, field.Type.Name())
				}
			default:
				t.Logf("%s, this field is a %s", field.Name, field.Type)
			}
		})
	})
	t.Run("get method type with struct", func(t *testing.T) {
		structVal := reflect.TypeOf(app)
		RecursionReadMethod(structVal, func(method *reflect.Method) {
			t.Logf("method name:%s,type: %s", method.Name, method.Type)
		})
	})
	t.Run("get tag with struct", func(t *testing.T) {
		structVal := reflect.TypeOf(app)
		RecursionGetField("", structVal, func(field *reflect.StructField) {
			t.Logf("%s, this field tag '%s'", field.Name, field.Tag)
		})
	})
}
