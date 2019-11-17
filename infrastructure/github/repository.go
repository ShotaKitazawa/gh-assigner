package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

// LabelToIssue is Infrastructure that label to GitHub Issue/PullRequest.
func (r GitInfrastructure) LabelToIssue(url, person, label string) error {
	// Create Body & Header
	body, err := json.Marshal(EditLabelRequest{
		Assignees: []string{person},
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
