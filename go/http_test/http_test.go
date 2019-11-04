/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package http_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"runtime"
	"runtime/debug"
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
