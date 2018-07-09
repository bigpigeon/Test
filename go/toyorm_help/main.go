package main

import (
	"encoding/json"
	"github.com/bigpigeon/toyorm"
)

type Invitation struct {
	ID               int    `toyorm:"primary key" json:"id"`
	MefeInvitationId string `json:"mefe___invitation___id"`
	// ...
}

func main() {
	jsonData := `{
	id: 22,
	mefe___invitation___id: "xx",
}`
	var data Invitation
	// use json lib to decode json bytes
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}
	// create database connect
	toy, err := toyorm.Open("mysql", "root:@tcp(localhost:3306)/toyorm_example?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	// insert data to mysql
	result, err := toy.Model(&data).Insert(&data)
	// check operation datta error
	if err != nil {
		panic(err)
	}
	// check mysql execute error
	if err := result.Err(); err != nil {
		panic(err)
	}

}
