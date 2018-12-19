package main

import (
	"bytes"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
			SamplingRefreshInterval: 1 * time.Second,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "192.168.0.15:6831",
		},
		ServiceName: "test ",
	}
	tracer, _, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	//defer closer.Close()
	someFunction("hello")
	someFunction("fuck")
	time.Sleep(5 * time.Second)
}

func someFunction(name string) {
	parent := opentracing.GlobalTracer().StartSpan(name)
	defer parent.Finish()

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
			child := opentracing.GlobalTracer().StartSpan(
				"world", opentracing.ChildOf(newSpan))
			child.LogFields(log.String("level", "info"), log.String("event", "abcd"))
			child.LogFields(log.String("level", "info2"), log.String("event", "abcd"))

			child.SetTag("error", true)

			defer child.Finish()
		}
		{
			newSpan, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, &opentracing.HTTPHeadersCarrier{})
			if err != nil {
				fmt.Printf("err %T\n", err)
				panic(err)
			}
			child := opentracing.GlobalTracer().StartSpan(
				"world http", opentracing.ChildOf(newSpan))
			child.LogFields(log.String("level", "info"), log.String("event", "abcd"))
			child.LogFields(log.String("level", "info2"), log.String("event", "abcd"))

			child.SetTag("error", true)

			defer child.Finish()
		}
	}

}
