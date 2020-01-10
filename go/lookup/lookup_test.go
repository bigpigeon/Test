/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package lookup

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"net"
	"net/http/httptrace"
	"testing"
)

func TestLookup(t *testing.T) {
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Println("dns start", info.Host)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Println("dns done", info.Addrs, info.Coalesced)
			if info.Err != nil {
				fmt.Println("dns done error", info.Err)
			}
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("connection start  ", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("connection done  ", network, addr)
		},
	}
	_default := net.DefaultResolver
	ctx := httptrace.WithClientTrace(context.Background(), trace)
	name, err := _default.LookupIPAddr(ctx, "google.com")
	require.NoError(t, err)
	t.Log(name)
}
