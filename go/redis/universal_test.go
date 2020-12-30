/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	cli := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"redis-0.redis:6379", "redis-1.redis", "redis-2.redis"},

		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		PoolSize:     4,
	})

	err := cli.Set("testkey", "123", 100*time.Second).Err()
	require.NoError(t, err)
	data := cli.Get("testkey").Val()
	t.Log(data)
	require.Equal(t, "123", data)
}
