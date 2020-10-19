/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package rocketmq_test

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/stretchr/testify/require"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func SendN(t testing.TB, p rocketmq.Producer, topic string, n int) {
	SendTagN(t, p, topic, n, "")
}

func SendTagN(t testing.TB, p rocketmq.Producer, topic string, n int, tag string) {
	for i := 0; i < n; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client!" + strconv.Itoa(i)),
		}
		msg.WithTag(tag)
		_, err := p.SendSync(context.Background(), msg)
		require.NoError(t, err)
		//t.Logf("%#v\n", rst)
	}
}

func TestProduce(t *testing.T) {
	endpoint := primitive.NamesrvAddr{"http://localhost:9876"}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		//producer.WithNsResovler(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(1),
		producer.WithGroupName("GID_xxxxxx"),
	)

	require.NoError(t, err)
	err = p.Start()
	require.NoError(t, err)
	defer p.Shutdown()
	SendN(t, p, "test", 10000)

}

func TestSlowFastProcess(t *testing.T) {
	endpoint := primitive.NamesrvAddr{"http://localhost:9876"}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		//producer.WithNsResovler(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(1),
		producer.WithGroupName("GID_xxxxxx"),
	)
	require.NoError(t, err)
	err = p.Start()
	require.NoError(t, err)
	defer p.Shutdown()
	SendN(t, p, "test", 10000)
	// slow consumer
	{
		c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(endpoint),
			consumer.WithConsumerModel(consumer.Clustering),
			consumer.WithGroupName("GID_XXXXXX"),
			consumer.WithConsumeMessageBatchMaxSize(512),
			consumer.WithPullBatchSize(40),
		)
		require.NoError(t, err)
		defer c.Shutdown()

		count := int64(0)
		go func() {
			for {
				time.Sleep(1 * time.Second)
				fmt.Printf("process count %d\n", count)
			}
		}()
		err = c.Subscribe("test", consumer.MessageSelector{},
			func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				for i, msg := range ext {
					fmt.Printf("subscribe %d callback: %v id: %s \n", i, string(msg.Body), msg.MsgId)
				}
				//atomic.AddInt64(&count, 1)
				if cc := atomic.AddInt64(&count, 1); cc%2 != 1 {
					//time.Sleep(100 * time.Second)
				}

				return consumer.ConsumeSuccess, nil
			})
		require.NoError(t, err)
		err = c.Start()
		require.NoError(t, err)
	}
	{
		c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(endpoint),
			consumer.WithConsumerModel(consumer.Clustering),
			consumer.WithGroupName("GID_XXXXXX1"),
			consumer.WithConsumeMessageBatchMaxSize(1),
			consumer.WithPullBatchSize(40),
		)
		require.NoError(t, err)
		defer c.Shutdown()

		count := int64(0)
		go func() {
			for {
				time.Sleep(1 * time.Second)
				fmt.Printf("1 process count %d\n", count)
			}
		}()
		err = c.Subscribe("test", consumer.MessageSelector{},
			func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				for i, msg := range ext {
					fmt.Printf("subscribe %d callback: %v id: %s \n", i, string(msg.Body), msg.MsgId)
				}
				//atomic.AddInt64(&count, 1)
				if cc := atomic.AddInt64(&count, 1); cc%2 != 1 {
					//time.Sleep(100 * time.Second)
				}

				return consumer.ConsumeSuccess, nil
			})
		require.NoError(t, err)
		err = c.Start()
		require.NoError(t, err)
	}
	time.Sleep(10000 * time.Second)
}

func defaultConsumer(ep []string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(ep),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_XXXXXX"),
		consumer.WithConsumeMessageBatchMaxSize(1),
		consumer.WithSuspendCurrentQueueTimeMillis(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return c
}

func defaultSubScribe(t *testing.T, c rocketmq.PushConsumer, processor int) {
	count := int64(0)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d process count %d\n", processor, count)
		}
	}()
	err := c.Subscribe("test", consumer.MessageSelector{},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i, msg := range ext {
				fmt.Printf("subscribe %d callback: %v id: %s \n", i, string(msg.Body), msg.MsgId)
				if cc := atomic.AddInt64(&count, 1); cc%2 != 1 {
					time.Sleep(100 * time.Second)
				}
			}

			return consumer.ConsumeSuccess, nil
		})
	require.NoError(t, err)
}

func testFuck(t *testing.T) {
	endpoint := primitive.NamesrvAddr{"http://localhost:9876"}

	for processor := 0; processor < 1; processor++ {
		c := defaultConsumer(endpoint)

		//waitStop := make(chan struct{})
		defaultSubScribe(t, c, processor)
		err := c.Start()

		require.NoError(t, err)
	}
	time.Sleep(10000 * time.Second)
}

func TestFuck(t *testing.T) {
	testFuck(t)
}

func TestFuck2(t *testing.T) {
	testFuck(t)
}

func TestTwoConsumer(t *testing.T) {

}

func TestReScribe(t *testing.T) {
	endpoint := primitive.NamesrvAddr{"http://localhost:9876"}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		//producer.WithNsResovler(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(1),
		producer.WithGroupName("GID_xxxxxx"),
	)

	require.NoError(t, err)
	err = p.Start()
	require.NoError(t, err)
	defer p.Shutdown()
	SendN(t, p, "test", 10000)

	c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(endpoint),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_XXXXXX"),
		consumer.WithConsumeMessageBatchMaxSize(1),
	)
	require.NoError(t, err)
	for ci := 0; ci < 2; ci++ {
		ciCopy := ci
		//waitStop := make(chan struct{})
		count := int64(0)
		go func() {
			for {
				time.Sleep(1 * time.Second)
				fmt.Printf("%d process count %d\n", ciCopy, count)
			}
		}()
		err = c.Subscribe("test", consumer.MessageSelector{},
			func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
				for i, msg := range ext {
					fmt.Printf("subscribe %d callback: %v id: %s \n", i, string(msg.Body), msg.MsgId)
				}
				//atomic.AddInt64(&count, 1)
				if cc := atomic.AddInt64(&count, 1); cc%2 != 1 {
					//time.Sleep(100 * time.Second)
				}

				return consumer.ConsumeSuccess, nil
			})
		require.NoError(t, err)
	}
	err = c.Start()
	require.NoError(t, err)
	time.Sleep(10000 * time.Second)
}

func TestTag(t *testing.T) {
	endpoint := primitive.NamesrvAddr{"http://localhost:9876"}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		//producer.WithNsResovler(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(1),
		producer.WithGroupName("GID_xxxxxx"),
	)

	require.NoError(t, err)
	err = p.Start()
	require.NoError(t, err)
	defer p.Shutdown()
	SendTagN(t, p, "test", 1, "aa")

	c, err := rocketmq.NewPushConsumer(consumer.WithNameServer(endpoint),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_XXXXXX"),
		consumer.WithConsumeMessageBatchMaxSize(1),
	)
	require.NoError(t, err)
	count := int64(0)
	err = c.Subscribe("test", consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "aa || bb",
	},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i, msg := range ext {
				fmt.Printf("subscribe %d callback: %v id: %s \n", i, string(msg.Body), msg.MsgId)
			}
			//atomic.AddInt64(&count, 1)
			if cc := atomic.AddInt64(&count, 1); cc%2 != 1 {
				//time.Sleep(100 * time.Second)
			}

			return consumer.ConsumeSuccess, nil
		})
	require.NoError(t, err)
	err = c.Start()
	require.NoError(t, err)
	defer c.Shutdown()
	time.Sleep(10 * time.Second)
}
