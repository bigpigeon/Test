/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package pubsub_test

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

const addr = "127.0.0.1:4150"
const lookupAddr = "127.0.0.1:4161"
const topic = "test"

func testPub(t *testing.T, n int) {
	w, err := nsq.NewProducer(addr, nsq.NewConfig())
	require.NoError(t, err)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		err := w.Publish(topic, []byte("data "+strconv.Itoa(i)))
		if err != nil {
			log.Panic("Could not connect")
		}
	}
	w.Stop()
}

func testSub(t *testing.T, channel string) *nsq.Consumer {
	q, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	require.NoError(t, err)

	n := int64(0)
	go func() {
		for {
			fmt.Printf("processd %d\n", n)
			time.Sleep(1 * time.Second)
		}
	}()
	q.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		message.Touch()

		//time.Sleep(time.Millisecond * (time.Duration(rand.Intn(1000) + 1000)))
		message.Finish()
		log.Printf("c %s p %d m: %v has res: %v \n", channel, n, string(message.Body), message.HasResponded())
		atomic.AddInt64(&n, 1)
		//if cn%4 != 0 {
		//	time.Sleep(1 * time.Second)
		//}
		return nil
	}), 4)
	err = q.ConnectToNSQLookupd(lookupAddr)
	return q
}

func TestPubSub(t *testing.T) {
	testPub(t, 10000)
	q := testSub(t, "1")
	<-q.StopChan
}

func TestPub(t *testing.T) {
	testPub(t, 10000)
}

func TestSub(t *testing.T) {
	q := testSub(t, "1")
	<-q.StopChan
}

func TestSub2(t *testing.T) {
	q := testSub(t, "1")
	<-q.StopChan
}
