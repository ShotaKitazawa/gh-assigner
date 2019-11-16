package usecase

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type GitRepositoryMock struct{}
type DatabaseRepositoryMock struct{}
type CalendarRepositoryMock struct{}

func newGitInteractor() *GitInteractor {
	return &GitInteractor{
		GitRepository:      &GitRepositoryMock{},
		DatabaseRepository: &DatabaseRepositoryMock{},
		CalendarRepository: &CalendarRepositoryMock{},
	}
}
func (r GitRepositoryMock) PostMessageToIssue(url, message string) error {
	// TODO
	return nil
}
func (r GitRepositoryMock) LabelToIssue(url, person, label string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreatePullRequest(issueID int, title, username, url string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateRequestAction(issueID int, username string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateReviewedAction(issueID int, username string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateUser(username string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) GetPullRequestTTL(issueID int) (time.Duration, error) {
	// TODO
	return 0, nil
}
func (r CalendarRepositoryMock) GetStaffThisWeek() (string, error) {
	// TODO
	return "", nil
}

// TODO
func TestGitInteractor(t *testing.T) {
	t.Run("OpenPullRequest()", func(t *testing.T) {
	})
	t.Run("CommentRequest()", func(t *testing.T) {
	})
	t.Run("CommentReviewed()", func(t *testing.T) {
	})
}
