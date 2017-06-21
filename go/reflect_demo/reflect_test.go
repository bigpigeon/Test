package reflectdemo

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestRelect(t *testing.T) {
	// basic
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
		var a []int = []int{1, 2, 3, 4, 5}
		// 当使用append遇到list的空间不足时会重分配n*2的空间，这里只是为了让cap和length不同才这么做
		a = append(a, 6)
		aVal := reflect.ValueOf(a)
		aType := reflect.TypeOf(a)
		t.Logf("variable a Value is %v", aVal.Interface())
		t.Logf("variable a Type is %s", aType.String())
		t.Logf("variable a length is %d,cap is %d, last int is %d", aVal.Len(), aVal.Cap(), aVal.Index(aVal.Len()-1))
	})
	t.Run("get map object attributes", func(t *testing.T) {
		var a map[int]string = map[int]string{1: "a", 2: "b", 3: "c"}
		aVal := reflect.ValueOf(a)
		aType := reflect.TypeOf(a)
		OneVal := reflect.ValueOf(1)
		t.Logf("variable a Value is %v", aVal.Interface())
		t.Logf("variable a Type is %s", aType.String())
		t.Logf("variable a length is %d,one index value is %s", aVal.Len(), aVal.MapIndex(OneVal))
	})

	t.Run("change object value", func(t *testing.T) {
		var a int32 = 10
		// 要改变reflect的中的值必须是一个指针的Elem()下的reflect.Value
		aVal := reflect.ValueOf(a)
		aPointVal := reflect.ValueOf(&a)
		t.Logf("variable a set status %v", aVal.CanSet())
		t.Logf("variable a in point set status %v", aPointVal.Elem().CanSet())
		// 不用担心因为SetInt传入的是一个int64而设置负数会导致不正确，在这个函数中会根据int具体类型而做转换
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
	RecursionGetField = func(prefix string, structType reflect.Type, f func(field *reflect.StructField)) {
		for i := 0; i < structType.NumField(); i++ {
			field := structType.Field(i)
			field.Name = prefix + "." + field.Name
			f(&field)
			if field.Type.Kind() == reflect.Struct {
				// 匿名结构体字段中的子字段归并到当前结构体中
				if field.Anonymous == true {
					RecursionGetField(prefix, field.Type, f)
				} else {
					RecursionGetField(field.Name, field.Type, f)
				}
			}
		}
	}

	var RecursionReadMethod func(reflect.Type, func(method *reflect.Method))
	RecursionReadMethod = func(structType reflect.Type, f func(method *reflect.Method)) {
		for i := 0; i < structType.NumMethod(); i++ {
			method := structType.Method(i)
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
		structType := reflect.TypeOf(app)
		RecursionReadMethod(structType, func(method *reflect.Method) {
			t.Logf("method name '%s' function source '%s' kind '%v'", method.Name, method.Type, method.Type.Kind())
		})
	})
	t.Run("get tag with struct", func(t *testing.T) {
		structType := reflect.TypeOf(app)
		RecursionGetField("", structType, func(field *reflect.StructField) {
			t.Logf("this field %s,tag is '%s' json field '%s'", field.Name, field.Tag, field.Tag.Get("json"))
		})
	})

	// practical
	t.Run("map field with struct to sql", func(t *testing.T) {
		// 这里只是做简单取类型和字段名,不涉及主键和附加属性的处理
		// tableFields[table name][field name][field kind name]
		tableFields := map[string]map[string]string{}
		tables := []string{}
		sqlTypeMap := map[reflect.Kind]string{
			reflect.Int:     "integer",
			reflect.String:  "varchar(255)",
			reflect.Float64: "real",
		}
		structType := reflect.TypeOf(MacApplication{})
		var GetTableField func(string, reflect.Type)

		GetTableField = func(tName string, structType reflect.Type) {

			if tableFields[tName] == nil {
				tableFields[tName] = map[string]string{}
			}
			for i := 0; i < structType.NumField(); i++ {
				field := structType.Field(i)
				if field.Type.Kind() == reflect.Struct {
					if field.Anonymous == true {
						GetTableField(tName, field.Type)
					} else {
						GetTableField(field.Type.Name(), field.Type)
					}
				} else {
					tableFields[tName][field.Name] = sqlTypeMap[field.Type.Kind()]
				}
			}
		}
		GetTableField(structType.Name(), structType)
		// 通过tableFields表组装sql到tables中
		for tName, fields := range tableFields {
			fNameTypeList := []string{}
			for fName, fType := range fields {
				fNameTypeList = append(fNameTypeList, fmt.Sprintf(`"%s" %s`, fName, fType))
			}
			tables = append(tables, fmt.Sprintf(`CREATE TABLE "%s"(%s)`, tName, strings.Join(fNameTypeList, ",")))
		}
		for _, tab := range tables {
			t.Log(tab)
		}
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
			//			structVal := reflect.ValueOf(&app)
			//			ApplicationVal := reflect.ValueOf(&app.Application)
			offset := unsafe.Offsetof(app.Application)
			assert.Equal(t, realFieldName, offsetmap[offset].Name)
			t.Logf("I guess this field is %s, real field is %s", offsetmap[offset].Name, realFieldName)
		}
		{
			realFieldName := "AppleStore"
			//			structVal := reflect.ValueOf(&app)
			//			AppleStoreVal := reflect.ValueOf(&app.AppleStore)
			offset := unsafe.Offsetof(app.AppleStore)
			assert.Equal(t, realFieldName, offsetmap[offset].Name)
			t.Logf("I guess this field is %s, real field is %s", offsetmap[offset].Name, realFieldName)
		}
	})
	t.Run("compare type", func(t *testing.T) {
		a := MoneyTypeCN
		b := 2
		aType := reflect.TypeOf(a)
		bType := reflect.TypeOf(b)
		// 虽然他们的种类一样但类型不一样(kind翻译为种类是为了不和Type混淆)
		t.Logf("a kind is %v, b kind is %v", aType.Kind(), bType.Kind())
		t.Logf("a Type is %v, b Type is %v", aType.String(), bType.String())
		// 也可以通过转为interface直接比较类型
		_, aIsInt := reflect.New(aType).Interface().(int)
		t.Logf("a is int?%v", aIsInt)
		_, bIsMoneyType := reflect.New(bType).Interface().(MoneyType)
		t.Logf("b is MoneyType?%v", bIsMoneyType)
	})
}
