/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package json_test

import (
	"encoding/json"
	"testing"
)

var smallFixture []byte = []byte(`{
    "st": 1,
    "sid": 486,
    "tt": "active",
    "gr": 0,
    "uuid": "de305d54-75b4-431b-adb2-eb6b9e546014",
    "ip": "127.0.0.1",
    "ua": "user_agent",
    "tz": -6,
    "v": 1
}`)

type SmallPayload struct {
	St   int    `json:"st"`
	Sid  int    `json:"-"`
	Tt   string `json:"-"`
	Gr   int    `json:"-"`
	Uuid string `json:"uuid"`
	Ip   string `json:"-"`
	Ua   string `json:"ua"`
	Tz   int    `json:"tz"`
	V    int    `json:"-"`
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		json.Unmarshal(smallFixture, &data)
	}
}
