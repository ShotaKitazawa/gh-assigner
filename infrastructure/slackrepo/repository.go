package slackrepo

import (
	"fmt"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
	"github.com/nlopes/slack"
)

type SlackInfrastructure struct {
	Client  *slack.Client
	Channel string
	Logger  interfaces.Logger
}

func (r SlackInfrastructure) SendMessage(msg, channel string) (err error) {
	if _, _, err := r.Client.PostMessage(channel, slack.MsgOptionText(msg, false)); err != nil {
		return err
	}
	return
}

func (r SlackInfrastructure) SendMessageToDefaultChannel(msg string) (err error) {
	fmt.Println(r.Channel)
	if _, _, err := r.Client.PostMessage(r.Channel, slack.MsgOptionText(msg, false)); err != nil {
		return err
	}
	return
}
