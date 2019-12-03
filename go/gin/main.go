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

type TestQuerySub struct {
	SubData string `form:"sub_data"`
}

type TestQuery struct {
	Name    string       `form:"name"`
	Address string       `form:"address"`
	Sub     TestQuerySub `form:"sub"`
}

func main() {
	r1 := gin.Default()
	//data := TestData{
	//	IntVal: 6446744073709551610,
	//}

	r1.GET("/abc", func(c *gin.Context) {
		var query TestQuery
		err := c.BindQuery(&query)
		if err != nil {
			panic(err)
		}
		c.JSON(200, query)
	})
	r1.Run()

}
