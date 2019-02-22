/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/grpc/helloworld"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	logger "log"
	"time"
)

const (
	address     = "my-server"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		panic(cerr)
	}
	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)

	conn, err := grpc.DialContext(context.Background(), address, grpc.WithBalancer(b), grpc.WithInsecure())

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

	md := metadata.New(map[string]string{"SpanEncode": "val1", "key2": "val2"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	for {
		r, err := c.SayHello(ctx, &helloworld.HelloRequest{
			Name: "1",
		})
		fmt.Println("??")
		if err != nil {
			logger.Fatalf("could not greet: %v", err)
		}
		logger.Printf("Greeting: %s", r.Message)
		time.Sleep(1 * time.Second)
	}
}
