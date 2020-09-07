/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	"github.com/bigpigeon/Test/go/grpc/error_example/options"
	"google.golang.org/grpc/status"
	"log"
	"net"

	api "github.com/bigpigeon/Test/go/grpc/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	if len(req.Name) >= 10 {
		return nil, status.Error(options.TooLong, "Too Long")
	}
	return &api.HelloReply{Message: fmt.Sprintf("Hey, %s!", req.GetName())}, nil
}

func Serve() {
	addr := fmt.Sprintf(":%d", 50051)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen to address %s", addr)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &HelloServer{})
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	Serve()
}
