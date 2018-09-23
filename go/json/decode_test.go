/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package json

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	IntVal int64
}

func TestJsonEncode(t *testing.T) {
	data := TestData{
		IntVal: 6446744073709551610,
	}
	buff := bytes.Buffer{}
	err := json.NewEncoder(&buff).Encode(data)
	assert.NoError(t, err)
	t.Log(buff.String())
}
