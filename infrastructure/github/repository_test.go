package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func serverInternalServerErrorMock(w http.ResponseWriter, r *http.Request) {
	// Response
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func serverTimeoutMock(w http.ResponseWriter, r *http.Request) {
	// Sleep 15 second
	time.Sleep(15 * time.Second)
	// Response
	w.WriteHeader(http.StatusOK)
	return
}

func TestGitHubInfrastructure(t *testing.T) {
	// Initialize
	t.Parallel()
	infra := &GitInfrastructure{
		User:   "test",
		Token:  "test",
		Logger: &Logger{},
	}

	t.Run("PostMessageToIssue()", func(t *testing.T) {
		// Initialize
		t.Parallel()
		msg := "これはtestです"

		t.Run("期待するリクエストを投げていることのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Test Header
					assert.Equal(t, "POST", r.Method)
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

					// Read Body
					body, err := ioutil.ReadAll(r.Body)
					assert.Nil(t, err)
					data := PostMessageRequest{}
					err = json.Unmarshal(body, &data)
					assert.Nil(t, err)

					// Test Body
					assert.Equal(t, msg, data.Body)

					// Response
					w.WriteHeader(http.StatusCreated)
					return
				},
			))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.PostMessageToIssue(ts.URL, msg)

			// Check
			assert.Nil(t, err)
		})

		t.Run("リクエスト先サーバがInternalServerErrorの際にエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverInternalServerErrorMock))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.PostMessageToIssue(ts.URL, msg)

			// Check
			assert.NotNil(t, err)

		})
		t.Run("10秒経ってもレスポンスが来ない場合Timeoutでエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverTimeoutMock))
			defer ts.Close()
			infra.Client = ts.Client()
			infra.Client.Timeout = time.Duration(10) * time.Second

			// Do
			err := infra.PostMessageToIssue(ts.URL, msg)

			// Check
			assert.NotNil(t, err)
		})
	})
	t.Run("LabelAndAssignIssue()", func(t *testing.T) {
		// Initialize
		t.Parallel()
		label := "test"
		assignee := "test"

		t.Run("期待するリクエストを投げていることのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Test Header
					assert.Equal(t, "PATCH", r.Method)
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

					// Read Body
					body, err := ioutil.ReadAll(r.Body)
					assert.Nil(t, err)
					data := EditLabelRequest{}
					err = json.Unmarshal(body, &data)
					assert.Nil(t, err)

					// Test Body
					assert.Equal(t, 1, len(data.Labels))
					assert.Equal(t, label, data.Labels[0])
					assert.Equal(t, 1, len(data.Assignees))
					assert.Equal(t, assignee, data.Assignees[0])

					// Response
					w.WriteHeader(http.StatusOK)
					return
				},
			))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.LabelAndAssignIssue(ts.URL, assignee, label)

			// Check
			assert.Nil(t, err)
		})
		t.Run("リクエスト先サーバがInternalServerErrorの際にエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverInternalServerErrorMock))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.LabelAndAssignIssue(ts.URL, assignee, label)

			// Check
			assert.NotNil(t, err)
		})
		t.Run("10秒経ってもレスポンスが来ない場合Timeoutでエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverTimeoutMock))
			defer ts.Close()
			infra.Client = ts.Client()
			infra.Client.Timeout = time.Duration(10) * time.Second

			// Do
			err := infra.LabelAndAssignIssue(ts.URL, assignee, label)

			// Check
			assert.NotNil(t, err)
		})
	})
	t.Run("UnlabelIssue()", func(t *testing.T) {
		// Initialize
		t.Parallel()

		t.Run("期待するリクエストを投げていることのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Test Header
					assert.Equal(t, "PATCH", r.Method)
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

					// Read Body
					body, err := ioutil.ReadAll(r.Body)
					assert.Nil(t, err)
					data := EditLabelRequest{}
					err = json.Unmarshal(body, &data)
					assert.Nil(t, err)

					// Test Body
					assert.Equal(t, 0, len(data.Labels))

					// Response
					w.WriteHeader(http.StatusOK)
					return
				},
			))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.UnlabelIssue(ts.URL)

			// Check
			assert.Nil(t, err)
		})
		t.Run("リクエスト先サーバがInternalServerErrorの際にエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverInternalServerErrorMock))
			defer ts.Close()
			infra.Client = ts.Client()

			// Do
			err := infra.UnlabelIssue(ts.URL)

			// Check
			assert.NotNil(t, err)
		})
		t.Run("10秒経ってもレスポンスが来ない場合Timeoutでエラーを返すことのテスト", func(t *testing.T) {
			t.Parallel()

			// Http Server Mock
			ts := httptest.NewServer(http.HandlerFunc(serverTimeoutMock))
			defer ts.Close()
			infra.Client = ts.Client()
			infra.Client.Timeout = time.Duration(10) * time.Second

			// Do
			err := infra.UnlabelIssue(ts.URL)

			// Check
			assert.NotNil(t, err)
		})
	})
}

// for test
type Logger struct{}

func (logger Logger) Debug(args ...interface{}) {}
func (logger Logger) Info(args ...interface{})  {}
func (logger Logger) Warn(args ...interface{})  {}
func (logger Logger) Error(args ...interface{}) {}
