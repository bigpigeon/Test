/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package tls

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	pb "github.com/bigpigeon/Test/go/grpc/helloworld"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"
)

type ServerImpl struct{}

func (s ServerImpl) SayHello(c context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello " + req.Name,
	}, nil
}

func TestTls(t *testing.T) {
	rootCAs, err := x509.SystemCertPool()
	require.NoError(t, err)
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	caCrtData, err := ioutil.ReadFile("ca.pem")
	require.NoError(t, err)
	if ok := rootCAs.AppendCertsFromPEM(caCrtData); !ok {
		t.Log("add cert failure")
		t.FailNow()
	}
	cliCerts := credentials.NewClientTLSFromCert(rootCAs, "")

	creds, err := credentials.NewServerTLSFromFile("local.pem", "local-key.pem")
	require.NoError(t, err)
	lis, err := net.Listen("tcp", "localhost:50051")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &ServerImpl{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	defer wg.Wait()
	defer s.Stop()
	{
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(cliCerts))
		require.NoError(t, err)
		// error handling omitted
		client := pb.NewGreeterClient(conn)
		res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "pigeon"})
		require.NoError(t, err)
		t.Log(res.Message)
	}
	{
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(cliCerts))
		require.NoError(t, err)
		// error handling omitted
		client := pb.NewGreeterClient(conn)
		_, err = client.SayHello(context.Background(), &pb.HelloRequest{Name: "pigeon"})
		require.Error(t, err)
		t.Log(err)
	}
}

func TestToken(t *testing.T) {
	rootCAs, err := x509.SystemCertPool()
	require.NoError(t, err)
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	caCrtData, err := ioutil.ReadFile("ca.pem")
	require.NoError(t, err)
	if ok := rootCAs.AppendCertsFromPEM(caCrtData); !ok {
		t.Log("add cert failure")
		t.FailNow()
	}
	cliCerts := credentials.NewClientTLSFromCert(rootCAs, "localhost")

	creds, err := credentials.NewServerTLSFromFile("local.pem", "local-key.pem")
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken:  "1123",
		TokenType:    "basic",
		RefreshToken: "132",
		Expiry:       time.Now().Add(-1 * time.Second),
	}
	token = token.WithExtra("extra")

	access := oauth.NewOauthAccess(token)

	lis, err := net.Listen("tcp", "localhost:50051")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok == false {
			return nil, errors.New("not metadata")
		}
		fmt.Println(md)
		return handler(ctx, req)
	}), grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &ServerImpl{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	defer wg.Wait()
	defer s.Stop()
	{
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithPerRPCCredentials(access), grpc.WithTransportCredentials(cliCerts))
		require.NoError(t, err)
		// error handling omitted
		client := pb.NewGreeterClient(conn)
		res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "pigeon"})
		require.NoError(t, err)
		t.Log(res.Message)
	}
}
