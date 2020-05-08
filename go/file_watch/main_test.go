/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFileWatch(t *testing.T) {
	f, err := os.Open("./")
	require.NoError(t, err)

	fileInfos, err := f.Readdir(0)
	require.NoError(t, err)
	for _, info := range fileInfos {

		t.Log(info.Name(), info.ModTime())

		t.Logf("%#v\n", info.Sys())
	}
}
