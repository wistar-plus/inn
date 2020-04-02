package gintool

import (
	"context"

	"github.com/gin-gonic/gin"
)

func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get("TracerContext")
	if exist == false {
		ok = false
		return
	}

	ctx, ok = v.(context.Context)
	return
}
