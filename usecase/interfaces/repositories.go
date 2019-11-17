package interfaces

import (
	"time"
)

type GitInfrastructure interface {
	PostMessageToIssue(string, string) error
	LabelToIssue(string, string, string) error
}

type DatabaseInfrastructure interface {
	CreatePullRequest(uint, uint, uint, string) error
	CreateRequestAction(uint, uint, uint) error
	CreateReviewedAction(uint, uint, uint) error
	CreateUserIfNotExists(string) (uint, error)
	CreateRepositoryIfNotExists(string, string) (uint, error)
	GetPullRequestTTL(uint) (time.Duration, error)
}

type CalendarInfrastructure interface {
	GetStaffThisWeek() (string, error)
}
