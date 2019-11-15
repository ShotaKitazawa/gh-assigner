package usecase

import (
	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

type GitInteractor struct {
	GitRepository      interfaces.GitRepository
	DatabaseRepository interfaces.DatabaseRepository
	CalendarRepository interfaces.CalendarRepository
	Logger             interfaces.Logger
}

const (
	request_label  = "review"
	reviewed_label = "wip"
)

func (i GitInteractor) OpenPullRequest(pullRequest domain.PullRequestEvent) (res domain.PullRequestEventResponse, err error) {
	// User to open PullRequest
	person := pullRequest.Sender.Login

	// Create User record if not exists
	err = i.DatabaseRepository.CreateUser(person)
	if err != nil {
		return
	}

	// Create PullRequest record
	err = i.DatabaseRepository.CreatePullRequest(pullRequest.PullRequest.ID, pullRequest.PullRequest.Title, pullRequest.PullRequest.HTMLURL, person)
	if err != nil {
		return
	}

	// Send message to PullRequest
	err = i.GitRepository.PostMessageToIssue(pullRequest.PullRequest.IssueURL, "これはtestです") // TODO
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

func (i GitInteractor) CommentRequest(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// TODO カレンダーから担当者を取ってくる
	//person, err := i.CalendarRepository.GetStaffThisWeek()
	//err = i.GitRepository.LabeledToIssue(issueComment.PullRequest.IssueURL, person, "review")

	// Labeled "review" & assign user to PullRequest
	err = i.GitRepository.LabeledToIssue(issueComment.Issue.URL, "ShotaKitazawa", request_label)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

func (i GitInteractor) CommentReviewed(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// User to open PullRequest
	person := issueComment.Issue.User.Login

	// Labeled "wip" & assign user to PullRequest
	err = i.GitRepository.LabeledToIssue(issueComment.Issue.URL, person, reviewed_label)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}
