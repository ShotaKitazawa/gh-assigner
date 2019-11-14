package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
)

// GitHubWebhookController is Controller
func GitHubWebhookController(c *gin.Context) {
	// Set gin.Context to context.Context
	ctx := ginContext2standardContext(c, "logger", "db", "gh_user", "gh_token")

	// Get Logger
	logger := ctx.Value("logger").(Logger)

	// Switch by Request Header
	switch c.Request.Header.Get("X-GitHub-Event") {
	case "pull_request":
		request := domain.PullRequestEvent{}
		c.Bind(&request)
		ctx = context.WithValue(ctx, "request", request)

		switch request.Action {
		case "opened", "reopened": // user Open/ReOpen PullRequest
			res, err := usecase.MessagePullRequestOpened(ctx)
			if err != nil {
				//logger.Error(errors.Wrap(err, "GitHubWebhookController: cannot get data"))
				logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}

	case "issue_comment":
		request := domain.IssueCommentEvent{}
		c.Bind(&request)
		switch request.Action {
		case "created": // User created Comment in PullRequest
			// TODO comment.body よりコメント内容に応じた処理
		}
	}
}
