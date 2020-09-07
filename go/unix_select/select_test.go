/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package unix_select

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"os"
	"testing"
)

func TestSelect(t *testing.T) {
	//xx, err := os.Open("testdata/select_file")
	//require.NoError(t, err)
	rfs := unix.FdSet{}
	rfs.Set(0)
	wfs := rfs
	efs := rfs
	t.Log("before select")
	t.Log("read fd set", rfs)
	t.Log("write fd set", wfs)
	t.Log("error fd set", efs)
	//rfs.Set(2)
	//rfs.Set(3)

	n, err := unix.Select(1024, &rfs, &wfs, &efs, nil)
	require.NoError(t, err)
	t.Log("after select")
	t.Log("read fd set", rfs)
	t.Log("write fd set", wfs)
	t.Log("error fd set", efs)
	t.Log("select result", n)

}

func TestPoll(t *testing.T) {
	fds := []unix.PollFd{
		{
			Fd:      0,
			Events:  unix.POLLIN,
			Revents: 0,
		},
	}
	n, err := unix.Poll(fds, -1)
	require.NoError(t, err)
	t.Log(n, fds[0])
}

func TestRWV(t *testing.T) {
	data := [][]byte{
		[]byte("abc"),
		[]byte("123"),
		[]byte("456\n"),
	}
	n, err := unix.Writev(int(os.Stdout.Fd()), data)
	require.NoError(t, err)
	t.Log("readv result", n)
}
