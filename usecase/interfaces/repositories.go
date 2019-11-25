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
	SelectPullRequestTTLs(string, string, int) ([]time.Duration, error)
}

type CalendarInfrastructure interface {
	GetCurrentStaff() (string, error)
}

type ChatInfrastructure interface {
	SendMessage(string, string) error
	SendMessageToDefaultChannel(string) error
	SendImage(string, string) error
	SendImageToDefaultChannel(string) error
}

type ImageInfrastructure interface {
	CreateGraphWithReviewWaitTime([]time.Duration) (string, error)
	DeleteFile(string) error
}
