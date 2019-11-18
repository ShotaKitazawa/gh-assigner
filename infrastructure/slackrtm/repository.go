package slackrtm

import (
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
	"github.com/ShotaKitazawa/gh-assigner/pkg/slackutils"
	"github.com/nlopes/slack"
)

type SlackInfrastructure struct {
	Listener *slackutils.Listener
	Logger   interfaces.Logger
}

func (r SlackInfrastructure) HelloMessage(msg, channel string) (err error) {
	if _, _, err := r.Listener.Client.PostMessage(channel, slack.MsgOptionText("Hello World!", false)); err != nil {
		return err
	}
	return
}
