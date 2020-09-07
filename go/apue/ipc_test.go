/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package apue

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
	"testing"
)

func TestSem(t *testing.T) {

}

func shmget(t *testing.T, length int) []byte {
	// map /dev/null
	data, err := syscall.Mmap(-1, 0, length, syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_ANON|syscall.MAP_SHARED)
	require.NoError(t, err)

	return data
}

func TestPoll(t *testing.T) {
	fd, err := os.OpenFile("testdata/poll_data", os.O_WRONLY, 0644)
	require.NoError(t, err)
	pollFd := []unix.PollFd{
		{
			Fd:      int32(fd.Fd()),
			Events:  unix.POLLOUT | unix.POLLERR,
			Revents: 0,
		},
	}
	n, err := unix.Poll(pollFd, 1)
	require.NoError(t, err)
	switch n {
	case -1:
		t.Log("poll error")
	case 0:
		t.Log("timeout")
	default:
		t.Logf("event value is %s", PollStr(pollFd[0].Revents))

	}
}

func PollStr(p int16) string {
	if p|unix.POLLERR == unix.POLLERR {
		return "ERR"
	}
	s := ""
	if p|unix.POLLOUT == unix.POLLOUT {
		s += "OUT|"
	}
	if p|unix.POLLIN == unix.POLLIN {
		s += "IN|"
	}
	return s
}
