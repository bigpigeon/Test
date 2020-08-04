/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TestData struct {
	IntVal int64
}

type TestQuerySub struct {
	SubData string `form:"sub_data"`
}

type TestQuery struct {
	Name    string       `form:"name"`
	Address string       `form:"address"`
	Sub     TestQuerySub `form:"sub"`
}

type ItemName struct {
	Name string `uri:"name"`
}

func main() {
	r1 := gin.Default()

	r1.GET("/aa", func(ctx *gin.Context) {
		ctx.Next()
		fmt.Println("11")
	}, func(ctx *gin.Context) {
		fmt.Println("22")
	})

	r1.GET("/ab", func(ctx *gin.Context) {
		ctx.Abort()
		fmt.Println("11")
	}, func(ctx *gin.Context) {
		fmt.Println("22")
	})

	r1.GET("/abc", func(c *gin.Context) {
		var query TestQuery
		err := c.BindQuery(&query)
		if err != nil {
			panic(err)
		}
		c.JSON(200, query)
	})

	r1.GET("/item/:name", func(c *gin.Context) {
		var item ItemName
		err := c.BindUri(&item)
		if err != nil {
			panic(err)
		}
		c.JSON(200, "")
	})

	r1.Run()

}
