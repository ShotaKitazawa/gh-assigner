package slackrtm

import (
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
	"github.com/nlopes/slack"
)

type SlackInfrastructure struct {
	Client *slack.Client
	Logger interfaces.Logger
}

func (r SlackInfrastructure) HelloMessage(msg, channel string) (err error) {
	if _, _, err := r.Client.PostMessage(channel, slack.MsgOptionText("Hello World!", false)); err != nil {
		return err
	}
	return
}
