/*
 * Copyright 2018 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"time"
)

func EchoServer(ws *websocket.Conn) {
	var err error
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
	//buff := bytes.Buffer{}
	//_, err = io.Copy(&buff, ws)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("logger ", buff.String())
	_, err = ws.Write([]byte("test"))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	_, err = ws.Write([]byte("test 2"))
	if err != nil {
		panic(err)
	}
	//ws.WriteClose(400)
}

func main() {
	r := gin.New()
	r.GET("/ws", func(c *gin.Context) {
		handler := websocket.Handler(EchoServer)
		handler.ServeHTTP(c.Writer, c.Request)
		c.JSON(400, gin.H{"error": "test"})
	})
	r.Run(":12345")
}
