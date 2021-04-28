/*
 * Copyright 2021 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package fmt_example

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"testing"
)

func TestScan(t *testing.T) {
	fmt.Print("please input 123: ")

	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	require.NoError(t, err)
	require.Equal(t, string(pass), "123")

}
