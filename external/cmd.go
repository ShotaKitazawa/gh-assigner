package external

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:           "gh-assigner",
	Short:         "SRE tools",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		// Set context
		ctx := context.Background()
		ctx = context.WithValue(ctx, hostContextKey,
			fmt.Sprintf("%s:%d",
				viper.GetString("bind-address"),
				viper.GetUint("bind-port"),
			),
		)
		ctx = context.WithValue(ctx, dsnContextKey,
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
				viper.GetString("db-user"),
				viper.GetString("db-password"),
				viper.GetString("db-host"),
				viper.GetUint("db-port"),
				viper.GetString("db-table-name"),
			))

		ctx = context.WithValue(ctx, ghUserContextKey, viper.GetString("github-user"))
		ctx = context.WithValue(ctx, ghTokenContextKey, viper.GetString("github-token"))
		ctx = context.WithValue(ctx, slackBotUserContextKey, viper.GetString("slack-bot-id"))
		ctx = context.WithValue(ctx, slackTokenContextKey, viper.GetString("slack-bot-token"))
		ctx = context.WithValue(ctx, slackChannelsContextKey, viper.GetString("slack-channel-ids"))
		ctx = context.WithValue(ctx, slackDefaultChannelContextKey, viper.GetString("slack-default-channel-id"))
		ctx = context.WithValue(ctx, googleCalendarContextKey, viper.GetString("google-calendar-id"))
		ctx = context.WithValue(ctx, gcpCredentialContextKey, viper.GetString("gcp-credential-path"))
		ctx = context.WithValue(ctx, crontabContextKey, viper.GetString("crontab"))

		// Run
		Run(ctx)
	},
}

// set up Cobra/Viper
func init() {

	var cfgFile string
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		}
		viper.AutomaticEnv()
		viper.ReadInConfig()
	})

	rootCmd.PersistentFlags().StringP("bind-address", "", "0.0.0.0", "Bind address")
	rootCmd.PersistentFlags().UintP("bind-port", "", 8080, "Bind port")
	rootCmd.PersistentFlags().StringP("db-user", "", "root", "User to connect DB")
	rootCmd.PersistentFlags().StringP("db-password", "", "password", "Password to connect DB")
	rootCmd.PersistentFlags().StringP("db-host", "", "127.0.0.1", "Host to connect DB")
	rootCmd.PersistentFlags().UintP("db-port", "", 3306, "Port to connect DB")
	rootCmd.PersistentFlags().StringP("db-table-name", "", "sample", "DB table name")
	rootCmd.PersistentFlags().StringP("github-user", "", "", "GitHub User")
	rootCmd.PersistentFlags().StringP("github-token", "", "", "GitHub token")
	rootCmd.PersistentFlags().StringP("slack-bot-id", "", "", "Slack Bot ID")
	rootCmd.PersistentFlags().StringP("slack-bot-token", "", "", "Slack Token for bot")
	rootCmd.PersistentFlags().StringP("slack-channel-ids", "", "", "Slack Channel IDs to activate bot (comma separated)")
	rootCmd.PersistentFlags().StringP("slack-default-channel-id", "", "", "Slack Channel ID to send message by bot initiatively")
	rootCmd.PersistentFlags().StringP("google-calendar-id", "", "", "GoogleCalendar ID")
	rootCmd.PersistentFlags().StringP("gcp-credential-path", "c", "", "Path to GCP Credential using to read GoogleCalendar")
	rootCmd.PersistentFlags().StringP("crontab", "", "", "Time when chat message is sent")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "Path to config file")
	viper.BindPFlag("bind-address", rootCmd.PersistentFlags().Lookup("bind-address"))
	viper.BindPFlag("bind-port", rootCmd.PersistentFlags().Lookup("bind-port"))
	viper.BindPFlag("db-user", rootCmd.PersistentFlags().Lookup("db-user"))
	viper.BindPFlag("db-password", rootCmd.PersistentFlags().Lookup("db-password"))
	viper.BindPFlag("db-host", rootCmd.PersistentFlags().Lookup("db-host"))
	viper.BindPFlag("db-port", rootCmd.PersistentFlags().Lookup("db-port"))
	viper.BindPFlag("db-table-name", rootCmd.PersistentFlags().Lookup("db-table-name"))
	viper.BindPFlag("github-user", rootCmd.PersistentFlags().Lookup("github-user"))
	viper.BindPFlag("github-token", rootCmd.PersistentFlags().Lookup("github-token"))
	viper.BindPFlag("slack-bot-id", rootCmd.PersistentFlags().Lookup("slack-bot-id"))
	viper.BindPFlag("slack-channel-ids", rootCmd.PersistentFlags().Lookup("slack-channel-ids"))
	viper.BindPFlag("slack-default-channel-id", rootCmd.PersistentFlags().Lookup("slack-default-channel-id"))
	viper.BindPFlag("slack-bot-token", rootCmd.PersistentFlags().Lookup("slack-bot-token"))
	viper.BindPFlag("google-calendar-id", rootCmd.PersistentFlags().Lookup("google-calendar-id"))
	viper.BindPFlag("gcp-credential-path", rootCmd.PersistentFlags().Lookup("gcp-credential-path"))
	viper.BindPFlag("crontab", rootCmd.PersistentFlags().Lookup("crontab"))
}

// Execute is entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
