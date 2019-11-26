package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

var (
	flagOpenPullRequest  = false
	flagMergePullRequest = false
	flagClosePullRequest = false
	flagCommentRequest   = false
	flagCommentReviewed  = false
)

type GitHubWebhookControllerInteractorMock struct{}
type GitHubWebhookControllerLoggerMock struct{}

func newGitHubWebhookController() *GitHubWebhookController {
	return &GitHubWebhookController{
		Interactor: &GitHubWebhookControllerInteractorMock{},
		Logger:     &GitHubWebhookControllerLoggerMock{},
	}
}
func (i GitHubWebhookControllerInteractorMock) OpenPullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error) {
	flagOpenPullRequest = true
	return domain.GitHubPullRequestResponse{}, nil
}
func (i GitHubWebhookControllerInteractorMock) MergePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error) {
	flagMergePullRequest = true
	return domain.GitHubPullRequestResponse{}, nil
}
func (i GitHubWebhookControllerInteractorMock) ClosePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error) {
	flagClosePullRequest = true
	return domain.GitHubPullRequestResponse{}, nil
}

func (i GitHubWebhookControllerInteractorMock) CommentRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error) {
	flagCommentRequest = true
	return domain.GitHubPullRequestResponse{}, nil
}
func (i GitHubWebhookControllerInteractorMock) CommentReviewed(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error) {
	flagCommentReviewed = true
	return domain.GitHubPullRequestResponse{}, nil
}
func (l GitHubWebhookControllerLoggerMock) Debug(args ...interface{}) {
}
func (l GitHubWebhookControllerLoggerMock) Info(args ...interface{}) {
}
func (l GitHubWebhookControllerLoggerMock) Warn(args ...interface{}) {
}
func (l GitHubWebhookControllerLoggerMock) Error(args ...interface{}) {
}

func TestGitHubWebhookController(t *testing.T) {
	// Initialize
	t.Parallel()
	gin.SetMode(gin.TestMode)

	t.Run("PostWebhook()", func(t *testing.T) {
		// Initialize
		t.Parallel()

		t.Run("PullRequestをOpenするとinteractor.OpenPullRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"action": "opened",
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "pull_request")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagOpenPullRequest, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("PullRequestをMergeするとinteractor.MergePullRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"action": "closed",
				"pull_request": map[string]interface{}{
					"merged": true,
				},
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "pull_request")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagMergePullRequest, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("PullRequestをCloseするとinteractor.ClosePullRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"action": "closed",
				"pull_request": map[string]interface{}{
					"merged": false,
				},
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "pull_request")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagClosePullRequest, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("PullRequestに'/request'とCommentするとinteractor.CommentRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"action": "created",
				"comment": map[string]interface{}{
					"body": "/request",
				},
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "issue_comment")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagCommentRequest, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("PullRequestに'/reviewed'とCommentするとinteractor.CommentReviewed()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"action": "created",
				"comment": map[string]interface{}{
					"body": "/reviewed",
				},
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "issue_comment")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagCommentReviewed, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("リポジトリのPush(未対応のイベント)のwebhookには403を返すことのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(map[string]interface{}{
				"dummy": "dummy",
			})
			assert.Nil(t, err)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-GitHub-Event", "push")
			ctx.Request = req

			// call
			controller := newGitHubWebhookController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, responseWriter.Code, http.StatusForbidden)
		})
	})
}
