package main

import (
	"fmt"
	"os"

	"github.com/opentracing/opentracing-go/log"
	"github.com/jaegertracing/opentracing-tutorial/go/lib/tracing"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := tracing.Init("hello-world") //生成trace实例，用于开始一个span。参数表示它用于将跟踪程序发出的所有跨距标记为源自hello world服务。
	defer closer.Close()

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")//每一个span都会给一个操作名
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println")

	span.Finish() //每一个span完成必须调用其finish函数。span的开始/结束时间戳会被tracer自动捕获。
}
