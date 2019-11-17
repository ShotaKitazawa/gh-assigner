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
	Logger                 interfaces.Logger
}

const (
	requestLabel  = "review"
	reviewedLabel = "wip"
)

var (
	// const variable
	pullRequestOpeningMessage = fmt.Sprintf(`
以下のコマンドをコメントすることでプルリクエストのやり取りを行います。

* %s/request%s : レビュイーがレビュアーにレビューをお願いするコマンド
* %s/reviewed%s : レビュアーによるレビューにて修正点のある場合、レビュイーに返すコマンド
`, "`", "`", "`", "`")
)

// OpenPullRequest is usecase
func (i GitInteractor) OpenPullRequest(pullRequest domain.PullRequestEvent) (res domain.PullRequestEventResponse, err error) {
	// User to open PullRequest
	person := pullRequest.Sender.Login

	// Create User record if not exists
	userID, err := i.DatabaseInfrastructure.CreateUserIfNotExists(person)
	if err != nil {
		return
	}

	// Create Repository record if not exists
	repositoryID, err := i.DatabaseInfrastructure.CreateRepositoryIfNotExists(pullRequest.Repository.Owner.Login, pullRequest.Repository.Name)
	if err != nil {
		return
	}

	// Create PullRequest record
	err = i.DatabaseInfrastructure.CreatePullRequest(userID, repositoryID, uint(pullRequest.PullRequest.Number), pullRequest.PullRequest.Title)
	if err != nil {
		return
	}

	// Send message to PullRequest
	err = i.GitInfrastructure.PostMessageToIssue(pullRequest.PullRequest.IssueURL, pullRequestOpeningMessage)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

// CommentRequest is usecase
func (i GitInteractor) CommentRequest(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// TODO: カレンダーから担当者を取ってくる
	//person, err := i.CalendarInfrastructure.GetStaffThisWeek()
	//err = i.GitInfrastructure.LabelToIssue(issueComment.PullRequest.IssueURL, person, "review")
	person := "ShotaKitazawa"

	// Create User record if not exists
	userID, err := i.DatabaseInfrastructure.CreateUserIfNotExists(person)
	if err != nil {
		return
	}

	// Create Repository record if not exists
	repositoryID, err := i.DatabaseInfrastructure.CreateRepositoryIfNotExists(issueComment.Repository.Owner.Login, issueComment.Repository.Name)
	if err != nil {
		return
	}

	// Create RequestAction record
	// TODO: DB の pullrequests table の state カラムよりプルリクエストの現在の状態を取得し、既に review ラベルが付いてるならreturnする
	// TODO: DB の pullrequests table の state カラムの update
	err = i.DatabaseInfrastructure.CreateRequestAction(userID, repositoryID, uint(issueComment.Issue.Number))
	if err != nil {
		return
	}

	// Labeled "review" & assign user to PullRequest
	err = i.GitInfrastructure.LabelToIssue(issueComment.Issue.URL, person, requestLabel)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}

// CommentReviewed is usecase
func (i GitInteractor) CommentReviewed(issueComment domain.IssueCommentEvent) (res domain.PullRequestEventResponse, err error) {
	// User to open PullRequest
	person := issueComment.Issue.User.Login

	// Create User record if not exists
	userID, err := i.DatabaseInfrastructure.CreateUserIfNotExists(person)
	if err != nil {
		return
	}

	// Create Repository record if not exists
	repositoryID, err := i.DatabaseInfrastructure.CreateRepositoryIfNotExists(issueComment.Repository.Owner.Login, issueComment.Repository.Name)
	if err != nil {
		return
	}

	// Create RequestAction record
	// TODO: DB の pullrequests table の state カラムよりプルリクエストの現在の状態を取得し、既に wip ラベルが付いてるならreturnする
	// TODO: DB の pullrequests table の state カラムの update
	err = i.DatabaseInfrastructure.CreateReviewedAction(userID, repositoryID, uint(issueComment.Issue.Number))
	if err != nil {
		return
	}

	// Labeled "wip" & assign user to PullRequest
	err = i.GitInfrastructure.LabelToIssue(issueComment.Issue.URL, person, reviewedLabel)
	if err != nil {
		return
	}
	return domain.PullRequestEventResponse{}, nil
}
