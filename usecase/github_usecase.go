package usecase

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

type GitHubInteractor struct {
	GitHubRepository interfaces.GitHubRepository
	Logger           interfaces.Logger
}

func (i GitHubInteractor) MessagePullRequestOpened(ctx context.Context) (res domain.PullRequestEventResponse, err error) {
	err = i.GitHubRepository.PostMessage(ctx, "これはtestです")
	return domain.PullRequestEventResponse{}, err
}
