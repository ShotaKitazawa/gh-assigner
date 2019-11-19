package external

import (
	"context"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/mysql"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/slackrtm"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
)

var (
	db          *sqlx.DB
	ghUser      string
	ghToken     string
	slackClient *slack.Client
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
	token, err := getContextString(ctx, slackTokenContextKey)
	if err != nil {
		panic(err)
	}
	slackClient = slack.New(token)

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
			ChatInfrastructure: &slackrtm.SlackInfrastructure{
				Client: slackClient,
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
		Interactor: &usecase.ChatInteractor{
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
			ChatInfrastructure: &slackrtm.SlackInfrastructure{
				Client: slackClient,
				Logger: &Logger{},
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
