package controller

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ginContext2standardContext(c *gin.Context, args ...string) context.Context {
	ctx := context.Background()
	for _, val := range args {
		inter, ok := c.Get(val)
		if !ok {
			panic(errors.New(fmt.Sprintf("gin.Context not in value %s", val)))
		}
		ctx = context.WithValue(ctx, val, inter)
	}
	return ctx
}
