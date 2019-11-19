package usecase

import (
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
	"github.com/nlopes/slack"
)

// ChatInteractor is Interactor
type ChatInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	Logger                 interfaces.Logger
}

func (i ChatInteractor) Hello(ev *slack.MessageEvent) (err error) {
	err = i.ChatInfrastructure.HelloMessage("hello", ev.Channel)
	return
}
