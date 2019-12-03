/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package dial

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"net/http/httptrace"
	"testing"
	"time"
)

var address = "localhost:50001"

func TestNetDial(t *testing.T) {
	{

		conn, err := net.Dial("tcp", address)
		t.Log(conn)
		t.Log(err)
	}
}

func TestGrpcDial(t *testing.T) {
	{
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		t.Log(conn)
		t.Log(err)
	}
	time.Sleep(100 * time.Second)
}

func TestHttpReq(t *testing.T) {
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
	}

	req, err := http.NewRequestWithContext(httptrace.WithClientTrace(context.Background(), trace), "GET", "http://"+address, nil)
	require.NoError(t, err)
	res, err := http.DefaultClient.Do(req)

	t.Log(err)
	t.Log(res)
}
