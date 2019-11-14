package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ShotaKitazawa/gh-assigner/external"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ctx context.Context

var rootCmd = &cobra.Command{
	Use:           "gh-assigner",
	Short:         "SRE tools",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		// Set context
		host := fmt.Sprintf("%s:%d",
			viper.GetString("bind-address"),
			viper.GetUint("bind-port"),
		)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			viper.GetString("db-user"),
			viper.GetString("db-password"),
			viper.GetString("db-host"),
			viper.GetUint("db-port"),
			"sample",
		)
		ctx = context.Background()
		ctx = context.WithValue(ctx, "host", host)
		ctx = context.WithValue(ctx, "dsn", dsn)
		ctx = context.WithValue(ctx, "gh_user", viper.GetString("github-user"))
		ctx = context.WithValue(ctx, "gh_token", viper.GetString("github-token"))

		// Run
		external.Run(ctx)
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
	rootCmd.PersistentFlags().StringP("github-user", "", "", "GitHub User")
	rootCmd.PersistentFlags().StringP("github-token", "", "", "GitHub token")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	viper.BindPFlag("bind-address", rootCmd.PersistentFlags().Lookup("bind-address"))
	viper.BindPFlag("bind-port", rootCmd.PersistentFlags().Lookup("bind-port"))
	viper.BindPFlag("db-user", rootCmd.PersistentFlags().Lookup("db-user"))
	viper.BindPFlag("db-password", rootCmd.PersistentFlags().Lookup("db-password"))
	viper.BindPFlag("db-host", rootCmd.PersistentFlags().Lookup("db-host"))
	viper.BindPFlag("db-port", rootCmd.PersistentFlags().Lookup("db-port"))
	viper.BindPFlag("github-user", rootCmd.PersistentFlags().Lookup("github-user"))
	viper.BindPFlag("github-token", rootCmd.PersistentFlags().Lookup("github-token"))

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
