/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package http_test

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
	"time"
)

func TestHttpNoClose(t *testing.T) {
	e := gin.New()
	e.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	go func() {
		e.Run(":50001")
	}()
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10000; i++ {
		req, err := http.NewRequest("GET", "http://localhost:50001/ping", nil)
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.Equal(t, res.StatusCode, 200)
	}
	debug.FreeOSMemory()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	t.Log(mem.HeapAlloc, mem.TotalAlloc)
}

func TestHttpGcClose(t *testing.T) {
	e := gin.New()
	e.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	go func() {
		e.Run(":50001")
	}()
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10000; i++ {
		func() {

			req, err := http.NewRequest("GET", "http://localhost:50001/ping", nil)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.Equal(t, res.StatusCode, 200)
			runtime.SetFinalizer(res, func(res *http.Response) {
				res.Body.Close()
			})
		}()
	}
	debug.FreeOSMemory()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	t.Log(mem.HeapAlloc, mem.TotalAlloc)
}

func TestHttpClose(t *testing.T) {
	e := gin.New()
	e.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "pong")
	})
	go func() {
		e.Run(":50001")
	}()
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10000; i++ {
		req, err := http.NewRequest("GET", "http://localhost:50001/ping", nil)
		require.NoError(t, err)
		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.Equal(t, res.StatusCode, 200)
		res.Body.Close()
	}
	debug.FreeOSMemory()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	t.Log(mem.HeapAlloc, mem.TotalAlloc)
}

func TestReadResponse(t *testing.T) {
	reader := strings.NewReader(`HTTP/1.1 200 OK
Date: Tue, 09 Jun 2020 08:42:34 GMT
Server: Apache
Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
ETag: "51-47cf7e6ee8400"
Accept-Ranges: bytes
Content-Length: 81
Cache-Control: max-age=86400
Expires: Wed, 10 Jun 2020 08:42:34 GMT
Connection: Keep-Alive
Content-Type: text/html

<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
`)
	buf := bufio.NewReader(reader)
	resp, err := http.ReadResponse(buf, nil)
	require.NoError(t, err)
	t.Log(resp.Header)
	//d, err := ioutil.ReadAll(resp.Body)
	//require.NoError(t, err)
	//t.Log(string(d))
	t.Log(buf.ReadString('\n'))
}

func TestRequestWrite(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < 2; i++ {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:1313", nil)
		require.NoError(t, err)
		err = req.Write(&buf)
		require.NoError(t, err)
	}
	t.Log(buf.String())
}

func TestChangeHost(t *testing.T) {
	req, err := http.NewRequest("GET", "http://192.168.0.1/path", nil)
	require.NoError(t, err)
	req.RequestURI = ""
	uri, err := url.Parse("http://baidu.com")
	require.NoError(t, err)
	req.URL.Host = uri.Host
	req.URL.Scheme = uri.Scheme
	//req.URL.Host = "baidu.com"
	req.Host = ""

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	data, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(data))
}
