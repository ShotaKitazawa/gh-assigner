package usecase

import (
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
	"github.com/nlopes/slack"
)

// SlackInteractor is Interactor
type SlackInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	SlackInfrastructure    interfaces.SlackInfrastructure
	Logger                 interfaces.Logger
}

func (i SlackInteractor) Hello(ev *slack.MessageEvent) (err error) {
	err = i.SlackInfrastructure.HelloMessage("hello", ev.Channel)
	return
}
