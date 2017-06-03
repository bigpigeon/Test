package reflectdemo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		t.Logf("variable a length is %d, last character is %c", aVal.Len(), aVal.Index(aVal.Len()-1))
	})
	t.Run("get list object attributes", func(t *testing.T) {
		var a []int = []int{1, 2, 3, 4}
		aVal := reflect.ValueOf(a)
		aType := reflect.TypeOf(a)
		t.Logf("variable a Value is %v", aVal.Interface())
		t.Logf("variable a Type is %s", aType.String())
		t.Logf("variable a length is %d, last int is %d", aVal.Len(), aVal.Index(aVal.Len()-1))
	})
	t.Run("change object value", func(t *testing.T) {
		var a int32 = 10
		// 要改变reflect的中的值必须是一个指针的Elem()后的reflect.Value
		aVal := reflect.ValueOf(a)
		aPointVal := reflect.ValueOf(&a)
		t.Logf("variable a set status %v", aVal.CanSet())
		t.Logf("variable a in point set status %v", aPointVal.Elem().CanSet())
		// 不用担心因为SetInt传入的是一个int64而设置负数会导致不正确，在这个函数中会根据int32而做转换
		aPointVal.Elem().SetInt(-20)
		t.Logf("variable a changed value %d", aPointVal.Elem().Interface())
		// Slice中元素都可以修改
		var b []int = []int{1, 2, 3, 4}
		bVal := reflect.ValueOf(b)
		t.Logf("slice b first element set status %v", bVal.Index(0).CanSet())
		// Map中元素都不能直接修改
		var c map[int]int = map[int]int{1: 2, 2: 4, 3: 6}
		cVal := reflect.ValueOf(c)
		oneValue := reflect.ValueOf(1)
		fourValue := reflect.ValueOf(4)

		t.Logf("map c the 1 key l element set status %v", cVal.MapIndex(oneValue).CanSet())
		// 不过可以使用SetMapIndex修改Val的值
		cVal.SetMapIndex(oneValue, oneValue)
		t.Logf("map c the value with 1 key is %d", cVal.MapIndex(oneValue))
		// 也可以对不存在的key设值
		cVal.SetMapIndex(fourValue, oneValue)
		t.Logf("now map c is %v", cVal)
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
	//	var RecursionGetValue func(reflect.Value, func(field *reflect.StructField))
	//	RecursionGetValue = func(structVal reflect.Value, f func(field *reflect.Value)) {
	//		for i := 0; i < structVal.NumField(); i++ {

	//		}
	//	}
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
				t.Logf("%s, this field is a %s and kind is %v", field.Name, field.Type, field.Type.Kind())
			}
		})
	})
	t.Run("get method type with struct", func(t *testing.T) {
		structVal := reflect.TypeOf(app)
		RecursionReadMethod(structVal, func(method *reflect.Method) {
			t.Logf("method name '%s' function source '%s' kind '%v'", method.Name, method.Type, method.Type.Kind())
		})
	})
	t.Run("get tag with struct", func(t *testing.T) {
		structVal := reflect.TypeOf(app)
		RecursionGetField("", structVal, func(field *reflect.StructField) {
			t.Logf("this field %s, json field '%s'", field.Name, field.Tag.Get("json"))
		})
	})
	t.Run("guess field with struct", func(t *testing.T) {
		structType := reflect.TypeOf(MacApplication{})
		offsetmap := map[uintptr]reflect.StructField{}

		for i := 0; i < structType.NumField(); i++ {
			field := structType.Field(i)
			offsetmap[field.Offset] = field
		}

		{
			realFieldName := "Application"
			structVal := reflect.ValueOf(&app)
			ApplicationVal := reflect.ValueOf(&app.Application)
			offset := ApplicationVal.Elem().UnsafeAddr() - structVal.Elem().UnsafeAddr()
			assert.Equal(t, realFieldName, offsetmap[offset].Name)
			t.Logf("I guess this field is %s, real field is %s", offsetmap[offset].Name, realFieldName)
		}
		{
			realFieldName := "AppleStore"
			structVal := reflect.ValueOf(&app)
			AppleStoreVal := reflect.ValueOf(&app.AppleStore)
			offset := AppleStoreVal.Elem().UnsafeAddr() - structVal.Elem().UnsafeAddr()
			assert.Equal(t, realFieldName, offsetmap[offset].Name)
			t.Logf("I guess this field is %s, real field is %s", offsetmap[offset].Name, realFieldName)
		}
	})
}
