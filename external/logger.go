package external

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
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
