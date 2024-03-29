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
type ChatInfrastructureMock struct{}

func newGitInteractor() *GitInteractor {
	return &GitInteractor{
		GitInfrastructure:      &GitInfrastructureMock{},
		DatabaseInfrastructure: &DatabaseInfrastructureMock{},
		CalendarInfrastructure: &CalendarInfrastructureMock{},
		ChatInfrastructure:     &ChatInfrastructureMock{},
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
func (r GitInfrastructureMock) UnlabelIssue(url string) error {
	// TODO
	return nil
}

func (r GitInfrastructureMock) GetPullRequestURL(organizationName, repositoryName string, issueID uint) (string, error) {
	// TODO
	return "", nil
}

func (r DatabaseInfrastructureMock) CreatePullRequest(username, organizationName, repositoryName string, issueID uint, title string) error {
	// TODO
	return nil
}

func (r DatabaseInfrastructureMock) MergePullRequest(username, organizationName, repositoryName string, issueID uint, title string) error {
	// TODO
	return nil
}

func (r DatabaseInfrastructureMock) ClosePullRequest(username, organizationName, repositoryName string, issueID uint, title string) error {
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

func (r DatabaseInfrastructureMock) GetPullRequestTTL(organizationName, repositoryName string, issueID uint) (time.Duration, error) {
	// TODO
	return 0, nil
}

func (r DatabaseInfrastructureMock) SelectPullRequestTTLs(organizationName, repositoryName string, period int) (map[uint]time.Duration, error) {
	// TODO
	return nil, nil
}

func (r CalendarInfrastructureMock) GetCurrentStaff() (string, error) {
	// TODO
	return "", nil
}

func (r ChatInfrastructureMock) SendMessage(string, string) error {
	// TODO
	return nil
}
func (r ChatInfrastructureMock) SendMessageToDefaultChannel(string) error {
	// TODO
	return nil
}
func (r ChatInfrastructureMock) SendImage(string, string) error {
	// TODO
	return nil
}
func (r ChatInfrastructureMock) SendImageToDefaultChannel(string) error {
	// TODO
	return nil
}

// TODO
func TestGitInteractor(t *testing.T) {
	t.Parallel()
	t.Run("OpenPullRequest()", func(t *testing.T) {
		t.Parallel()
	})
	t.Run("CommentRequest()", func(t *testing.T) {
		t.Parallel()
	})
	t.Run("CommentReviewed()", func(t *testing.T) {
		t.Parallel()
	})
}
