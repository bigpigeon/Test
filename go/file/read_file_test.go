/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package file

import (
	"encoding/base64"
	"encoding/binary"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"os"
	"strconv"
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
	//os.Remove("testxattr2")
	//fd, err := os.OpenFile("testxattr", os.O_WRONLY|os.O_CREATE, 0644)
	//require.NoError(t, err)
	//fd.Close()
	for i := 4; i < 6; i++ {
		err := syscall.Setxattr("testxattr", "user."+strconv.Itoa(i), []byte("hash value"), 0)
		require.NoError(t, err)
	}

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

func TestWriteOverride(t *testing.T) {
	fd, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	require.NoError(t, err)
	defer fd.Close()
	_, err = fd.WriteAt([]byte("123456"), 0)
	require.NoError(t, err)
	_, err = fd.Seek(0, io.SeekStart)
	require.NoError(t, err)
	_, err = fd.WriteAt([]byte("789"), 0)
	require.NoError(t, err)
	err = fd.Truncate(3)
	require.NoError(t, err)
}

func TestSeek(t *testing.T) {
	fd, err := os.OpenFile("testfile", os.O_RDONLY|os.O_CREATE, 0644)
	require.NoError(t, err)
	_, err = fd.Seek(5, io.SeekStart)
	require.NoError(t, err)
	off, err := fd.Seek(-6, io.SeekCurrent)
	t.Log(off)
	t.Log(err)
}

func TestFileStat(t *testing.T) {
	for _, f := range []string{"/dev/sda", "/dev", "/dev/sdb", "/home", "/home/jia", "/"} {

		var statfs syscall.Statfs_t
		err := syscall.Statfs(f, &statfs)
		require.NoError(t, err)

		b := make([]byte, 8)
		binary.BigEndian.PutUint32(b[:4], uint32(statfs.Fsid.X__val[0]))
		binary.BigEndian.PutUint32(b[4:], uint32(statfs.Fsid.X__val[1]))
		t.Logf("%s sys stat \t %+v\n", f, statfs)
		t.Logf("uuid %s\n", base64.RawURLEncoding.EncodeToString(b))
	}

}

func TestOsExecute(t *testing.T) {
	execVal, err := os.Executable()
	require.NoError(t, err)
	t.Logf(execVal)
}
