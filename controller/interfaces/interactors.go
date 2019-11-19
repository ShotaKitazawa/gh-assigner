package interfaces

import (
	"github.com/ShotaKitazawa/gh-assigner/domain"
)

type GitInteractor interface {
	OpenPullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	MergePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	ClosePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	CommentRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	CommentReviewed(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
}

type ChatInteractor interface {
	Hello(domain.SlackMessage) (err error)
}
