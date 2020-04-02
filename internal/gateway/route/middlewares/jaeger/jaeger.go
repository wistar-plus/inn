package jaeger

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sp opentracing.Span
		md := make(map[string]string)

		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
			defer sp.Finish()
		} else {
			sp = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer sp.Finish()
		}

		if err = opentracing.GlobalTracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
			fmt.Println(err)
		}

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		ctx = metadata.NewContext(ctx, md)
		c.Set("TracerContext", ctx)

		c.Next()

	}
}
