package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var conn *rados.Conn

const (
	poolName     = "jiapool"
	radosBusy    = -16
	radosExisted = -17
	radosEnoent  = -2
)

func cephInitByFile(conn *rados.Conn, dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if info.IsDir() == false {
			return conn.ReadConfigFile(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func newConn() *rados.Conn {
	conn, err := rados.NewConn()
	if err != nil {
		panic(err)
	}
	cephInitByFile(conn, "k8s_conf")
	err = conn.Connect()
	if err != nil {
		panic(err)
	}
	return conn
}

func init() {
	conn = newConn()
}

func TestListPool(t *testing.T) {
	pools, err := conn.ListPools()
	require.NoError(t, err)
	for _, pool := range pools {
		t.Log(pool)
	}
}

func TestGetSet(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
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
	t.Log(string(bytesOut))

	stat, err := ioctx.Stat("benjamin_test")
	assert.NoError(t, err)
	t.Log(stat.ModTime, stat.Size)

	_, err = ioctx.Stat("benjamin_test_2")
	if err == rados.RadosErrorNotFound {
		t.Log("not found:", err)
	}
}

func TestXattr(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	assert.NoError(t, err)
	// write some data
	bytesIn := []byte("input attr data")
	objectID := "attr_test"
	err = ioctx.Write(objectID, bytesIn, 0)
	assert.NoError(t, err)
	// write xattr
	attrIn := []byte("{1:2}")
	err = ioctx.SetXattr(objectID, "test_attr", attrIn)
	assert.NoError(t, err)
	// read the data back out
	bytesOut := make([]byte, len(bytesIn))
	n, err := ioctx.Read(objectID, bytesOut, 0)
	assert.NoError(t, err)
	assert.Equal(t, bytesIn, bytesOut)
	t.Log(n, string(bytesOut))
	attrOut := make([]byte, len(attrIn))
	// read xattr
	_, err = ioctx.GetXattr(objectID, "test_attr", attrOut)
	assert.NoError(t, err)
	assert.Equal(t, attrIn, attrOut)
	t.Log(string(attrOut))
}

func TestOmap(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	// write some data
	bytesIn := []byte("input omap data")
	objectID := "omap_test"
	err = ioctx.Write(objectID, bytesIn, 0)
	assert.NoError(t, err)
	// write omap
	omapIn := map[string][]byte{
		"key1": []byte("data1"),
		"key2": []byte("data2"),
	}
	err = ioctx.SetOmap(objectID, omapIn)
	assert.NoError(t, err)
	// read the data back out
	bytesOut := make([]byte, len(bytesIn))
	_, err = ioctx.Read(objectID, bytesOut, 0)
	assert.NoError(t, err)
	assert.Equal(t, bytesIn, bytesOut)
	t.Log(string(bytesOut))
	// read omap
	omapOut, err := ioctx.GetOmapValues(objectID, "", "", 100)
	assert.NoError(t, err)
	for key, val := range omapOut {
		assert.Equal(t, val, omapIn[key])
		t.Log(key, string(val))
	}
}

func TestListObject(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	i := 0
	err = ioctx.ListObjects(func(oid string) {
		if i > 100 {
			return
		}
		i++
		t.Log(oid)
	})
	assert.NoError(t, err)
}

func TestIter(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)

	iter, err := ioctx.Iter()
	require.NoError(t, err)
	for i := 0; i < 1000000 && iter.Next() != false; i++ {
		if iter.Value() == "omap_test" {
			t.Log("ns", iter.Namespace())
			t.Log("val", iter.Value())
			break
		}
	}
}

func TestWriteBigData(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)

	data := make([]byte, 1024*1024*50)
	_, err = rand.Read(data)
	require.NoError(t, err)
	for i := 0; i < 2; i++ {
		err = ioctx.Write("big_data"+fmt.Sprint(i), data, 0)
		require.NoError(t, err)
	}
}

func TestWriteFull(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	oid := "write_full_example"

	err = ioctx.WriteFull(oid, []byte("123"))
	require.NoError(t, err)
	err = ioctx.WriteFull(oid, []byte("2345"))
	require.NoError(t, err)

	data := make([]byte, 4)
	_, err = ioctx.Read(oid, data, 0)
	require.NoError(t, err)
	t.Log(string(data))
	assert.Equal(t, string(data), "2345")
}

func TestLockExclusive(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	oid := "lock_exclusive_example"
	oidLock := "lock"
	cookie := "cookie"
	maxExpired := time.Minute

	res, err := ioctx.LockExclusive(oid, oidLock, cookie, "lock test", maxExpired, nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	// verify lock ex
	info, err := ioctx.ListLockers(oid, oidLock)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(info.Clients))
	assert.Equal(t, true, info.Exclusive)

	// fail to lock ex again, lock is busy
	res, err = ioctx.LockExclusive(oid, oidLock, cookie+"2", "this is a description", time.Millisecond, nil)
	assert.NoError(t, err)
	assert.Equal(t, radosBusy, res)

	// unlock
	res, err = ioctx.Unlock(oid, oidLock, cookie)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
}

func TestLockExclusiveDifferentCtx(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	oid := "lock_exclusive_example"
	oidLock := "lock"
	cookie := "cookie"
	maxExpired := time.Minute

	res, err := ioctx.LockExclusive(oid, oidLock, cookie, "lock test", maxExpired, nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	// verify lock ex
	info, err := ioctx.ListLockers(oid, oidLock)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(info.Clients))
	t.Log("clients", info.Clients)
	assert.Equal(t, true, info.Exclusive)

	{
		// check busy with same cookie must in different client
		conn := newConn()
		defer conn.Shutdown()
		// open a pool handle
		ioctx, err := conn.OpenIOContext(poolName)
		require.NoError(t, err)
		defer ioctx.Destroy()
		// fail to lock ex again, lock exist
		res, err = ioctx.LockExclusive(oid, oidLock, cookie, "this is a description", time.Millisecond, nil)
		assert.NoError(t, err)
		assert.Equal(t, radosBusy, res)
	}
	// unlock
	res, err = ioctx.Unlock(oid, oidLock, cookie)
	assert.NoError(t, err)
	if err != nil {
		info, err := ioctx.ListLockers(oid, oidLock)
		require.NoError(t, err)
		for _, client := range info.Clients {
			res, err := ioctx.BreakLock(oid, oidLock, client, cookie)
			require.NoError(t, err)
			require.Equal(t, 0, res)
		}
	}
	assert.Equal(t, 0, res)
}

func TestLockShared(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)
	oid := "lock_shared_example"
	oidLock := "lock"
	escCookie := "escCookie"
	cookie := "cookie"
	tag := "read"
	diffTag := "write"
	maxExpired := time.Minute

	res, err := ioctx.LockShared(oid, oidLock, cookie, tag, "lock test", maxExpired, nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	res, err = ioctx.LockShared(oid, oidLock, escCookie, tag, "lock test", maxExpired, nil)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	res, err = ioctx.LockShared(oid, oidLock, escCookie+"2", diffTag, "lock test", time.Microsecond, nil)
	assert.NoError(t, err)
	assert.Equal(t, -16, res)

	res, err = ioctx.Unlock(oid, oidLock, cookie)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)

	res, err = ioctx.Unlock(oid, oidLock, escCookie)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
}

