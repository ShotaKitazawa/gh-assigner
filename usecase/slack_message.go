package usecase

import "fmt"

const (
	CommandPing      = "ping"
	CommandReviewTTL = "review-ttl"
	CommandHelp      = "help"
)

var (
	DefaultMessage = fmt.Sprintf(`
コマンドが存在しません。 以下のコマンドよりヘルプメッセージを確認してください。
> %s
`, "`%[1]s "+CommandHelp+"`")

	invalidCommandSlackMessage = fmt.Sprintf(`
コマンドの引数が誤っています。以下のコマンドよりヘルプメッセージを確認してください。
> %s
`, "`%[1]s "+CommandHelp+"`")

	HelpMessage = fmt.Sprintf(`
> %s
疎通確認
> %s <OrganizationName> <RepositoryName> <Period>
<Period> 日前までにCloseしたPullRequestにおけるレビュー時間の取得
> %s
ヘルプ出力
`,
		"`%[1]s "+CommandPing+"`",
		"`%[1]s "+CommandReviewTTL+"`",
		"`%[1]s "+CommandHelp+"`",
	)
)
