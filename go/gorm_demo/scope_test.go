package gormdemo

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestScope(t *testing.T) {
	TraversalTest(t, func(db *gorm.DB) {
		tables := []interface{}{&Category{}, &Email{}, &Origin{}, &Language{}, &Product{}}
		db.DropTableIfExists(tables...)
		db.CreateTable(tables...)
	})
	t.Run("a easy test", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			db.Scopes()
		})
	})
	TraversalTest(t, func(db *gorm.DB) {
		if err := db.Close(); err != nil {
			t.Error(err)
		} else {
			t.Log("database was closed")
		}
	})

}
