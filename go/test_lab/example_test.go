/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package test_lab_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"test_lab"
)

func ExampleToNotPoint() {
	type TestSub struct {
		Data   *string
		Name   *string
		IntVal *int
	}
	type TestData struct {
		Sub  *TestSub
		Text *string
	}

	data := &TestData{}
	{
		var buff bytes.Buffer
		err := json.NewEncoder(&buff).Encode(&data)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s", buff.String())
	}

	test_lab.ToNotPoint(&data)
	{
		var buff bytes.Buffer
		err := json.NewEncoder(&buff).Encode(&data)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s", buff.String())
	}

	// Output:
	// {"Sub":null,"Text":null}
	// {"Sub":{"Data":"","Name":"","IntVal":0},"Text":""}
}
