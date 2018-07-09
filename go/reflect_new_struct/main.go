/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StructA struct {
	Username string `json:"user"`
	Process  string `json:"process"`
}

func main() {

	var test1 StructA
	err := json.Unmarshal([]byte(`{"user": "user123", "process": "something"}`), &test1)
	if err != nil {
		fmt.Println(err)
	}

	// do some work with test1

	jsonByte, err := json.Marshal(&test1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonByte))

}

func (u *StructA) MarshalJSON() ([]byte, error) {
	// get old struct fields
	uType := reflect.TypeOf(u).Elem()
	userNameField, _ := uType.FieldByName("Username")
	// set username field tag
	userNameField.Tag = `json:"username"`
	processField, _ := uType.FieldByName("Process")
	newType := reflect.StructOf([]reflect.StructField{userNameField, processField})
	// set new value field
	oldValue := reflect.ValueOf(u).Elem()
	newtValue := reflect.New(newType).Elem()
	for i := 0; i < oldValue.NumField(); i++ {
		newtValue.Field(i).Set(oldValue.Field(i))
	}
	return json.Marshal(newtValue.Interface())
}
