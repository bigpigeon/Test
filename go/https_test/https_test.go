/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package https_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func customTransport(cli *http.Client) {
	dialer := &net.Dialer{
		Timeout:   1 * time.Second,
		KeepAlive: 1 * time.Second,
		DualStack: true,
	}
	cli.Transport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
		if addr == "myhost.me:8080" {
			addr = "127.0.0.1:8080"
		}
		return dialer.DialContext(ctx, network, addr)
	}
}

func TestHttpsCaServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	server := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		err := server.ListenAndServeTLS("cert/my.pem", "cert/my-key.pem")
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	defer server.Close()

	// try use not ca client
	{
		cli := &http.Client{Transport: &http.Transport{}}
		customTransport(cli)
		_, err := cli.Get("https://myhost.me:8080/hello?name=bigpigeon")
		t.Log(err)
		require.Error(t, err)
	}

	rootCAs, err := x509.SystemCertPool()
	require.NoError(t, err)
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	certs, err := ioutil.ReadFile("cert/ca.pem")
	require.NoError(t, err)
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		t.Logf("Failed to append %q to RootCAs: %v", "cert/ca.pem", err)
		t.FailNow()
	}
	config := &tls.Config{
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	_, err = client.Get("https://localhost:8080/hello?name=bigpigeon")
	t.Log(err)
	require.Error(t, err)

	customTransport(client)
	res, err := client.Get("https://myhost.me:8080/hello?name=bigpigeon")
	require.NoError(t, err)
	data, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	t.Log(string(data))
}
