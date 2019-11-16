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
	"gopkg.in/go-playground/assert.v1"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

var (
	flagOpenPullRequest = false
	flagCommentRequest  = false
	flagCommentReviewed = false
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
func (i InteractorMock) CommentRequest(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error) {
	flagCommentRequest = true
	return domain.PullRequestEventResponse{}, nil
}
func (i InteractorMock) CommentReviewed(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error) {
	flagCommentReviewed = true
	return domain.PullRequestEventResponse{}, nil
}

func TestGitHubWebhookController(t *testing.T) {
	t.Run("PostWebhook()", func(t *testing.T) {
		t.Run("PullRequestをOpenするとinteractor.OpenPullRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			gin.SetMode(gin.TestMode)
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(domain.PullRequestEvent{
				Action: "opened",
			})
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
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
		t.Run("PullRequestに'/request'とCommentするとinteractor.CommentRequest()が呼ばれることのテスト", func(t *testing.T) {
			// Initialize
			gin.SetMode(gin.TestMode)
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(domain.IssueCommentEvent{
				Action:  "created",
				Comment: GitHubComment{Body: "/request"},
			})
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
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
			gin.SetMode(gin.TestMode)
			responseWriter := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseWriter)

			// Create Body & Header
			body, err := json.Marshal(domain.IssueCommentEvent{
				Action:  "created",
				Comment: GitHubComment{Body: "/reviewed"},
			})
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(
				"POST",
				"http://localhost:8080/",
				bytes.NewBuffer(body),
			)
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
	})
	t.Run("trimNewlineChar()", func(t *testing.T) {
		t.Run("改行や余分な空白が取り除かれているテスト", func(t *testing.T) {
			expected := "hoge fuga piyo"
			actual := trimNewlineChar(`hoge
fuga
    piyo
`)
			assert.Equal(t, expected, actual)
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