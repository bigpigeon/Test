package gormdemo

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestAlpha(t *testing.T) {
	t.Run("create table", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			tables := []interface{}{&Category{}, &Email{}, &Origin{}, &Language{}, &Product{}, &GreekAlphabet{}}
			db.DropTableIfExists(tables...)
			db.CreateTable(tables...)
		})
	})
	t.Run("add data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			for _, product := range SampleProductData() {
				assert.Nil(t, db.Create(&product).Error)
			}
			for _, char := range SampleGreeceCharacterData() {
				assert.Nil(t, db.Create(&char).Error)
			}

		})
	})
}
