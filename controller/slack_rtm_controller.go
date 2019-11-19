package controller

import (
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
	//TODO 引数を domain に寄せる
	err = c.Interactor.Hello(domain.SlackMessage{
		ChannelID: ev.Channel,
		Message:   ev.Msg.Text,
	})
	return err
}
