package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

// GitRepository is Repository
type GitRepository struct {
	User   string
	Token  string
	Logger interfaces.Logger
}

func (r GitRepository) PostMessageToIssue(url, message string) error {
	// Create Body & Header
	body, err := json.Marshal(domain.GitHubPostMessageRequest{Body: message})
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
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusCreated {
		respBody, _ := ioutil.ReadAll(resp.Body)
		r.Logger.Error(string(respBody))
		return fmt.Errorf("Response status is %d, expected %d", resp.StatusCode, http.StatusCreated)
	}

	return nil
}

func (r GitRepository) LabeledToIssue(url, person, label string) error {
	// Create Body & Header
	body, err := json.Marshal(domain.GitHubEditLabelRequest{
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
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		r.Logger.Error(string(respBody))
		return fmt.Errorf("Response status is %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	return nil
}
