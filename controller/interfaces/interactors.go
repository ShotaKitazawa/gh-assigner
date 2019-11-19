package interfaces

import (
	"github.com/ShotaKitazawa/gh-assigner/domain"
)

// GitInteractor is interface of Interactor
type GitInteractor interface {
	OpenPullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	MergePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	ClosePullRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	CommentRequest(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
	CommentReviewed(domain.GitHubPullRequest) (domain.GitHubPullRequestResponse, error)
}

// ChatInteractor is interface of Interactor
type ChatInteractor interface {
	Pong(domain.SlackMessage) error
	ShowDefault(domain.SlackMessage) error
	ShowHelp(domain.SlackMessage) error
}
