package gormdemo

import (
	"errors"
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

			UpdateClassic := func(frequentNames []string, isclassic bool) func(db *gorm.DB) *gorm.DB {
				return func(db *gorm.DB) *gorm.DB {
					return db.Where("latin_name in (?)", frequentNames).Update("is_frequent", isclassic)
				}
			}
			frequentNames := []string{"Alpha", "Beta", "Gamma", "Delta", "Pi", "Lambda"}
			assert.Nil(t, db.Model(&GreekAlphabet{}).Scopes(UpdateClassic(frequentNames, true)).Error)
			frequents := []GreekAlphabet{}
			db.Where(&GreekAlphabet{IsFrequent: true}).Find(&frequents)
			t.Logf("%10s\t%s\t%s\t%s", "name", "upper", "lower", "frequent")
			for _, c := range frequents {
				t.Logf("%10s\t%c\t%c\t%v", c.LatinName, c.UpperCode, c.LowerCode, c.IsFrequent)
			}
		})
	})

	// 利用scopes做一些定制方法
	fieldPreload := func(offset uintptr) func(db *gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			val := db.Value
			structType := reflect.TypeOf(val)
			for structType.Kind() == reflect.Slice || structType.Kind() == reflect.Ptr {
				structType = structType.Elem()
			}
			name, ok := OffsetSelector.OffsetMap[structType][offset]
			if ok == false {
				db.AddError(errors.New("offset is invalid"))
			}
			t.Log(OffsetSelector.OffsetMap[structType])
			return db.Preload(name)
		}
	}
	t.Run("embed get struct field", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			var product Product
			originPreload := fieldPreload(unsafe.Offsetof(product.Origin))
			db.Model(&Product{}).Where(&Product{Name: "xiaomi6"}).Scopes(originPreload).First(&product)
			t.Logf("this product name '%s', the address is '%v'", product.Name, product.Origin.Address1)
		})
	})
}
