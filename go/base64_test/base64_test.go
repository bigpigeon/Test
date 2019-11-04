/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package base64_test

import (
	"crypto/rand"
	"encoding/ascii85"
	"encoding/base64"
	"encoding/hex"
	"github.com/stretchr/testify/require"

	"testing"
)

func BenchmarkBase64Encode(b *testing.B) {
	data := make([]byte, 600*4)
	_, err := rand.Read(data)
	require.NoError(b, err)
	dst := make([]byte, 1200*4)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.Encode(dst, data)
	}
}

func BenchmarkBase85Encode(b *testing.B) {
	data := make([]byte, 600*4)
	_, err := rand.Read(data)
	require.NoError(b, err)
	dst := make([]byte, 800*4)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ascii85.Encode(dst, data)
	}
}

func BenchmarkHexEncode(b *testing.B) {
	data := make([]byte, 600*4)
	_, err := rand.Read(data)
	require.NoError(b, err)
	dst := make([]byte, 1200*4)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hex.Encode(dst, data)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	data := make([]byte, 600*4)
	_, err := rand.Read(data)
	encoder := base64.StdEncoding.EncodeToString(data)

	require.NoError(b, err)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := base64.StdEncoding.DecodeString(encoder)
		require.NoError(b, err)
	}
}
