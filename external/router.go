package external

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //justifying
	"github.com/nlopes/slack"

	"github.com/ShotaKitazawa/gh-assigner/external/slackrtm"
)

var (
	httpRunnnerChan  chan error
	slackRunnnerChan chan error
)

// Run is entrypoint
func Run(ctx context.Context) {
	initializeDeferFunc := Initialize(ctx)
	defer initializeDeferFunc()

	go httpRunner(ctx)
	go slackRunner(ctx)

	var err error
	for {
		select {
		case err = <-httpRunnnerChan:
			panic(err)
		case err = <-slackRunnnerChan:
			panic(err)
		}
	}
}

func httpRunner(ctx context.Context) {
	githubWebhookController := NewGitHubWebhookController(ctx)

	r := gin.New()
	r.Use(gin.Recovery(), Log())

	g := r.Group("/webhook")
	g.POST("/github", func(c *gin.Context) { githubWebhookController.PostWebhook(c) })

	host, err := getContextString(ctx, hostContextKey)
	if err != nil {
		panic(err)
	}
	httpRunnnerChan <- r.Run(host)
}

func slackRunner(ctx context.Context) {
	slackRTMController := NewSlackRTMController(ctx)

	// Get Slack Channel & Token
	channelsStr, err := getContextString(ctx, slackChannelsContextKey)
	if err != nil {
		panic(err)
	}
	channels := strings.Split(channelsStr, ",")
	botID, err := getContextString(ctx, slackBotUserContextKey)
	if err != nil {
		panic(err)
	}
	token, err := getContextString(ctx, slackTokenContextKey)
	if err != nil {
		panic(err)
	}
	c := slackrtm.New(token, channels, botID)

	slackRunnnerChan <- c.Run(func(ev *slack.MessageEvent) { slackRTMController.MessageEvent(ev) })
}
