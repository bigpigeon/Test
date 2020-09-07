/*
 * Copyright 2020 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package recover

import (
	"context"
	"google.golang.org/grpc/stats"
	"log"
)

type statHandle struct{}

func (s statHandle) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (s statHandle) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	switch stat.(type) {
	case *stats.End:
		log.Print("grpc finish")
	}
}

func (s statHandle) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (s statHandle) HandleConn(ctx context.Context, stats stats.ConnStats) {

}
