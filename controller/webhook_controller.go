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
		pullRequest := PullRequestEvent{}
		err = ctx.Bind(&pullRequest)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
		request := domain.GitHubPullRequest{
			Organization:   pullRequest.Repository.Owner.Login,
			Repository:     pullRequest.Repository.Name,
			Number:         uint(pullRequest.PullRequest.Number),
			Title:          pullRequest.PullRequest.Title,
			URL:            pullRequest.PullRequest.IssueURL,
			SenderUsername: pullRequest.Sender.Login,
			OpenedUsername: pullRequest.PullRequest.User.Login,
		}

		// Switch by Request Body
		switch pullRequest.Action {
		case "opened", "reopened": // user Open/ReOpen PullRequest
			res, err := c.Interactor.OpenPullRequest(request)
			if err != nil {
				c.Logger.Error(err.Error())
				ctx.JSON(http.StatusInternalServerError, err)
				return err
			}
			ctx.JSON(http.StatusOK, res)
		case "closed":
			switch pullRequest.PullRequest.Merged {
			case true:
				res, err := c.Interactor.MergePullRequest(domain.GitHubPullRequest{})
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				ctx.JSON(http.StatusOK, res)
			case false:
				res, err := c.Interactor.ClosePullRequest(request)
				if err != nil {
					c.Logger.Error(err.Error())
					ctx.JSON(http.StatusInternalServerError, err)
					return err
				}
				ctx.JSON(http.StatusOK, res)
			}
		}

	case "issue_comment":
		issue := IssueCommentEvent{}
		err := ctx.Bind(&issue)
		if err != nil {
			c.Logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return err
		}
		request := domain.GitHubPullRequest{
			Organization:   issue.Repository.Owner.Login,
			Repository:     issue.Repository.Name,
			Number:         uint(issue.Issue.Number),
			Title:          issue.Issue.Title,
			URL:            issue.Issue.URL,
			SenderUsername: issue.Sender.Login,
			OpenedUsername: issue.Issue.User.Login,
		}

		// Switch by Request Body
		switch issue.Action {
		case "created": // User created Comment in PullRequest
			command := trimNewlineChar(issue.Comment.Body)
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
