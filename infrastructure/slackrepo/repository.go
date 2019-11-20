package slackrepo

import (
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
	_, _, err = r.Client.PostMessage(channel, slack.MsgOptionText(msg, false))
	return
}

func (r ChatInfrastructure) SendMessageToDefaultChannel(msg string) (err error) {
	return r.SendMessage(msg, r.Channel)
}

func (r ChatInfrastructure) SendImage(filepath, channel string) (err error) {
	params := slack.FileUploadParameters{
		Filetype: "image",
		File:     filepath,
		Channels: []string{channel},
	}
	_, err = r.Client.UploadFile(params)
	return
}

func (r ChatInfrastructure) SendImageToDefaultChannel(filepath string) (err error) {
	return r.SendImage(filepath, r.Channel)
}
