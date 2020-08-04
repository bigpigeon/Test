/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package file

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	fd, err := os.OpenFile("testfile", os.O_RDONLY|os.O_CREATE, 0644)
	require.NoError(t, err)
	n, err := fd.Seek(4, 0)
	require.NoError(t, err)
	t.Log(n)
	data, err := ioutil.ReadAll(fd)
	require.NoError(t, err)
	t.Log(string(data))
}

func TestContinueWrite(t *testing.T) {
	fd, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	defer fd.Close()
	require.NoError(t, err)
	for {
		_, err := fd.Write([]byte("some data\n"))
		require.NoError(t, err)
		time.Sleep(1 * time.Second)
	}
}

func TestReadNotDir(t *testing.T) {
	_, err := os.OpenFile("no_exist/testfile", os.O_WRONLY|os.O_CREATE, 0644)
	require.True(t, os.IsNotExist(err))
	t.Log(err)
}

func TestFileXAttr(t *testing.T) {
	os.Remove("testxattr2")
	fd, err := os.OpenFile("testxattr", os.O_WRONLY|os.O_CREATE, 0644)
	require.NoError(t, err)
	fd.Close()
	err = syscall.Setxattr("testxattr", "user.x", []byte("hash value"), 0)
	require.NoError(t, err)
	data := make([]byte, 1024)
	sz, err := syscall.Getxattr("testxattr", "user.x", data)
	require.NoError(t, err)
	t.Log(string(data[:sz]))
	err = os.Link("testxattr", "testxattr2")
	require.NoError(t, err)

	sz, err = syscall.Getxattr("testxattr2", "user.x", data)
	require.NoError(t, err)
	t.Log(string(data[:sz]))
}
