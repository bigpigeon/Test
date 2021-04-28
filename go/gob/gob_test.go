/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package gob_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type Doc struct {
	Title   string
	Content string
}

type OtherDoc Doc

const gobB64Data = "J/+BAwEBA0RvYwH/ggABAgEFVGl0bGUBDAABB0NvbnRlbnQBDAAAAB3/ggEKc29tZSB0aXRsZQEMc29tZSBkYXRhCgoKAA=="

var doc = Doc{
	Title:   "some title",
	Content: "some data\n\n\n",
}

func TestGobEncode(t *testing.T) {

	var buff bytes.Buffer

	err := gob.NewEncoder(&buff).Encode(&doc)
	require.NoError(t, err)
	b64EncodeData := base64.StdEncoding.EncodeToString(buff.Bytes())
	t.Log(b64EncodeData)
	require.Equal(t, b64EncodeData, gobB64Data)
	{
		var ddoc Doc
		err := gob.NewDecoder(&buff).Decode(&ddoc)
		require.NoError(t, err)
		t.Logf("%#v\n", ddoc)
	}
}

func TestGobDecode(t *testing.T) {
	buff, err := base64.StdEncoding.DecodeString(gobB64Data)
	require.NoError(t, err)
	var ddoc Doc
	err = gob.NewDecoder(bytes.NewReader(buff)).Decode(&ddoc)
	require.NoError(t, err)
	t.Logf("%#v\n", ddoc)
}

func TestGobDecodeInterface(t *testing.T) {
	buff, err := base64.StdEncoding.DecodeString(gobB64Data)
	require.NoError(t, err)
	var ddoc interface{}
	err = gob.NewDecoder(bytes.NewReader(buff)).Decode(&ddoc)
	require.NoError(t, err)
	t.Logf("%#v\n", ddoc)
}

func TestGobHasher(t *testing.T) {
	hasher := sha256.New()
	data := []byte("1234567890")
	_, err := hasher.Write(data)
	require.NoError(t, err)
	right := hasher.Sum(nil)
	t.Log(base64.RawStdEncoding.EncodeToString(right))
	{
		hasher := sha256.New()
		_, err := hasher.Write(data[:5])
		require.NoError(t, err)
		var buff bytes.Buffer
		gobEncoder := gob.NewEncoder(&buff)
		err = gobEncoder.EncodeValue(reflect.ValueOf(hasher).Elem())
		require.NoError(t, err)
		gobDecoder := gob.NewDecoder(&buff)
		decoderHasher := sha256.New()
		err = gobDecoder.Decode(decoderHasher)
		require.NoError(t, err)
		_, err = decoderHasher.Write(data[5:])
		require.NoError(t, err)
		t.Log("gob hash", base64.RawStdEncoding.EncodeToString(right))
	}
}
