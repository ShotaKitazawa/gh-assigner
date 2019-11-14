package github

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

func PostMessage(ctx context.Context, message string) error {
	// Get Logger
	logger := ctx.Value("logger").(interfaces.Logger)
	logger.Info("PostMessage")

	// Get requested URL
	url := ctx.Value("request").(domain.PullRequestEvent).PullRequest.IssueURL

	// Get GitHub User & Token
	ghUser := ctx.Value("gh_user").(string)
	ghToken := ctx.Value("gh_token").(string)

	// TODO: Get GitHub Secret

	// Create Body & Header
	body, err := json.Marshal(domain.GitHubMessage{Body: message})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		url+"/comments",
		bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.SetBasicAuth(ghUser, ghToken)
	req.Header.Set("Content-Type", "application/json")

	// Post message
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check response
	if resp.StatusCode != http.StatusCreated {
		respBody, _ := ioutil.ReadAll(resp.Body)
		logger.Error(string(respBody))
		return errors.New(fmt.Sprintf("Response status is %d, expected 201", resp.StatusCode))
	}

	return nil
}
