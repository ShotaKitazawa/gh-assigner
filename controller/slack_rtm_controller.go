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
	switch commands[1] {
	case "ping":
		err = c.Interactor.Pong(domain.SlackMessage{
			ChannelID:  ev.Channel,
			SenderName: ev.User,
			Commands:   commands,
		})
		if err != nil {
			c.Logger.Error(err.Error())
			return
		}
	case "help":
		err = c.Interactor.ShowHelp(domain.SlackMessage{
			ChannelID:  ev.Channel,
			SenderName: ev.User,
			Commands:   commands,
		})
		if err != nil {
			c.Logger.Error(err.Error())
			return
		}
	default:
		err = c.Interactor.ShowDefault(domain.SlackMessage{
			ChannelID:  ev.Channel,
			SenderName: ev.User,
			Commands:   commands,
		})
		if err != nil {
			c.Logger.Error(err.Error())
			return
		}
	}
	return
}
