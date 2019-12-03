/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/grpc/handlers"
	"github.com/bigpigeon/Test/go/grpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	logger "log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func runOnce() {
	// Set up a connection to the server.
	gCtx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
	trace := handlers.NewTrace()
	th := handlers.NewTraceHandler(trace)
	conn, err := grpc.DialContext(gCtx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithStatsHandler(th))
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}

	r, err := c.SayHello(gCtx, &helloworld.HelloRequest{
		Name: "1",
	})
	fmt.Println("??")
	if err != nil {
		logger.Fatalf("could not greet: %v", err)
	}
	logger.Printf("Greeting: %s", r.Message)
}

func main() {
	runOnce()
}
