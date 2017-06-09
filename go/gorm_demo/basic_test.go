package gormdemo

import (
	//	"encoding/json"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	t.Run("create table", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			tables := []interface{}{&Category{}, &Email{}, &Origin{}, &Language{}, &Product{}}
			db.DropTableIfExists(tables...)
			db.CreateTable(tables...)
		})
	})
	t.Run("add data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			for _, product := range SampleData() {
				assert.Nil(t, db.Create(&product).Error)
			}

		})
	})
	t.Run("query data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			{
				var product Product
				assert.Nil(t, db.First(&product).Error)
				str, err := Encode(&product)
				assert.Nil(t, err)
				t.Log(str)
			}

			{
				var product Product
				assert.Nil(t, db.Last(&product).Error)
				str, err := Encode(&product)
				assert.Nil(t, err)
				t.Log(str)
			}
			{
				var products []Product
				assert.Nil(t, db.Find(&products).Error)
				str, err := Encode(&products)
				assert.Nil(t, err)
				t.Log(str)
			}
			{
				var product Product
				assert.Nil(t, db.Where(&Product{Name: "xiaomi6"}).First(&product).Error)
				str, err := Encode(&product)
				assert.Nil(t, err)
				t.Log(str)
			}
			// 使用Proload查询Product中的Origin字段
			{
				var product Product
				assert.Nil(t, db.Preload("Origin").Where(&Product{Name: "xiaomi6"}).First(&product).Error)
				str, err := Encode(&product)
				assert.Nil(t, err)
				t.Log(str)
			}
		})
	})
	t.Run("upadte data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			// Save不会忽略0值
			{
				var xiaomi Product
				db.Where(&Product{Name: "xiaomi6"}).First(&xiaomi)
				xiaomi.Sid = 0
				assert.Nil(t, db.Save(&xiaomi).Error)
				var product Product
				db.Where(&Product{Name: "xiaomi6"}).First(&product)
				assert.Equal(t, product.Sid, 0)
			}
			// update 可以一次更新多条数据
			{
				assert.Nil(t, db.Model(&Product{}).Update("Description", "also nothing here").Error)
				products := []Product{}
				db.Find(&products)
				for _, p := range products {
					t.Log(p.Name, p.Description)
				}
			}
		})
	})
	t.Run("delete data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			// 删除数据时要保证被删除数据的主键不能为空，不然会吧整个表的数据都删掉
			// 因为product中包含DeleteAt字段，所以并不会数据并不会真的被删除，只是设置了DeleteAt为当前时间
			{
				var meat Product
				db.Where(&Product{Name: "wild boar meat"}).First(&meat)
				assert.NotEqual(t, meat.ID, 0)
				assert.Nil(t, db.Delete(&meat).Error)
				var product Product
				db.Where(&Product{Name: "wild boar meat"}).First(&product)
				assert.Equal(t, &product, &Product{})

			}
			// 如果数据表没有DeletedAt字段，那么调用Delete会物理删除数据
			{
				var email Email
				db.First(&email)
				id := email.ID
				db.Delete(&email)
				var nullEmail Email
				db.Find(&Email{ID: id}).First(&nullEmail)
				assert.Equal(t, &nullEmail, &Email{})
			}
		})
	})
}
