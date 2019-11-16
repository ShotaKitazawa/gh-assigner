package github

// PostMessageRequest is GitHub API post to Issue/PullRequest
type PostMessageRequest struct {
	Body string `json:"body"`
}

// EditLabelRequest is GitHub API label to Issue/PullRequest
type EditLabelRequest struct {
	Assignees []string `json:"assignees"`
	Labels    []string `json:"labels"`
}
