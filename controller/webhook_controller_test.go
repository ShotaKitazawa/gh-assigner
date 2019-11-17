package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

var (
	flagOpenPullRequest  = false
	flagMergePullRequest = false
	flagClosePullRequest = false
	flagCommentRequest   = false
	flagCommentReviewed  = false
)

type InteractorMock struct{}

func newController() *GitHubWebhookController {
	return &GitHubWebhookController{
		Interactor: &InteractorMock{},
	}
}
func (i InteractorMock) OpenPullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error) {
	flagOpenPullRequest = true
	return domain.PullRequestEventResponse{}, nil
}
func (i InteractorMock) MergePullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error) {
	flagMergePullRequest = true
	return domain.PullRequestEventResponse{}, nil
}
func (i InteractorMock) ClosePullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error) {
	flagClosePullRequest = true
	return domain.PullRequestEventResponse{}, nil
}

func (i InteractorMock) CommentRequest(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error) {
	flagCommentRequest = true
	return domain.PullRequestEventResponse{}, nil
}
func (i InteractorMock) CommentReviewed(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error) {
	flagCommentReviewed = true
	return domain.PullRequestEventResponse{}, nil
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
			body, err := json.Marshal(domain.PullRequestEvent{
				Action: "opened",
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
			controller := newController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagOpenPullRequest, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		/*
			t.Run("PullRequestをMergeするとinteractor.MergePullRequest()が呼ばれることのテスト", func(t *testing.T) {
				// Initialize
				t.Parallel()
				responseWriter := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseWriter)

				// Create Body & Header
				// TODO: marshal に interface を渡す
				body, err := json.Marshal(domain.PullRequestEvent{
					Action:      "closed",
					PullRequest: GitHubPullRequest{Merged: true},
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
				controller := newController()
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
				// TODO: marshal に interface を渡す
				body, err := json.Marshal(domain.PullRequestEvent{
					Action: "closed",
					PullRequest: GitHubPullRequest{Merged: true},
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
				controller := newController()
				controller.PostWebhook(ctx)

				// assert
				assert.Equal(t, flagClosePullRequest, true)
				assert.Equal(t, responseWriter.Code, http.StatusOK)
			})
		*/
		t.Run("PullRequestに'/request'とCommentするとinteractor.CommentRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(domain.IssueCommentEvent{
				Action:  "created",
				Comment: GitHubComment{Body: "/request"},
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
			controller := newController()
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
			body, err := json.Marshal(domain.IssueCommentEvent{
				Action:  "created",
				Comment: GitHubComment{Body: "/reviewed"},
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
			controller := newController()
			controller.PostWebhook(ctx)

			// assert
			assert.Equal(t, flagCommentReviewed, true)
			assert.Equal(t, responseWriter.Code, http.StatusOK)
		})
		t.Run("リポジトリのPushをしても何も発火しないことのテスト", func(t *testing.T) {
			// TODO
		})
	})
}

type GitHubComment struct {
	URL               string            `json:"url"`
	HTMLURL           string            `json:"html_url"`
	IssueURL          string            `json:"issue_url"`
	ID                int               `json:"id"`
	NodeID            string            `json:"node_id"`
	User              domain.GitHubUser `json:"user"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
	AuthorAssociation string            `json:"author_association"`
	Body              string            `json:"body"`
}
