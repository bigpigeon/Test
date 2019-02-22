/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"fmt"
	pb "github.com/bigpigeon/Test/go/grpc/helloworld"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		span.LogFields(log.String("user def", "test data"))
	}
	fmt.Println("append : ", ctx.Value("append"))
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	//th := handlers.NewTraceHandler(handlers.NewTrace())
	//s := grpc.NewServer(grpc.StatsHandler(th))
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		panic(cerr)
	}
	r := &etcdnaming.GRPCResolver{Client: cli}
	err = r.Update(context.Background(), "my-server", naming.Update{Op: naming.Add, Addr: "localhost", Metadata: "..."})
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
