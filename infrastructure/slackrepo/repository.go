package slackrepo

import (
	"fmt"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
	"github.com/nlopes/slack"
)

// ChatInfrastructure is Repository
type ChatInfrastructure struct {
	Client  *slack.Client
	Channel string
	Logger  interfaces.Logger
}

func (r ChatInfrastructure) SendMessage(msg, channel string) (err error) {
	if _, _, err := r.Client.PostMessage(channel, slack.MsgOptionText(msg, false)); err != nil {
		return err
	}
	return
}

func (r ChatInfrastructure) SendMessageToDefaultChannel(msg string) (err error) {
	if _, _, err := r.Client.PostMessage(r.Channel, slack.MsgOptionText(msg, false)); err != nil {
		return err
	}
	return
}
