package interfaces

import (
	"time"
)

type GitRepository interface {
	PostMessageToIssue(string, string) error
	LabelToIssue(string, string, string) error
}

type CalendarRepository interface {
	GetStaffThisWeek() (string, error)
}

type DatabaseRepository interface {
	CreatePullRequest(int, string, string, string) error
	CreateRequestAction(int, string) error
	CreateReviewedAction(int, string) error
	CreateUser(string) error
	GetPullRequestTTL(int) (time.Duration, error)
}
