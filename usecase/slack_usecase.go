package usecase

import (
	"fmt"
	"strconv"

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

func (i ChatInteractor) Pong(msg domain.SlackMessage) (err error) {
	err = i.ChatInfrastructure.SendMessage("pong", msg.ChannelID)

	return
}

func (i ChatInteractor) ShowDefault(msg domain.SlackMessage) (err error) {
	sendMsg := fmt.Sprintf(DefaultMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) ShowHelp(msg domain.SlackMessage) (err error) {
	sendMsg := fmt.Sprintf(HelpMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) SendImageWithReviewWaitTimeGraph(msg domain.SlackMessage) (err error) {
	if len(msg.Commands) < 5 {
		err = i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
		return
	}
	organization := msg.Commands[2]
	repository := msg.Commands[3]
	period, err := strconv.Atoi(msg.Commands[4])
	if err != nil {
		err = i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
		return
	}

	times, err := i.DatabaseInfrastructure.SelectPullRequestTTLs(organization, repository, period)
	if err != nil {
		return
	}

	filepath, err := i.ImageInfrastructure.CreateGraphWithReviewWaitTime(times)
	if err != nil {
		return
	}
	err = i.ChatInfrastructure.SendImage(filepath, msg.ChannelID)
	if err != nil {
		return
	}
	err = i.ImageInfrastructure.DeleteFile(filepath)
	if err != nil {
		return
	}
	return
}
