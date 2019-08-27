/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package test_lab

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	// 创建一个gin的http服务
	g := gin.New()
	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// 使用httptest直接调用http服务
	recorder := httptest.NewRecorder()
	g.ServeHTTP(recorder, httptest.NewRequest("GET", "/test", nil))
	t.Log(recorder.Code)
	t.Log(recorder.Body.String())
}
