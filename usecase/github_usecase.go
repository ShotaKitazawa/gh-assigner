package usecase

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

type GitHubInteractor struct {
	GitRepository      interfaces.GitRepository
	CalendarRepository interfaces.CalendarRepository
	Logger             interfaces.Logger
}

func (i GitHubInteractor) OpenPullRequest(ctx context.Context) (res domain.PullRequestEventResponse, err error) {
	err = i.GitRepository.PostMessageToIssue(ctx, "これはtestです")
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

func (i GitHubInteractor) CommentRequest(ctx context.Context) (res domain.PullRequestEventResponse, err error) {
	// TODO カレンダーから担当者を取ってくる
	//person, err := i.CalendarRepository.GetStaffThisWeek()
	//err = i.GitRepository.LabeledToIssue(ctx, person, "review")

	err = i.GitRepository.LabeledToIssue(ctx, "ShotaKitazawa", "review")
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

func (i GitHubInteractor) CommentReviewed(ctx context.Context) (res domain.PullRequestEventResponse, err error) {
	person, err := i.GitRepository.GetPersonToOpenPullRequest(ctx)
	if err != nil {
		return
	}
	err = i.GitRepository.LabeledToIssue(ctx, person, "wip")
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}
