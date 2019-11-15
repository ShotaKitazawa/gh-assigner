package interfaces

import (
	"context"
)

type GitRepository interface {
	PostMessageToIssue(context.Context, string) error
	GetPersonToOpenPullRequest(context.Context) (string, error)
	LabeledToIssue(context.Context, string, string) error
}

type CalendarRepository interface {
	GetStaffThisWeek() (string, error)
}
