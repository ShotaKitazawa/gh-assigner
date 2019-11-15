package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
	"github.com/ShotaKitazawa/gh-assigner/domain"
)

type GitHubWebhookController struct {
	Interactor interfaces.GitHubInteractor
}

// GitHubWebhookController is Controller
func (controller GitHubWebhookController) PostWebhook(c *gin.Context) {
	// Set gin.Context to context.Context
	ctx := ginContext2standardContext(c, "logger", "db", "gh_user", "gh_token")

	// Switch by Request Header
	switch c.Request.Header.Get("X-GitHub-Event") {
	case "pull_request":
		request := domain.PullRequestEvent{}
		err := c.Bind(&request)
		if isInternalServerError(c, err) {
			return
		}
		ctx = context.WithValue(ctx, "request", request)

		switch request.Action {
		case "opened", "reopened": // user Open/ReOpen PullRequest
			res, err := controller.Interactor.MessagePullRequestOpened(ctx)
			if isInternalServerError(c, err) {
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
