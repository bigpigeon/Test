package gormdemo

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestScope(t *testing.T) {
	TraversalTest(t, func(db *gorm.DB) {
		tables := []interface{}{&Category{}, &Email{}, &Origin{}, &Language{}, &Product{}}
		db.DropTableIfExists(tables...)
		db.CreateTable(tables...)
		for _, p := range SampleData() {
			assert.Nil(t, db.Create(&p).Error)
		}
	})
	t.Run("a simple test", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			scope := func(db *gorm.DB) *gorm.DB {
				return db.Where("score > ?", 2.0)
			}
			products := []Product{}
			db.Model(&Product{}).Scopes(scope).Find(&products)
			for _, p := range products {
				t.Logf("'%s' product has high socre", p.Name)
			}
		})
	})

}
