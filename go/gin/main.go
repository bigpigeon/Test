/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/gin-gonic/gin"
)

type TestData struct {
	IntVal int64
}

func main() {
	r := gin.Default()
	data := TestData{
		IntVal: 6446744073709551610,
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, data)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
