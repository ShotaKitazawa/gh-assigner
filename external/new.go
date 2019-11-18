package external

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/mysql"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/slackrtm"
	"github.com/ShotaKitazawa/gh-assigner/pkg/slackutils"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
)

var (
	db          *sqlx.DB
	ghUser      string
	ghToken     string
	slackClient *slackutils.Listener
)

// Initialize is initialize shared setting with all Controller
func Initialize(ctx context.Context) func() {
	// Get DB connection
	dsn, err := getContextString(ctx, dsnContextKey)
	if err != nil {
		panic(err)
	}
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to DB"))
	}

	// Get Github User & Token
	ghUser, err = getContextString(ctx, ghUserContextKey)
	if err != nil {
		panic(err)
	}
	ghToken, err = getContextString(ctx, ghTokenContextKey)
	if err != nil {
		panic(err)
	}

	// Get Slack Channel & Token
	channelsStr, err := getContextString(ctx, slackChannelsContextKey)
	if err != nil {
		panic(err)
	}
	channels := strings.Split(channelsStr, ",")
	botUsername, err := getContextString(ctx, slackBotUserContextKey)
	if err != nil {
		panic(err)
	}
	token, err := getContextString(ctx, slackTokenContextKey)
	if err != nil {
		panic(err)
	}
	slackClient = slackutils.New(token, channels, botUsername)

	return func() {
		db.Close()
	}
}

// NewGitHubWebhookController is initialize Controller, Interactor and Infrastructure.
func NewGitHubWebhookController(ctx context.Context) *controller.GitHubWebhookController {
	return &controller.GitHubWebhookController{
		Interactor: &usecase.GitInteractor{
			GitInfrastructure: &github.GitInfrastructure{
				Client: &http.Client{Timeout: time.Duration(10) * time.Second},
				User:   ghUser,
				Token:  ghToken,
				Logger: &Logger{},
			},
			DatabaseInfrastructure: &mysql.DatabaseInfrastructure{
				DB:     db,
				Logger: &Logger{},
			},
			/*
				CalendarInfrastructure: &googlecalendar.CalendarInfrastructure{
					Credential: TODO,
				}
			*/
			Logger: &Logger{},
		},
		Logger: &Logger{},
	}
}

// NewSlackRTMController is initialize Controller, Interactor and Infrastructure.
func NewSlackRTMController(ctx context.Context) *controller.SlackRTMController {
	return &controller.SlackRTMController{
		Interactor: &usecase.SlackInteractor{
			GitInfrastructure: &github.GitInfrastructure{
				Client: &http.Client{Timeout: time.Duration(10) * time.Second},
				User:   ghUser,
				Token:  ghToken,
				Logger: &Logger{},
			},
			DatabaseInfrastructure: &mysql.DatabaseInfrastructure{
				DB:     db,
				Logger: &Logger{},
			},
			SlackInfrastructure: &slackrtm.SlackInfrastructure{
				Listener: slackClient,
				Logger:   &Logger{},
			},
			/*
				CalendarInfrastructure: &googlecalendar.CalendarInfrastructure{
					Credential: TODO,
				}
			*/
			Logger: &Logger{},
		},
	}
}
