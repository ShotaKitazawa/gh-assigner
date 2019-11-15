package external

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
	"github.com/gin-gonic/gin"
)

type Logger struct{}

func (logger Logger) Debug(args ...interface{}) {
	logger.log(int(logging.Debug), args...)
}
func (logger Logger) Info(args ...interface{}) {
	logger.log(int(logging.Info), args...)
}
func (logger Logger) Warn(args ...interface{}) {
	logger.log(int(logging.Warning), args...)
}
func (logger Logger) Error(args ...interface{}) {
	logger.log(int(logging.Error), args...)
}

func (logger Logger) log(severity int, args ...interface{}) {
	entry := map[string]interface{}{
		"time":     time.Now().Format(time.RFC3339Nano),
		"severity": severity,
		"message":  args,
	}
	b, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

type Message struct {
	Status  int    `json:"status"`
	Latency int64  `json:"latency"`
	Source  string `json:"source"`
	Method  string `json:"method"`
	Path    string `json:"path"`
}

// Log is set gin.Context, using everywhere
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		logger := Logger{}
		c.Set("logger", logger)

		c.Next()

		latency := time.Since(start)
		message := Message{
			Status:  c.Writer.Status(),
			Latency: int64(latency),
			Source:  c.ClientIP(),
			Method:  c.Request.Method,
			Path:    c.Request.URL.Path,
		}

		entry := map[string]interface{}{
			"time":     time.Now().Format(time.RFC3339Nano),
			"severity": message.Status,
			"message":  message,
		}
		b, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}
}
