/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package recover

import (
	"context"
	pb "github.com/bigpigeon/Test/go/grpc/helloworld"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"runtime/debug"
	"testing"
)

// server is used to implement helloworld.GreeterServer.
type Impl struct{}

// SayHello implements helloworld.GreeterServer
func (s *Impl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	panic("cool, boom!!!")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func serverStart(ser *grpc.Server) {
	lis = bufconn.Listen(bufSize)
	go func() {
		if err := ser.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func CustomRecover(p interface{}) (err error) {
	stack := debug.Stack()
	log.Print(string(stack))
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func TestServerRecover(t *testing.T) {

	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(CustomRecover)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(CustomRecover)),
		),
		grpc.StatsHandler(&statHandle{}),
	)
	pb.RegisterGreeterServer(server, &Impl{})
	serverStart(server)
	{
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			t.Log(err)
		}
		cli := pb.NewGreeterClient(conn)
		_, err = cli.SayHello(context.Background(), &pb.HelloRequest{Name: "jia"})
		if err != nil {
			t.Log(err)
		}
	}
}
