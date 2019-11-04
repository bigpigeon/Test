/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package grpc_noclose

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime"
	"testing"
)

type ecServer struct {
	addr string
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}
func (s *ecServer) ServerStreamingEcho(*ecpb.EchoRequest, ecpb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) ClientStreamingEcho(ecpb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) BidirectionalStreamingEcho(ecpb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ecpb.RegisterEchoServer(s, &ecServer{addr: addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Conn struct{ *grpc.ClientConn }

func TestGrpcNoClose(t *testing.T) {
	addr := "localhost:50051"
	//go startServer(addr)

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	t.Log("total", stats.TotalAlloc, "heap", stats.HeapAlloc)
	for i := 0; i < 5000; i++ {
		conn, err := grpc.DialContext(context.Background(), addr, grpc.WithInsecure())
		require.NoError(t, err)
		cli := ecpb.NewEchoClient(conn)
		runtime.SetFinalizer(cli, func(cli ecpb.EchoClient) {
			fmt.Println("lease once")
			conn.Close()
		})
		_, err = cli.UnaryEcho(context.Background(), &ecpb.EchoRequest{Message: "abc"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}

	runtime.GC()
	t.Log("total", stats.TotalAlloc, "heap", stats.HeapAlloc)
	//time.Sleep(time.Second * 10000)
}
