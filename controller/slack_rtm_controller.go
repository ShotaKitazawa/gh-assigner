package controller

import (
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/nlopes/slack"
)

// GitHubWebhookController is Controller
type SlackRTMController struct {
	Interactor interfaces.ChatInteractor
	Logger     interfaces.Logger
}

func (c SlackRTMController) MessageEvent(ev *slack.MessageEvent) (err error) {
	commands := strings.Split(trimNewlineChar(ev.Msg.Text), " ")
	err = c.Interactor.Hello(domain.SlackMessage{
		ChannelID:  ev.Channel,
		SenderName: ev.User,
		Commands:   commands,
	})
	return err
}
