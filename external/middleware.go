package external

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DB is set gin.Context, using in Repository
func DB() gin.HandlerFunc {
	var err error

	// Connect DB
	dsn := getContext(ctx, "dsn").(string)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to DB"))
	}

	// TODO: Create & Migrate DB

	// Set gin.Context
	return func(c *gin.Context) {
		c.Set("db", db)
	}
}

// GitHub is set gin.Context, using in Repository
func GitHub() gin.HandlerFunc {
	// Get Github User & Token
	ghUser := getContext(ctx, "gh_user").(string)
	ghToken := getContext(ctx, "gh_token").(string)

	// Set gin.Context
	return func(c *gin.Context) {
		c.Set("gh_user", ghUser)
		c.Set("gh_token", ghToken)
	}
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
