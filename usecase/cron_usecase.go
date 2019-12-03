package usecase

import (
	"fmt"

	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

// CronInteractor is Interactor
type CronInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	ImageInfrastructure    interfaces.ImageInfrastructure
	Logger                 interfaces.Logger
}

func (i CronInteractor) SendImageWithReviewWaitTimeGraph(organization, repository string, period int) (err error) {
	// Get PullRequest TTL last `period` days
	times, err := i.DatabaseInfrastructure.SelectPullRequestTTLs(organization, repository, period)
	if err != nil {
		return
	}

	// Send TTL info
	var reviewWaitTimeMsg string
	for id, time := range times {
		issueURL, err := i.GitInfrastructure.GetPullRequestURL(organization, repository, id)
		if err != nil {
			return err
		}
		reviewWaitTimeMsg += fmt.Sprintf("%s\n> %v\n", issueURL, time)
	}
	err = i.ChatInfrastructure.SendMessageToDefaultChannel(reviewWaitTimeMsg)
	if err != nil {
		return
	}

	// Create Bar Graph Image
	filepath, err := i.ImageInfrastructure.CreateGraphWithReviewWaitTime(times)
	if err != nil {
		return
	}

	// Send Bar Graph Image
	err = i.ChatInfrastructure.SendImageToDefaultChannel(filepath)
	if err != nil {
		return
	}

	// Delete Bar Graph Image
	err = i.ImageInfrastructure.DeleteFile(filepath)
	if err != nil {
		return
	}

	return
}
