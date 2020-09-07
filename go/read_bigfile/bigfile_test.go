/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package read_bigfile

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestReadBigFile(t *testing.T) {
	f, err := os.Open("/home/benjamin/下载/ubuntu-18.04.2-live-server-amd64.iso")
	require.NoError(t, err)
	buff := make([]byte, 800<<20)
	n, err := f.Read(buff)
	require.NoError(t, err)
	t.Log(n)
}

func TestReadV(t *testing.T) {
	f, err := os.Open("/home/benjamin/下载/ubuntu-18.04.2-live-server-amd64.iso")
	require.NoError(t, err)
	buff := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		buff[i] = make([]byte, 100<<20)
	}
	n, err := unix.Readv(int(f.Fd()), buff)
	require.NoError(t, err)
	t.Log(n)
}

func writeSequence(data []byte) {
	const seq = 'Z' - 'A' + 1
	for i := 0; i < len(data); i++ {
		data[i] = byte(i%seq) + 'A'
	}
}

func TestMMap(t *testing.T) {
	tmpfile := path.Join(os.TempDir(), "mmap.tmp")
	// write some rand data
	{
		os.Remove(tmpfile)
		t.Log("temp file", tmpfile)
		data := make([]byte, 10<<10)
		writeSequence(data)

		err := ioutil.WriteFile(tmpfile, data, 0644)
		require.NoError(t, err)
	}
	{

		fd, err := unix.Open(tmpfile, unix.O_RDWR, 0644)
		require.NoError(t, err)
		data, err := unix.Mmap(fd, 0, 1024, unix.PROT_READ, unix.MAP_SHARED)
		require.NoError(t, err)
		//data[2] = 'd' // will be panic
		t.Log(len(data), cap(data))
		unix.Close(fd)
	}
	{
		fd, err := unix.Open(tmpfile, unix.O_RDWR, 0644)
		require.NoError(t, err)
		data, err := unix.Mmap(fd, 0, 1024, unix.PROT_WRITE, unix.MAP_SHARED)
		require.NoError(t, err)
		//data[2] = 'd' // will be panic
		t.Log(string(data[:26]))
		data[2] = 'A'
		//err = unix.Msync(data, 0)
		//err = unix.Munmap(data)
		require.NoError(t, err)
		t.Log(len(data), cap(data), &data[0])
		unix.Close(fd)

	}
}
