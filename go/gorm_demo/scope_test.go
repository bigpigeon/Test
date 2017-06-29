package gormdemo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestScope(t *testing.T) {
	t.Run("a simple test", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			firstAndLast := func(db *gorm.DB) *gorm.DB {
				return db.Where("latin_name in (?)", []string{"Alpha", "Omega"})
			}
			chars := []GreekAlphabet{}
			assert.Nil(t, db.Model(&GreekAlphabet{}).Scopes(firstAndLast).Find(&chars).Error)
			t.Logf("first character name is %s upper code is %c lower code is %c", chars[0].LatinName, chars[0].UpperCode, chars[0].LowerCode)
			t.Logf("last character name is %s upper code is %c lower code is %c", chars[1].LatinName, chars[1].UpperCode, chars[1].LowerCode)
		})
	})
	t.Run("update test", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			WhereIn := func(fieldOffset uintptr, set interface{}) func(db *gorm.DB) *gorm.DB {
				return func(db *gorm.DB) *gorm.DB {
					val := db.Value
					structType := reflect.TypeOf(val)
					// 获取非list或指针的reflect.Type
					for structType.Kind() == reflect.Slice || structType.Kind() == reflect.Ptr {
						structType = structType.Elem()
					}
					// Where的查询语句中用的是表字段名
					dbname, ok := OffsetSelector.DBNameMap[structType][fieldOffset]
					if ok == false {
						db.AddError(errors.New("offset is invalid"))
					}
					query := fmt.Sprintf("%s in (?)", dbname)
					return db.Where(query, set)
				}
			}
			frequentNames := []string{"Alpha", "Beta", "Gamma", "Delta", "Pi", "Lambda"}
			//获取GreekAlphabet.LatinName的offset,记住Offsetof中的参数是表达式，所以不能传参,比如xx := GreekAlphabet{}.LatinName;unsafe.Offsetof(xx)这样是不行的
			latinNameOffset := unsafe.Offsetof(GreekAlphabet{}.LatinName)
			assert.Nil(t, db.Model(&GreekAlphabet{}).
				Scopes(WhereIn(latinNameOffset, frequentNames)).
				Updates(&GreekAlphabet{IsFrequent: true}).Error,
			)
			frequents := []GreekAlphabet{}
			// 查看所有IsFrequent=true的集合
			db.Where(&GreekAlphabet{IsFrequent: true}).Find(&frequents)
			t.Logf("%10s\t%s\t%s\t%s", "name", "upper", "lower", "frequent")
			for _, c := range frequents {
				t.Logf("%10s\t%c\t%c\t%v", c.LatinName, c.UpperCode, c.LowerCode, c.IsFrequent)
			}
		})
	})

	// 利用scopes做一些定制方法
	t.Run("embed get struct field", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			FieldPreload := func(offset uintptr) func(db *gorm.DB) *gorm.DB {
				return func(db *gorm.DB) *gorm.DB {
					val := db.Value
					structType := reflect.TypeOf(val)
					// 获取非list或指针的reflect.Type
					for structType.Kind() == reflect.Slice || structType.Kind() == reflect.Ptr {
						structType = structType.Elem()
					}
					name, ok := OffsetSelector.NameMap[structType][offset]
					if ok == false {
						db.AddError(errors.New("offset is invalid"))
					}
					t.Log(OffsetSelector.NameMap[structType])
					return db.Preload(name)
				}
			}
			var product Product
			fieldOffset := unsafe.Offsetof(product.Origin)
			originPreload := FieldPreload(fieldOffset)
			assert.Nil(t, db.Model(&Product{}).Where(&Product{Name: "xiaomi6"}).Scopes(originPreload).First(&product).Error)
			// 看看查询结构是否正确
			t.Logf("this product name '%s', the address is '%v'", product.Name, product.Origin.Address1)
		})
	})
}
