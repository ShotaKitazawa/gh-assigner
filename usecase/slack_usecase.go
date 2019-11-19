package usecase

import (
	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

// ChatInteractor is Interactor
type ChatInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	Logger                 interfaces.Logger
}

func (i ChatInteractor) Hello(msg domain.SlackMessage) (err error) {
	err = i.ChatInfrastructure.HelloMessage("hello", msg.ChannelID)
	return
}
