package main

import (
	"github.com/bigpigeon/toyorm"
	"log"
)

type Location struct {
	ID  uint32  `toyorm:"primary key;auto_increment"`
	Lat float64 `toyorm:"type:decimal(10,8)"`
	Lng float64 `toyorm:"type:decimal(11,8)"`
}

func main() {
	DB, err := toyorm.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	location := DB.Model(&Location{})
}
