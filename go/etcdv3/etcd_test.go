/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package etcdv3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/namespace"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"sync"
	"testing"
	"time"
)

var cli *clientv3.Client

func init() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"http://localhost:2379"},
	})
	if err != nil {
		panic(err)
	}

}

func TestKV(t *testing.T) {
	ctx := context.Background()
	var lastVersion int64
	{
		rsp, err := cli.Put(ctx, "/test", "val4", clientv3.WithPrevKV())
		assert.NoError(t, err)
		if rsp.PrevKv != nil {
			t.Log("pre kv", string(rsp.PrevKv.Key), string(rsp.PrevKv.Value))
		}
		rsp, err = cli.Put(ctx, "test2/not_root2", "val3")
		assert.NoError(t, err)
		if rsp.PrevKv != nil {
			t.Log(rsp.PrevKv.Key, rsp.PrevKv.Value)
		}
	}
	{
		rsp, err := cli.Get(ctx, "/test")
		assert.NoError(t, err)
		t.Logf("key count %d more %v", rsp.Count, rsp.More)
		for _, kv := range rsp.Kvs {
			t.Logf("key %s val %s version %v Lease %d", kv.Key, kv.Value, kv.Version, kv.Lease)
			lastVersion = kv.Version
		}

	}
	{
		rsp, err := cli.Get(ctx, "/test", clientv3.WithPrefix())
		assert.NoError(t, err)
		for _, kv := range rsp.Kvs {
			t.Logf("key %s val %s version %v", kv.Key, kv.Value, kv.Version)
		}
	}

	{
		rsp, err := cli.Get(ctx, "/test/no_exist", clientv3.WithPrefix())
		assert.NoError(t, err)
		for _, kv := range rsp.Kvs {
			t.Logf("key %s val %s version %v", kv.Key, kv.Value, kv.Version)
		}
	}

	{
		rsp, err := cli.Compact(ctx, lastVersion)
		assert.NoError(t, err)
		t.Logf("header %#v", rsp.Header)
	}

	{
		kv := namespace.NewKV(cli, "/test/")
		_, err := kv.Put(ctx, "namespace", "123")
		assert.NoError(t, err)
		rsp, err := kv.Get(ctx, "namespace")
		assert.NoError(t, err)
		t.Logf("namespace key %s val %s", rsp.Kvs[0].Key, rsp.Kvs[0].Value)
	}
}

func TestLease(t *testing.T) {
	ctx := context.Background()
	lease := clientv3.NewLease(cli)
	rsp, err := lease.Grant(ctx, 10)
	require.NoError(t, err)
	t.Logf("lease id %d", rsp.ID)
	rsp, err = lease.Grant(ctx, 10)
	require.NoError(t, err)
	t.Logf("lease id again %d", rsp.ID)
	rspChan, err := lease.KeepAlive(ctx, rsp.ID)
	require.NoError(t, err)
	{
		for i := 0; i < 1; i++ {
			rsp, more := <-rspChan
			if more == false {
				break
			}
			t.Log("keep alive", rsp.String(), rsp.TTL)
		}
	}

}

func TestWatch(t *testing.T) {
	ctx := context.Background()
	{
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			rsp := cli.Watch(ctx, "/test/watch", clientv3.WithCreatedNotify(), clientv3.WithPrevKV())
			for {
				watchData, ok := <-rsp
				if ok == false {
					break
				}
				fmt.Printf("canceled %v created %v revision %v\n", watchData.Canceled, watchData.Created, watchData.CompactRevision)
				for _, event := range watchData.Events {
					if event.IsCreate() {
						fmt.Printf("key %s val %s  watch event created\n", event.Kv.Key, event.Kv.Value)
					} else if event.Type == mvccpb.DELETE {
						fmt.Printf("key %s val %s  watch del event\n", event.Kv.Key, event.PrevKv.Value)
					} else {
						fmt.Printf("key %s val %s  watch event created %v is modify %v\n", event.Kv.Key, event.Kv.Value, event.IsCreate(), event.IsModify())
					}

				}
			}
		}()
		for i := 1; i < 4; i++ {
			_, err := cli.Put(ctx, "/test/watch", fmt.Sprint(i))
			assert.NoError(t, err)
		}

		_, err := cli.Delete(ctx, "/test/watch")
		assert.NoError(t, err)
		time.Sleep(10 * time.Microsecond)
		cli.Watcher.Close()
		wg.Wait()
	}
}
