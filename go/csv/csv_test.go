/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package csv

import (
	"encoding/csv"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"testing"
)

const testData = `
aa,bb,cc,dd`

func TestReader(t *testing.T) {
	reader := csv.NewReader(strings.NewReader(testData))
	record, err := reader.Read()
	require.NoError(t, err)
	t.Log(record)
}

func TestOffset(t *testing.T) {
	fp, err := os.Open("test.csv")
	require.NoError(t, err)
	reader := csv.NewReader(fp)
	require.NoError(t, err)

	start, err := fp.Seek(0, io.SeekCurrent)
	require.NoError(t, err)
	record, err := reader.Read()
	require.NoError(t, err)
	end, err := fp.Seek(0, io.SeekCurrent)
	require.NoError(t, err)
	t.Log(record, start, end)
	{
		record, err := reader.Read()
		require.NoError(t, err)
		end, err := fp.Seek(0, io.SeekCurrent)
		t.Log(record, start, end)
	}

}
