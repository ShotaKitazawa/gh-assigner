package controller

import (
	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
	"github.com/nlopes/slack"
)

// GitHubWebhookController is Controller
type SlackRTMController struct {
	Interactor interfaces.SlackInteractor
	Logger     interfaces.Logger
}

func (c SlackRTMController) MessageEvent(ev *slack.MessageEvent) (err error) {
	//TODO 引数を domain に寄せる
	err = c.Interactor.Hello(ev)
	return err
}
