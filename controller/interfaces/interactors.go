package interfaces

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

type GitHubInteractor interface {
	OpenPullRequest(context.Context) (domain.PullRequestEventResponse, error)
	CommentRequest(context.Context) (domain.PullRequestEventResponse, error)
	CommentReviewed(context.Context) (domain.PullRequestEventResponse, error)
}
