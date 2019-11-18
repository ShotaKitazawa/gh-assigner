package slackutils

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

type Listener struct {
	Client     *slack.Client
	ChannelIDs []string
	BotUser    string
}

func New(token string, channels []string, botUser string) *Listener {
	return &Listener{
		Client:     slack.New(token),
		ChannelIDs: channels,
		BotUser:    botUser,
	}
}

func (s *Listener) Run(f func(ev *slack.MessageEvent)) error {
	rtm := s.Client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			for _, channelName := range s.ChannelIDs {
				if ev.Channel == channelName {
					// for debug
					fmt.Println(ev.Msg.Text)
					fmt.Println(fmt.Sprintf("<@%s> ", ev.BotID))

					// Only response mention to bot. Ignore else.
					if strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.BotUser)) {
						f(ev)
					}
					break
				}
			}
		}
	}
	return nil
}
