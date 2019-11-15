package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/gin-gonic/gin"
)

// GitHubWebhookController is Controller
type GitHubWebhookController struct {
	Interactor interfaces.GitInteractor
	Logger     interfaces.Logger
}

// PostWebhook is called by GitHub Webhook
func (c GitHubWebhookController) PostWebhook(ctx *gin.Context) (err error) {

	// Switch by Request Header
	switch ctx.Request.Header.Get("X-GitHub-Event") {
	case "pull_request":
		request := domain.PullRequestEvent{}
		err = ctx.Bind(&request)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}

		// Switch by Request Body
		switch request.Action {
		case "opened", "reopened": // user Open/ReOpen PullRequest
			res, err := c.Interactor.OpenPullRequest(request)
			if err != nil {
				c.Logger.Error(err.Error())
				ctx.JSON(http.StatusInternalServerError, err)
				return err
			}
			ctx.JSON(http.StatusOK, res)
		case "closed":
			switch request.PullRequest.Merged {
			case true: // TODO
			case false: // TODO
			}
		}

	case "issue_comment":
		request := domain.IssueCommentEvent{}
		err := ctx.Bind(&request)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}

		// Switch by Request Body
		switch request.Action {
		case "created": // User created Comment in PullRequest
			command := trimNewlineChar(request.Comment.Body)
			if !strings.HasPrefix(command, "/") {
				return nil
			}
			commands := strings.Split(strings.TrimLeft(command, "/"), " ")

			// Switch by command
			switch commands[0] {
			case "request":
				res, err := c.Interactor.CommentRequest(request)
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				ctx.JSON(http.StatusOK, res)
			case "reviewed":
				res, err := c.Interactor.CommentReviewed(request)
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				ctx.JSON(http.StatusOK, res)
			}
		}
	default:
		c.Logger.Info(fmt.Sprintf("X-GitHub-Event: %s: Skiped.", ctx.Request.Header.Get("X-GitHub-Event")))
	}
	return nil
}
