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
func (r DatabaseRepositoryMock) CreatePullRequest(userID, repositoryID, issueID uint, title string) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateRequestAction(userID, repositoryID, issueID uint) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateReviewedAction(userID, repositoryID, issueID uint) error {
	// TODO
	return nil
}
func (r DatabaseRepositoryMock) CreateUserIfNotExists(username string) (uint, error) {
	// TODO
	return 0, nil
}
func (r DatabaseRepositoryMock) CreateRepositoryIfNotExists(organication, repository string) (uint, error) {
	// TODO
	return 0, nil
}
func (r DatabaseRepositoryMock) GetPullRequestTTL(issueID uint) (time.Duration, error) {
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
