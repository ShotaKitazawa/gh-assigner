package external

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
	"github.com/gin-gonic/gin"
)

// Logger is struct for using Debug/Info/Warn/Error methods.
// Those methods is custumize for Google Stackdriver
type Logger struct{}

// Debug is debug message (severity 100)
func (logger Logger) Debug(args ...interface{}) {
	logger.log(int(logging.Debug), args...)
}

// Info is information message (severity 200)
func (logger Logger) Info(args ...interface{}) {
	logger.log(int(logging.Info), args...)
}

// Warn is warning message (severity 400)
func (logger Logger) Warn(args ...interface{}) {
	logger.log(int(logging.Warning), args...)
}

// Error is error message (severity 500)
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

// Log is set gin.Context, using everywhere
func Log() gin.HandlerFunc {
	type Message struct {
		Status  int    `json:"status"`
		Latency int64  `json:"latency"`
		Source  string `json:"source"`
		Method  string `json:"method"`
		Path    string `json:"path"`
	}
	return func(c *gin.Context) {
		start := time.Now()

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
