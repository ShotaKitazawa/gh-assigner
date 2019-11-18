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

type GitInfrastructureMock struct{}
type DatabaseInfrastructureMock struct{}
type CalendarInfrastructureMock struct{}

func newGitInteractor() *GitInteractor {
	return &GitInteractor{
		GitInfrastructure:      &GitInfrastructureMock{},
		DatabaseInfrastructure: &DatabaseInfrastructureMock{},
		CalendarInfrastructure: &CalendarInfrastructureMock{},
	}
}
func (r GitInfrastructureMock) PostMessageToIssue(url, message string) error {
	// TODO
	return nil
}
func (r GitInfrastructureMock) LabelAndAssignIssue(url, person, label string) error {
	// TODO
	return nil
}
func (r DatabaseInfrastructureMock) CreatePullRequest(username, organizationName, repositoryName string, issueID uint, title string) error {
	// TODO
	return nil
}
func (r DatabaseInfrastructureMock) CreateRequestAction(username, organizationName, repositoryName string, issueID uint) error {
	// TODO
	return nil
}
func (r DatabaseInfrastructureMock) CreateReviewedAction(username, organizationName, repositoryName string, issueID uint) error {
	// TODO
	return nil
}
func (r DatabaseInfrastructureMock) GetPullRequestTTL(issueID uint) (time.Duration, error) {
	// TODO
	return 0, nil
}
func (r CalendarInfrastructureMock) GetStaffThisWeek() (string, error) {
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
