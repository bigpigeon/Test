/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package http

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMultipartForm(t *testing.T) {
	engine := gin.New()
	engine.POST("/form", func(ctx *gin.Context) {

	})
}
