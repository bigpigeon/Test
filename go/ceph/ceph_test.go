package main

import (
	"github.com/ceph/go-ceph/rados"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSet(t *testing.T) {
	conn, err := rados.NewConn()
	assert.NoError(t, err)
	err = conn.ReadConfigFile("conf/ceph.client.admin.keyring")
	assert.NoError(t, err)
	err = conn.ReadConfigFile("conf/ceph.mon.keyring")
	assert.NoError(t, err)
	err = conn.ReadConfigFile("conf/ceph.conf")
	//err = conn.ReadDefaultConfigFile()
	assert.NoError(t, err)
	//args := []string{"--mon-host", "192.168.0.81"}
	//err = conn.ParseCmdLineArgs(args)

	assert.NoError(t, err)
	conn.Connect()
	// open a pool handle
	ioctx, err := conn.OpenIOContext("test")
	assert.NoError(t, err)
	// write some data
	bytesIn := []byte("input data")
	err = ioctx.Write("benjamin_test", bytesIn, 0)
	assert.NoError(t, err)
	// read the data back out
	bytesOut := make([]byte, len(bytesIn))
	_, err = ioctx.Read("benjamin_test", bytesOut, 0)
	assert.NoError(t, err)
	assert.Equal(t, bytesIn, bytesOut)
}
