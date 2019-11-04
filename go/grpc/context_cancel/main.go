/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"context"
	"github.com/bigpigeon/Test/go/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func getConn() *grpc.ClientConn {
	// Set up a connection to the server.
	gCtx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	conn, err := grpc.DialContext(gCtx, address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func client() {

	conn := getConn()
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	for {
		res, err := c.SayHello(context.Background(), &helloworld.HelloRequest{
			Name: "bigpigeon",
		})
		if err != nil {
			log.Println("err", err)
		} else {
			log.Println(res.Message)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	client()
}
