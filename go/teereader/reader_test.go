/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package teereader

import (
	"crypto/sha256"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestTeeReader(t *testing.T) {
	data := "12345678901234567890"
	hasher := sha256.New()
	reader := io.TeeReader(strings.NewReader(data), hasher)
	rData, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	require.Equal(t, string(rData), data)
	sum := hasher.Sum(nil)
	sum2 := sha256.Sum256([]byte(data))
	require.Equal(t, sum, sum2[:])
}
