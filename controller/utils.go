package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
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

func trimNewlineChar(str string) string {
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, "\r", " ", -1)
	str = strings.Replace(str, "\r\n", " ", -1)
	return str
}

func isInternalServerError(c *gin.Context, err error) bool {
	if err != nil {
		loggerInterface, ok := c.Get("logger")
		if !ok {
			panic(errors.New(fmt.Sprintf("gin.Context not in value %s", "logger")))
		}
		logger := loggerInterface.(interfaces.Logger)

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return true
	}
	return false
}
