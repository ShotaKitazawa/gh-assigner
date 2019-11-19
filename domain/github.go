package domain

type GitHubPullRequest struct {
	Organization   string
	Repository     string
	Number         uint
	Title          string
	URL            string
	SenderUsername string
	OpenedUsername string
}
type GitHubPullRequestResponse struct{}
