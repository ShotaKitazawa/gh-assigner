package interfaces

import (
	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/nlopes/slack"
)

type GitInteractor interface {
	OpenPullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error)
	MergePullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error)
	ClosePullRequest(domain.PullRequestEvent) (domain.PullRequestEventResponse, error)
	CommentRequest(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error)
	CommentReviewed(domain.IssueCommentEvent) (domain.PullRequestEventResponse, error)
}

type SlackInteractor interface {
	Hello(*slack.MessageEvent) (err error)
}
