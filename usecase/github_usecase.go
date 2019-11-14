package usecase

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
)

func MessagePullRequestOpened(ctx context.Context) (res domain.PullRequestEventResponse, err error) {
	err = github.PostMessage(ctx, "これはtestです")
	return domain.PullRequestEventResponse{}, err
}
