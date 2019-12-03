package usecase

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

// ChatInteractor is Interactor
type ChatInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	ImageInfrastructure    interfaces.ImageInfrastructure
	Logger                 interfaces.Logger
}

func (i ChatInteractor) ShowDefault(msg domain.SlackMessage) (err error) {
	// Send default Message To Slack
	sendMsg := fmt.Sprintf(DefaultMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) ShowHelp(msg domain.SlackMessage) (err error) {
	// Send help Message To Slack
	sendMsg := fmt.Sprintf(HelpMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) Pong(msg domain.SlackMessage) (err error) {
	// Send "pong" Message To Slack
	err = i.ChatInfrastructure.SendMessage("pong", msg.ChannelID)

	return
}

func (i ChatInteractor) SendImageWithReviewWaitTimeGraph(msg domain.SlackMessage) (err error) {
	if len(msg.Commands) < 4 {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessage(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]), msg.ChannelID)
	}
	u, err := url.Parse(strings.TrimRight(strings.TrimLeft(msg.Commands[2], "<"), ">"))
	if err != nil {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessage(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]), msg.ChannelID)
	}
	organization := strings.Split(u.Path, "/")[1]
	repository := strings.Split(u.Path, "/")[2]

	period, err := strconv.Atoi(msg.Commands[3])
	if err != nil {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessage(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]), msg.ChannelID)
	}

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
			// Send command-miss Message To Slack
			return i.ChatInfrastructure.SendMessage(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]), msg.ChannelID)
		}
		reviewWaitTimeMsg += fmt.Sprintf("%s\n> %v\n", issueURL, time)
	}
	err = i.ChatInfrastructure.SendMessage(reviewWaitTimeMsg, msg.ChannelID)
	if err != nil {
		return
	}

	// Create Bar Graph Image
	filepath, err := i.ImageInfrastructure.CreateGraphWithReviewWaitTime(times)
	if err != nil {
		return
	}

	// Send Bar Graph Image
	err = i.ChatInfrastructure.SendImage(filepath, msg.ChannelID)
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
