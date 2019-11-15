package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
	"github.com/ShotaKitazawa/gh-assigner/domain"
)

// GitHubWebhookController is Controller
type GitHubWebhookController struct {
	Interactor interfaces.GitInteractor
	Logger     interfaces.Logger
}

func (c GitHubWebhookController) PostWebhook(ctx *gin.Context) {

	// Switch by Request Header
	switch ctx.Request.Header.Get("X-GitHub-Event") {
	case "pull_request":
		request := domain.PullRequestEvent{}
		err := ctx.Bind(&request)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		//ctx = context.WithValue(ctx, "request", request)

		// Switch by Request Body
		switch request.Action {
		case "opened", "reopened": // user Open/ReOpen PullRequest
			res, err := c.Interactor.OpenPullRequest(request)
			if err != nil {
				c.Logger.Error(err.Error())
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
	case "closed":
		switch request.PullRequest.Merged {
		case true: // TODO
		case false: // TODO
		}

	case "issue_comment":
		request := domain.IssueCommentEvent{}
		err := ctx.Bind(&request)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		//ctx = context.WithValue(ctx, "request", request)

		// Switch by Request Body
		switch request.Action {
		case "created": // User created Comment in PullRequest
			command := trimNewlineChar(request.Comment.Body)
			if !strings.HasPrefix(command, "/") {
				return
			}
			commands := strings.Split(strings.TrimLeft(command, "/"), " ")

			// Switch by command
			switch commands[0] {
			case "request":
				res, err := c.Interactor.CommentRequest(request)
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				ctx.JSON(http.StatusOK, res)
				return
			case "reviewed":
				res, err := c.Interactor.CommentReviewed(request)
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				ctx.JSON(http.StatusOK, res)
				return
			}
		}
	default:
		c.Logger.Info(fmt.Sprintf("X-GitHub-Event: %s: Skiped.", ctx.Request.Header.Get("X-GitHub-Event")))
	}
}
