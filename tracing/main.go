package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerTransport "github.com/uber/jaeger-client-go/transport"
	jaegerPrometheus "github.com/uber/jaeger-lib/metrics/prometheus"
)

// initJaeger 将jaeger tracer设置为全局tracer
func initJaeger(service, endpoint string) (opentracing.Tracer, io.Closer) {
	cfg := jaegercfg.Configuration{
		// 将采样频率设置为1，每一个span都记录，方便查看测试结果
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址
			CollectorEndpoint: endpoint,
		},
	}
	closer, err := cfg.InitGlobalTracer(service, jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	// 获取jaeger tracer
	return opentracing.GlobalTracer(), closer
}

func initProbabilisticJaeger(service, endpoint string) (opentracing.Tracer, io.Closer) {
	sampler, err := jaeger.NewProbabilisticSampler(0.5)
	if err != nil {
		panic(err)
	}
	jaegerTracer, jaegerCloser, err := jaegercfg.Configuration{
		ServiceName: service,
	}.NewTracer(jaegercfg.Sampler(sampler),
		jaegercfg.Reporter(jaeger.NewRemoteReporter(jaegerTransport.NewHTTPTransport(endpoint))),
		jaegercfg.Metrics(jaegerPrometheus.New()),
	)
	if err != nil {
		panic(err)
	}
	return jaegerTracer, jaegerCloser
}

func initNullJaeger(service, endpoint string) (opentracing.Tracer, io.Closer) {
	return jaeger.NewTracer(service, jaeger.NewConstSampler(false), jaeger.NewNullReporter())
}

func main() {
	//t, closer := initJaeger("in-process", "http://localhost:14268/api/traces")
	t, closer := initProbabilisticJaeger("in-process", "http://localhost:14268/api/traces")
	defer closer.Close()
	// 创建root span
	sp := t.StartSpan("in-process-service")
	// main执行完结束这个span
	defer sp.Finish()
	// 将span传递给Foo
	ctx := opentracing.ContextWithSpan(context.Background(), sp)
	Foo(ctx)
}

func Foo(ctx context.Context) {
	// 开始一个span, 设置span的operation_name=Foo
	span, ctx := opentracing.StartSpanFromContext(ctx, "Foo")
	defer span.Finish()
	// 将context传递给Bar
	Bar(ctx)
	// 模拟执行耗时
	time.Sleep(1 * time.Second)
}
func Bar(ctx context.Context) {
	// 开始一个span，设置span的operation_name=Bar
	span, ctx := opentracing.StartSpanFromContext(ctx, "Bar")
	defer span.Finish()
	// 模拟执行耗时
	time.Sleep(2 * time.Second)

	// 假设Bar发生了某些错误
	err := errors.New("something wrong")
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)
	span.SetTag("error", true)
}
