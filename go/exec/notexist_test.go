/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package exec

import (
	"context"
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
	"time"
)

func TestExecNotExistCmd(t *testing.T) {
	cmd := exec.Command("xxx")
	out, err := cmd.Output()
	require.NoError(t, err)
	t.Log(out)
}

func TestExecCtxTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/sleep", "100")
	err := cmd.Start()
	require.NoError(t, err)
	err = cmd.Wait()
	require.NoError(t, err)

}
