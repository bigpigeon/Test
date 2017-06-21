package gormdemo

import (
	"reflect"

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

type GreekAlphabet struct {
	ID         uint   `gorm:"primary_key"`
	LatinName  string `gorm:"unique_index"`
	UpperCode  rune
	LowerCode  rune
	IsFrequent bool `gorm:"index"`
}

var FieldSelector struct {
	Product       Product
	GreekAlphabet GreekAlphabet
}

var OffsetSelector struct {
	OffsetMap map[reflect.Type]map[uintptr]string
}

func init() {
	OffsetSelector.OffsetMap = map[reflect.Type]map[uintptr]string{}
	fieldSelectVal := reflect.ValueOf(&FieldSelector)
	fieldSelectType := reflect.TypeOf(FieldSelector)
	for i := 0; i < fieldSelectType.NumField(); i++ {

		fieldVal := fieldSelectVal.Elem().Field(i)
		scope := &gorm.Scope{Value: fieldVal.Interface()}
		table := scope.GetModelStruct().ModelType
		gormFields := scope.Fields()
		OffsetSelector.OffsetMap[table] = map[uintptr]string{}
		for j := 0; j < len(gormFields); j++ {
			subfield := gormFields[j]
			offset := subfield.StructField.Struct.Offset
			OffsetSelector.OffsetMap[table][offset] = subfield.Name
		}
	}
}

func SampleProductData() []Product {
	return []Product{
		Product{
			Name: "iphone7",
			Sid:  1211,
			Categories: []Category{
				Category{"mobile phone", "a hand-held mobile radiotelephone for use in an area divided into small sections, each with its own short-range transmitter/receiver"},
				Category{"apple", ""},
			},
			Emails:    []Email{Email{Email: "example@domain.com", Subscribed: false}},
			Origin:    &Origin{Address1: "apple company address", Address2: "test"},
			Languages: []Language{Language{Name: "中国", Code: "cn"}, Language{Name: "美国", Code: "us"}},
			Score:     func(f float64) *float64 { return &f }(0.0),
		},
		Product{
			Name: "xiaomi6",
			Sid:  1311,
			Categories: []Category{
				Category{"mobile phone", "a hand-held mobile radiotelephone for use in an area divided into small sections, each with its own short-range transmitter/receiver"},
				Category{"xiaomi", ""},
			},
			Emails:    []Email{Email{Email: "example2@domain.com", Subscribed: false}},
			Origin:    &Origin{Address1: "xiaomi company address", Address2: ""},
			Languages: []Language{Language{Name: "中国", Code: "cn"}},
			Score:     func(f float64) *float64 { return &f }(2.0),
		},
		Product{
			Name: "wild boar meat",
			Sid:  9999,
			Categories: []Category{
				Category{"food", " sth solid for eating"},
				Category{"meat", ""},
			},
			Emails:    []Email{Email{Email: "example3@domain.com", Subscribed: false}},
			Origin:    &Origin{Address1: "163 company address", Address2: "163 company address2"},
			Languages: []Language{Language{Name: "中国", Code: "cn"}},
			Score:     func(f float64) *float64 { return &f }(3.0),
		},
	}
}

func SampleGreeceCharacterData() []GreekAlphabet {
	upperCodeIter := 'Α'
	lowerCodeIter := 'α'
	nameList := []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon", "Zeta",
		"Eta", "Theta", "Iota", "Kappa", "Lambda", "Mu", "Nu", "Xi", "Omicron",
		"Pi", "Rho"}
	charactList := []GreekAlphabet{}
	for _, name := range nameList {
		charactList = append(
			charactList,
			GreekAlphabet{
				LatinName: name,
				UpperCode: upperCodeIter,
				LowerCode: lowerCodeIter,
			},
		)
		upperCodeIter++
		lowerCodeIter++
	}
	upperCodeIter++
	lowerCodeIter++
	nameList = []string{"Sigma", "Tau", "Upsilon", "Phi", "Chi", "Psi", "Omega"}
	for _, name := range nameList {
		charactList = append(
			charactList,
			GreekAlphabet{
				LatinName: name,
				UpperCode: upperCodeIter,
				LowerCode: lowerCodeIter,
			},
		)
		upperCodeIter++
		lowerCodeIter++
	}
	return charactList
}
