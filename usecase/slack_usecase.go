package usecase

import (
	"fmt"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

const (
	CommandHelp = "help"
	CommandPing = "ping"
)

var (
	DefaultMessage = fmt.Sprintf(`
コマンドが存在しません。 以下のコマンドよりヘルプメッセージを確認してください。
> %s
`, "`%[1]s "+CommandHelp+"`")

	HelpMessage = fmt.Sprintf(`
> %s
疎通確認
> %s
ヘルプ出力
`,
		"`%[1]s "+CommandPing+"`",
		"`%[1]s "+CommandHelp+"`",
	)
)

// ChatInteractor is Interactor
type ChatInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	Logger                 interfaces.Logger
}

func (i ChatInteractor) Pong(msg domain.SlackMessage) (err error) {
	err = i.ChatInfrastructure.SendMessage("pong", msg.ChannelID)

	return
}
func (i ChatInteractor) ShowDefault(msg domain.SlackMessage) (err error) {
	sendMsg := fmt.Sprintf(DefaultMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}
func (i ChatInteractor) ShowHelp(msg domain.SlackMessage) (err error) {
	sendMsg := fmt.Sprintf(HelpMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}
