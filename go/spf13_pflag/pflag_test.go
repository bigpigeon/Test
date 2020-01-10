/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPflagGetString(t *testing.T) {
	fset := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fset.ParseErrorsWhitelist.UnknownFlags = true
	err := fset.Parse([]string{
		"--addr", "localhost", "--port", "8080",
	})
	require.NoError(t, err)

}
