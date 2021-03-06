package main

import (
	"bytes"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"

	"time"
)

var otherTrace opentracing.Tracer

func main() {

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:                    "const",
			Param:                   1,
			SamplingRefreshInterval: 1 * time.Second,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,

			//LocalAgentHostPort:  "192.168.0.1:6831",
		},
		ServiceName: "test ",
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	var otherCloser io.Closer
	otherTrace, otherCloser, err = cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(err)
	}
	defer otherCloser.Close()

	opentracing.SetGlobalTracer(tracer)
	//defer closer.Close()
	someFunction("hello")
	someFunction("fuck")
	//time.Sleep(5 * time.Second)

}

func someFunction(name string) {
	parent := opentracing.GlobalTracer().StartSpan(name, opentracing.ChildOf(nil))

	parent.Finish()
	ctx := parent.Context().(jaeger.SpanContext)
	fmt.Printf("parent trace %s parent %s span %s\n", ctx.TraceID(), ctx.ParentID(), ctx.SpanID())

	{
		buff := bytes.Buffer{}
		err := opentracing.GlobalTracer().Inject(parent.Context(), opentracing.Binary, &buff)
		if err != nil {
			panic(err)
		}
		fmt.Printf("buff %v\n", buff.String())

		httpBuff := opentracing.HTTPHeadersCarrier{}
		err = opentracing.GlobalTracer().Inject(parent.Context(), opentracing.HTTPHeaders, &httpBuff)
		if err != nil {
			panic(err)
		}
		fmt.Printf("http header buff %v\n", httpBuff)

		{

			newSpan, err := opentracing.GlobalTracer().Extract(opentracing.Binary, &buff)
			if err != nil {
				panic(err)
			}

			child := otherTrace.StartSpan(
				"world", opentracing.ChildOf(newSpan))
			child.LogFields(log.String("level", "info"), log.String("event", "abcd"))
			child.LogFields(log.String("level", "info2"), log.String("event", "abcd"))
			ctx := child.Context().(jaeger.SpanContext)

			fmt.Printf("child trace %s parent %s span %s\n", ctx.TraceID(), ctx.ParentID(), ctx.SpanID())
			child.SetTag("error", true)

			defer child.Finish()
		}
		//{
		//	newSpan, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, &opentracing.HTTPHeadersCarrier{})
		//	if err != nil {
		//		fmt.Printf("err %T\n", err)
		//		panic(err)
		//	}
		//	child := opentracing.GlobalTracer().StartSpan(
		//		"world http", opentracing.ChildOf(newSpan))
		//	child.LogFields(log.String("level", "info"), log.String("event", "abcd"))
		//	child.LogFields(log.String("level", "info2"), log.String("event", "abcd"))
		//
		//	child.SetTag("error", true)
		//
		//	defer child.Finish()
		//}
	}

}
