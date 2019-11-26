package usecase

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/ShotaKitazawa/gh-assigner/domain"
	"github.com/ShotaKitazawa/gh-assigner/usecase/interfaces"
)

// Commands & Options
const (
	CommandPing            = "ping"
	CommandReviewTTL       = "review-ttl"
	CommandReviewTTLOption = "<RepositoryURL> <Period>"
	CommandHelp            = "help"
)

// message displayed to Chat
var (
	DefaultMessage = fmt.Sprintf(`
コマンドが存在しません。 以下のコマンドよりヘルプメッセージを確認してください。
> %[1]s%[2]s%[1]s
`, "`", "%[1]s "+CommandHelp)

	invalidCommandSlackMessage = fmt.Sprintf(`
コマンドの引数が誤っています。以下のコマンドよりヘルプメッセージを確認してください。
> %[1]s%[2]s%[1]s
`, "`", "%[1]s "+CommandHelp)

	HelpMessage = fmt.Sprintf(`
> %[1]s%[2]s%[1]s
疎通確認
> %[1]s%[3]s%[1]s
<Period> 日前までにCloseしたPullRequestにおけるレビュー時間の取得
> %[1]s%[4]s%[1]s
ヘルプ出力
`, "`",
		"`%[1]s "+CommandPing+"`",
		"`%[1]s "+CommandReviewTTL+" "+CommandReviewTTLOption+"`",
		"`%[1]s "+CommandHelp+"`",
	)
)

// ChatInteractor is Interactor
type ChatInteractor struct {
	GitInfrastructure      interfaces.GitInfrastructure
	DatabaseInfrastructure interfaces.DatabaseInfrastructure
	CalendarInfrastructure interfaces.CalendarInfrastructure
	ChatInfrastructure     interfaces.ChatInfrastructure
	ImageInfrastructure    interfaces.ImageInfrastructure
	Logger                 interfaces.Logger
}

func (i ChatInteractor) ShowDefault(msg domain.SlackMessage) (err error) {
	// Send default Message To Slack
	sendMsg := fmt.Sprintf(DefaultMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) ShowHelp(msg domain.SlackMessage) (err error) {
	// Send help Message To Slack
	sendMsg := fmt.Sprintf(HelpMessage, msg.Commands[0])
	err = i.ChatInfrastructure.SendMessage(sendMsg, msg.ChannelID)

	return
}

func (i ChatInteractor) Pong(msg domain.SlackMessage) (err error) {
	// Send "pong" Message To Slack
	err = i.ChatInfrastructure.SendMessage("pong", msg.ChannelID)

	return
}

func (i ChatInteractor) SendImageWithReviewWaitTimeGraph(msg domain.SlackMessage) (err error) {
	if len(msg.Commands) < 5 {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
	}
	//General form
	u, err := url.Parse(msg.Commands[2])
	if err != nil {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
	}
	organization := strings.Split(u.Path, "/")[1]
	repository := strings.Split(u.Path, "/")[2]

	period, err := strconv.Atoi(msg.Commands[3])
	if err != nil {
		// Send command-miss Message To Slack
		return i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
	}

	// Get PullRequest TTL last `period` days
	times, err := i.DatabaseInfrastructure.SelectPullRequestTTLs(organization, repository, period)
	if err != nil {
		return
	}

	// Send TTL info
	var reviewWaitTimeMsg string
	for id, time := range times {
		issueURL, err := i.DatabaseInfrastructure.GetPullRequestURL(organization, repository, id)
		if err != nil {
			// Send command-miss Message To Slack
			return i.ChatInfrastructure.SendMessageToDefaultChannel(fmt.Sprintf(invalidCommandSlackMessage, msg.Commands[0]))
		}
		reviewWaitTimeMsg += fmt.Sprintf("> %s\n%v\n", issueURL, time)
	}
	err = i.ChatInfrastructure.SendMessage(reviewWaitTimeMsg, msg.ChannelID)
	if err != nil {
		return
	}

	// Create Bar Graph Image
	filepath, err := i.ImageInfrastructure.CreateGraphWithReviewWaitTime(times)
	if err != nil {
		return
	}

	// Send Bar Graph Image
	err = i.ChatInfrastructure.SendImage(filepath, msg.ChannelID)
	if err != nil {
		return
	}

	// Delete Bar Graph Image
	err = i.ImageInfrastructure.DeleteFile(filepath)
	if err != nil {
		return
	}

	return
}
