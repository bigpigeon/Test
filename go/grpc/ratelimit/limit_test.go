/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package ratelimit

import (
	"context"
	"github.com/bigpigeon/Test/go/grpc/helloworld"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"google.golang.org/grpc"
	"testing"
)

type server struct {
	taskPool chan int
}

func (s *server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {

}

func (s *server) doSomething() {

}

func TestRateLimit(t *testing.T) {
	grpc.NewServer(grpc.UnaryInterceptor(ratelimit.UnaryServerInterceptor(ratelimit.Limiter())))
}