type ReadWriteLock struct {
	Oid          string
	Ctx          *rados.IOContext
	sharedCookie string
}

func (l ReadWriteLock) lockName() string {
	return "default"
}

func (l ReadWriteLock) tagName() string {
	return "default"
}

func (l ReadWriteLock) cookie() string {
	return l.sharedCookie
}

func (l ReadWriteLock) RLock() error {
	timeout := time.After(time.Second)
	for {
		select {
		case <-timeout:
			return errors.New("lock timeout")
		default:
			res, err := l.Ctx.LockShared(l.Oid, l.lockName(), l.cookie(), l.tagName(), "rwlock", time.Minute, nil)
			if err := lockErrorProcess(res, err); err != nil {
				return err
			}
			if res == 0 {
				return nil
			}
		}
	}
}

func lockErrorProcess(res int, err error) error {
	if err != nil {
		return err
	}
	if res == radosExisted || res == radosBusy {
		time.Sleep(time.Millisecond * 100)
	} else if res != 0 {
		return errors.New("unknown failure, code=" + fmt.Sprint(res))
	}
	return nil
}

func unlockErrorProcess(res int, err error) error {
	if err != nil {
		return err
	}
	if res == radosEnoent {
		return errors.New("unlock not existed response")
	}
	if res != 0 {
		return errors.New("unknown failure, code=" + fmt.Sprint(res))
	}
	return nil
}

func (l ReadWriteLock) RUnLock() error {
	res, err := l.Ctx.Unlock(l.Oid, l.lockName(), l.cookie())
	return unlockErrorProcess(res, err)
}

func (l ReadWriteLock) Lock() error {
	timeout := time.After(time.Second)
	for {
		select {
		case <-timeout:
			return errors.New("lock timeout")
		default:
			res, err := l.Ctx.LockExclusive(l.Oid, l.lockName(), l.cookie(), "rwlock", time.Minute, nil)
			if err := lockErrorProcess(res, err); err != nil {
				return err
			}
			if res == 0 {
				return nil
			}
		}
	}
}

func (l ReadWriteLock) UnLock() error {
	res, err := l.Ctx.Unlock(l.Oid, l.lockName(), l.cookie())
	return unlockErrorProcess(res, err)
}

func TestRWLock(t *testing.T) {
	// open a pool handle
	ioctx, err := conn.OpenIOContext(poolName)
	require.NoError(t, err)

	lock := ReadWriteLock{
		Oid:          "test_rw_lock",
		Ctx:          ioctx,
		sharedCookie: "test_rw_lock_cookie",
	}
	err = lock.Lock()
	assert.NoError(t, err)

	err = lock.RLock()
	assert.Error(t, err)
	t.Log(err)

	err = lock.UnLock()
	assert.NoError(t, err)

}
