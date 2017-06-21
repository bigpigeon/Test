package gormdemo

import (
	//	"encoding/json"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	t.Run("query data", func(t *testing.T) {
		TraversalTest(t, func(db *gorm.DB) {
			{
				var product Product
				assert.Nil(t, db.First(&product).Error)
				t.Logf("%+v", product)
			}
			{
				var product Product
				assert.Nil(t, db.Last(&product).Error)
				t.Logf("%+v", product)
			}
			{
				var chars []GreekAlphabet
				assert.Nil(t, db.Find(&chars).Error)
				t.Logf("%10s\t%s\t%s", "name", "upper", "lower")
				for _, c := range chars {
					t.Logf("%10s\t%c\t%c", c.LatinName, c.UpperCode, c.LowerCode)
				}
			}
			{
				var products []Product
				assert.Nil(t, db.Find(&products).Error)
				for _, p := range products {
					t.Logf("%+v", p)
				}
			}
			{
				var product Product
				assert.Nil(t, db.Where(&Product{Name: "xiaomi6"}).First(&product).Error)
				t.Logf("%+v", product)
			}
			// 使用Proload查询Product中的Origin字段
			{
				var product Product
				assert.Nil(t, db.Preload("Origin").Where(&Product{Name: "xiaomi6"}).First(&product).Error)
				t.Logf("%+v", product)
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
				assert.Nil(t, db.Model(&Product{}).Updates(&Product{Description: "also nothing here"}).Error)
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
