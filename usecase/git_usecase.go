package usecase

import (
	"fmt"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

// GitInteractor is Interactor
type GitInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	Logger                 interfaces.Logger
}

const (
	requestLabel  = "review"
	reviewedLabel = "wip"
)

var (
	// const variable
	pullRequestOpenedMessage = fmt.Sprintf(`
以下のコマンドをコメントすることでプルリクエストのやり取りを行います。

* %s/request%s : レビュイーがレビュアーにレビューをお願いするコマンド
* %s/reviewed%s : レビュアーによるレビューにて修正点のある場合、レビュイーに返すコマンド
`, "`", "`", "`", "`")
)

// OpenPullRequest is usecase
func (i GitInteractor) OpenPullRequest(pullRequest domain.PullRequestEvent) (res domain.PullRequestEventResponse, err error) {
	// Webhook request variable
	senderUserName := pullRequest.Sender.Login
	organizationName := pullRequest.Repository.Owner.Login
	repositoryName := pullRequest.Repository.Name
	pullRequestID := uint(pullRequest.PullRequest.Number)
	pullRequestTitle := pullRequest.PullRequest.Title
	pullRequestURL := pullRequest.PullRequest.IssueURL

	// Create PullRequest record
	err = i.DatabaseInfrastructure.CreatePullRequest(senderUserName, organizationName, repositoryName, pullRequestID, pullRequestTitle)
	if err != nil {
		return
	}

	// Send message to PullRequest
	err = i.GitInfrastructure.PostMessageToIssue(pullRequestURL, pullRequestOpenedMessage)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

// MergePullRequest is usecase
func (i GitInteractor) MergePullRequest(pullRequest domain.PullRequestEvent) (res domain.PullRequestEventResponse, err error) {
	// Webhook request variable
	senderUserName := pullRequest.Sender.Login
	organizationName := pullRequest.Repository.Owner.Login
	repositoryName := pullRequest.Repository.Name
	pullRequestID := uint(pullRequest.PullRequest.Number)
	pullRequestTitle := pullRequest.PullRequest.Title
	pullRequestURL := pullRequest.PullRequest.IssueURL

	// update PullRequst record
	err = i.DatabaseInfrastructure.MergePullRequest(senderUserName, organizationName, repositoryName, pullRequestID, pullRequestTitle)
	if err != nil {
		return
	}

	// get PullRequest TTL
	pullRequestTTL, err := i.DatabaseInfrastructure.GetPullRequestTTL(organizationName, repositoryName, pullRequestID)
	pullRequestMergedMessage := fmt.Sprintf("レビュー待ち時間総計: %v", pullRequestTTL)

	// Send message to PullRequest
	err = i.GitInfrastructure.PostMessageToIssue(pullRequestURL, pullRequestMergedMessage)
	if err != nil {
		return
	}

	// Unlabeled
	err = i.GitInfrastructure.UnlabelIssue(pullRequestURL)
	if err != nil {
		return
	}

	return domain.PullRequestEventResponse{}, nil
}

// ClosePullRequest is usecase
func (i GitInteractor) ClosePullRequest(pullRequest domain.PullRequestEvent) (res domain.PullRequestEventResponse, err error) {
	// Webhook request variable
	senderUserName := pullRequest.Sender.Login
	organizationName := pullRequest.Repository.Owner.Login
	repositoryName := pullRequest.Repository.Name
	pullRequestID := uint(pullRequest.PullRequest.Number)
	pullRequestTitle := pullRequest.PullRequest.Title
	pullRequestURL := pullRequest.PullRequest.IssueURL

	// update PullRequst record
	err = i.DatabaseInfrastructure.ClosePullRequest(senderUserName, organizationName, repositoryName, pullRequestID, pullRequestTitle)
	if err != nil {
		return
	}

	// Unlabeled
	err = i.GitInfrastructure.UnlabelIssue(pullRequestURL)
	if err != nil {
		return
	}

	return domain.PullRequestEventResponse{}, nil
}

// CommentRequest is usecase
func (i GitInteractor) CommentRequest(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// Webhook request variable
	senderUserName := issueComment.Sender.Login
	organizationName := issueComment.Repository.Owner.Login
	repositoryName := issueComment.Repository.Name
	pullRequestID := uint(issueComment.Issue.Number)
	pullRequestURL := issueComment.Issue.URL

	// Create RequestAction record
	err = i.DatabaseInfrastructure.CreateRequestAction(senderUserName, organizationName, repositoryName, pullRequestID)
	if err != nil {
		return
	}

	// TODO: カレンダーから担当者を取ってくる
	//assigneeUserName, err := i.CalendarInfrastructure.GetStaffThisWeek()
	//err = i.GitInfrastructure.LabelAndAssignIssue(issueComment.PullRequest.IssueURL, person, "review")
	assigneeUserName := "ShotaKitazawa"

	// Labeled "review" & assign user to PullRequest
	err = i.GitInfrastructure.LabelAndAssignIssue(pullRequestURL, assigneeUserName, requestLabel)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

// CommentReviewed is usecase
func (i GitInteractor) CommentReviewed(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// Webhook request variable
	senderUserName := issueComment.Sender.Login
	organizationName := issueComment.Repository.Owner.Login
	repositoryName := issueComment.Repository.Name
	pullRequestID := uint(issueComment.Issue.Number)
	pullRequestURL := issueComment.Issue.URL
	openedPullRequestUserName := issueComment.Issue.User.Login

	// Create RequestAction record
	err = i.DatabaseInfrastructure.CreateReviewedAction(senderUserName, organizationName, repositoryName, pullRequestID)
	if err != nil {
		return
	}

	// Labeled "wip" & assign user to PullRequest
	err = i.GitInfrastructure.LabelAndAssignIssue(pullRequestURL, openedPullRequestUserName, reviewedLabel)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}
