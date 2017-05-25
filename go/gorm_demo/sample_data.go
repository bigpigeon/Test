package gormdemo

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	Name        string `gorm:"primary_key"`
	Description string `gorm:"size:255;default:'nothing in here'"`
}

type Email struct {
	ID         int
	UserId     int
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Origin struct {
	ID        int
	ProductID uint
	Address1  string `gorm:"not null;unique"`
	Address2  string `gorm:"unique"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}

type Product struct {
	gorm.Model
	Name string `gorm:"index;size:255"`

	Sid         int        `gorm:"unique_index"`
	Categories  []Category `gorm:"many2many:categories_product;"`
	Emails      []Email    `gorm:"ForeignKey:UserId"`
	Origin      *Origin
	Languages   []Language
	Score       *float64 `gorm:"not null;default:1.0"`
	Description string   `gorm:"size:255;default:'nothing in here'"`
}
