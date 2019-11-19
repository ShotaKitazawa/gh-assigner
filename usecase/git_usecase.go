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
func (i GitInteractor) OpenPullRequest(pr domain.GitHubPullRequest) (res domain.GitHubPullRequestResponse, err error) {
	// Create PullRequest record
	err = i.DatabaseInfrastructure.CreatePullRequest(pr.SenderUsername, pr.Organization, pr.Repository, pr.Number, pr.Title)
	if err != nil {
		return
	}

	// Send message to PullRequest
	err = i.GitInfrastructure.PostMessageToIssue(pr.URL, pullRequestOpenedMessage)
	if err != nil {
		return
	}
	return domain.GitHubPullRequestResponse{}, nil
}

// MergePullRequest is usecase
func (i GitInteractor) MergePullRequest(pr domain.GitHubPullRequest) (res domain.GitHubPullRequestResponse, err error) {
	// update PullRequst record
	err = i.DatabaseInfrastructure.MergePullRequest(pr.SenderUsername, pr.Organization, pr.Repository, pr.Number, pr.Title)
	if err != nil {
		return
	}

	// get PullRequest TTL
	pullRequestTTL, err := i.DatabaseInfrastructure.GetPullRequestTTL(pr.Organization, pr.Repository, pr.Number)
	pullRequestMergedMessage := fmt.Sprintf("レビュー待ち時間総計: %v", pullRequestTTL)

	// Send message to PullRequest
	err = i.GitInfrastructure.PostMessageToIssue(pr.URL, pullRequestMergedMessage)
	if err != nil {
		return
	}

	// Unlabeled
	err = i.GitInfrastructure.UnlabelIssue(pr.URL)
	if err != nil {
		return
	}

	return domain.GitHubPullRequestResponse{}, nil
}

// ClosePullRequest is usecase
func (i GitInteractor) ClosePullRequest(pr domain.GitHubPullRequest) (res domain.GitHubPullRequestResponse, err error) {
	// update PullRequst record
	err = i.DatabaseInfrastructure.ClosePullRequest(pr.SenderUsername, pr.Organization, pr.Repository, pr.Number, pr.Title)
	if err != nil {
		return
	}

	// Unlabeled
	err = i.GitInfrastructure.UnlabelIssue(pr.URL)
	if err != nil {
		return
	}

	return domain.GitHubPullRequestResponse{}, nil
}

// CommentRequest is usecase
func (i GitInteractor) CommentRequest(pr domain.GitHubPullRequest) (res domain.GitHubPullRequestResponse, err error) {
	// Create RequestAction record
	err = i.DatabaseInfrastructure.CreateRequestAction(pr.SenderUsername, pr.Organization, pr.Repository, pr.Number)
	if err != nil {
		return
	}

	// Get current SRE-Order staff
	assigneeUserName, err := i.CalendarInfrastructure.GetCurrentStaff()
	if err != nil {
		return
	}

	// Labeled "review" & assign user to PullRequest
	err = i.GitInfrastructure.LabelAndAssignIssue(pr.URL, assigneeUserName, requestLabel)
	if err != nil {
		return
	}

	// Send message to Slack
	err = i.ChatInfrastructure.SendMessageToDefaultChannel("これはtestです")
	if err != nil {
		return
	}

	return domain.GitHubPullRequestResponse{}, nil
}

// CommentReviewed is usecase
func (i GitInteractor) CommentReviewed(pr domain.GitHubPullRequest) (res domain.GitHubPullRequestResponse, err error) {
	// Create RequestAction record
	err = i.DatabaseInfrastructure.CreateReviewedAction(pr.SenderUsername, pr.Organization, pr.Repository, pr.Number)
	if err != nil {
		return
	}

	// Labeled "wip" & assign user to PullRequest
	err = i.GitInfrastructure.LabelAndAssignIssue(pr.URL, pr.OpenedUsername, reviewedLabel)
	if err != nil {
		return
	}
	return domain.GitHubPullRequestResponse{}, nil
}
