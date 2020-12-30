/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package net

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestNetIPContainer(t *testing.T) {
	_, ip1, err := net.ParseCIDR("192.168.0.241/24")
	require.NoError(t, err)
	_, ip2, err := net.ParseCIDR("192.168.1.24/24")
	require.NoError(t, err)
	t.Log(ip1.Contains(ip2.IP))
	fmt.Println(uint32(ip2.IP[0]) << 24)
}

func TestListen(t *testing.T) {
	li, err := net.Listen("tcp", "localhost:10000")
	require.NoError(t, err)
	t.Log(li.Addr())
	{
		li, err := net.Listen("tcp", "localhost:10000")
		require.NoError(t, err)
		t.Log(li.Addr())
	}
}
