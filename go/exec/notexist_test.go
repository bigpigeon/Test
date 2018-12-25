/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package exec

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
)

func TestExecNotExistCmd(t *testing.T) {
	cmd := exec.Command("xxx")
	out, err := cmd.Output()
	require.NoError(t, err)
	t.Log(out)
}
