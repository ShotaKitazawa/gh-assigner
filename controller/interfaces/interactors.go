package interfaces

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

type GitHubInteractor interface {
	MessagePullRequestOpened(context.Context) (domain.PullRequestEventResponse, error)
}
