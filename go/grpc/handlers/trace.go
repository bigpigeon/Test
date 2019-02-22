/*
 * Copyright 2019 bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package handlers

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc/stats"
	"time"
)

var GRPCComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "gRPC"}

type TraceHandler struct {
	tracer opentracing.Tracer
}

func NewTrace() opentracing.Tracer {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:                    "const",
			Param:                   1,
			SamplingRefreshInterval: 1 * time.Second,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:6831",
		},
		ServiceName: "test ",
	}
	tracer, _, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(err)
	}
	return tracer
}

// NewTraceHandler creates a gRPC stats.Handler instance that instruments RPCs with Opentracing trace contexts
func NewTraceHandler(tracer opentracing.Tracer) *TraceHandler {
	return &TraceHandler{
		tracer: tracer,
	}

}

const (
	EventKey   = "event"
	PayloadKey = "payload"
)

func (th *TraceHandler) TagRPC(ctx context.Context, tagInfo *stats.RPCTagInfo) context.Context {

	spanCtx := extractSpanContext(th.tracer, ctx)
	span := th.tracer.StartSpan(tagInfo.FullMethodName, opentracing.FollowsFrom(spanCtx), GRPCComponentTag)
	newCtx, _ := injectSpanToMetadata(th.tracer, span, ctx)

	return opentracing.ContextWithSpan(newCtx, span)
}

func (th *TraceHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}

	switch t := s.(type) {
	case *stats.Begin:
		span.LogFields(log.String(EventKey, "RPC started"))
	case *stats.InPayload:
		e := log.String(EventKey, fmt.Sprintf("Payload received: Wire length=%d", t.WireLength))
		span.LogFields(e, log.Object(PayloadKey, t.Payload))
	case *stats.InHeader:
		span.LogFields(log.String(EventKey, fmt.Sprintf("Header received: Remote addr=%s, Local addr=%s", t.RemoteAddr, t.LocalAddr)))
	case *stats.InTrailer:
		span.LogFields(log.String(EventKey, "Trailer received"))
	case *stats.OutPayload:
		e := log.String(EventKey, fmt.Sprintf("Payload sent: Wire length=%d", t.WireLength))
		span.LogFields(e, log.Object(PayloadKey, t.Payload))
	case *stats.OutHeader:
		span.LogFields(log.String(EventKey, fmt.Sprintf("Header sent: Remote addr=%s, Local addr=%s", t.RemoteAddr, t.LocalAddr)))
	case *stats.OutTrailer:
		span.LogFields(log.String(EventKey, "Trailer sent"))
	case *stats.End:
		if t.IsClient() {
			span.SetTag(string(ext.SpanKind), ext.SpanKindRPCClientEnum)
		} else {
			span.SetTag(string(ext.SpanKind), ext.SpanKindRPCServerEnum)
		}

		if t.Error != nil {
			span.SetTag(string(ext.Error), true)
			span.LogFields(log.String(EventKey, "RPC failed"), log.Error(t.Error))
		} else {
			span.LogFields(log.String(EventKey, "RPC ended"))
		}
		span.Finish()
	}
}

func (th *TraceHandler) TagConn(ctx context.Context, tagInfo *stats.ConnTagInfo) context.Context {
	return ctx
}

func (th *TraceHandler) HandleConn(ctx context.Context, s stats.ConnStats) {}
