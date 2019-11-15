package domain

type GitHubPostMessageRequest struct {
	Body string `json:"body"`
}
type GitHubEditLabelRequest struct {
	Assignees []string `json:"assignees"`
	State     string   `json:"state"`
	Labels    []string `json:"labels"`
}
