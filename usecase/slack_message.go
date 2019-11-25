package usecase

import "fmt"

const (
	CommandPing            = "ping"
	CommandReviewTTL       = "review-ttl"
	CommandReviewTTLOption = "<OrganizationName> <RepositoryName> <Period>"
	CommandHelp            = "help"
)

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
