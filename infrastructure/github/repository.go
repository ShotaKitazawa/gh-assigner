package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

// GitInfrastructure is Infrastructure
type GitInfrastructure struct {
	Client *http.Client
	User   string
	Token  string
	Logger interfaces.Logger
}

// PostMessageToIssue is Infrastructure that post message to GitHub Issue/PullRequest.
func (r GitInfrastructure) PostMessageToIssue(url, message string) error {
	// Create Body & Header
	body, err := json.Marshal(PostMessageRequest{Body: message})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		url+"/comments",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.SetBasicAuth(r.User, r.Token)
	req.Header.Set("Content-Type", "application/json")

	// Request
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusCreated {
		if respBody, err := ioutil.ReadAll(resp.Body); err == nil && respBody != nil {
			r.Logger.Error(respBody)
		}
		return fmt.Errorf("Response status is %d, expected %d", resp.StatusCode, http.StatusCreated)
	}

	return nil
}

// LabelAndAssignIssue is Infrastructure that label to GitHub Issue/PullRequest.
func (r GitInfrastructure) LabelAndAssignIssue(url, username, label string) error {
	// Create Body & Header
	body, err := json.Marshal(EditLabelRequest{
		Assignees: []string{username},
		Labels:    []string{label},
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.SetBasicAuth(r.User, r.Token)
	req.Header.Set("Content-Type", "application/json")

	// Request
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		if respBody, err := ioutil.ReadAll(resp.Body); err == nil && respBody != nil {
			r.Logger.Error(string(respBody))
		}
		return fmt.Errorf("Response status is %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	return nil
}

// UnlabelIssue is Infrastructure that unlabel to GitHub Issue/PullRequest.
func (r GitInfrastructure) UnlabelIssue(url string) error {
	// Create Body & Header
	body, err := json.Marshal(EditLabelRequest{
		Labels: make([]string, 0),
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.SetBasicAuth(r.User, r.Token)
	req.Header.Set("Content-Type", "application/json")

	// Request
	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		if respBody, err := ioutil.ReadAll(resp.Body); err == nil && respBody != nil {
			r.Logger.Error(string(respBody))
		}
		return fmt.Errorf("Response status is %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	return nil
}

// GetPullRequestURL is getting PullRequestID based organizationName & repositoryName & issueID
func (r GitInfrastructure) GetPullRequestURL(organizationName, repositoryName string, issueID uint) (string, error) {
	schema := "https://"
	hostname := "github.com/"
	path := strings.Join([]string{organizationName, repositoryName, "pull", strconv.Itoa(int(issueID))}, "/")
	return schema + hostname + path, nil
}
