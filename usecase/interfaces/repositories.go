package interfaces

import (
	"time"
)

type GitInfrastructure interface {
	PostMessageToIssue(string, string) error
	LabelAndAssignIssue(string, string, string) error
	UnlabelIssue(string) error
}

type DatabaseInfrastructure interface {
	CreatePullRequest(string, string, string, uint, string) error
	MergePullRequest(string, string, string, uint, string) error
	ClosePullRequest(string, string, string, uint, string) error
	CreateRequestAction(string, string, string, uint) error
	CreateReviewedAction(string, string, string, uint) error
	GetPullRequestTTL(string, string, uint) (time.Duration, error)
	//CreateUserIfNotExists(string) (uint, error)
	//CreateRepositoryIfNotExists(string, string) (uint, error)
}

type CalendarInfrastructure interface {
	GetStaffThisWeek() (string, error)
}

type ChatInfrastructure interface {
	SendMessage(string, string) error
	SendMessageToDefaultChannel(string) error
}
