package external

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/googlecalendar"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/image"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/mysql"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/slackrepo"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
)

var (
	db                  *sqlx.DB
	ghUser              string
	ghToken             string
	slackClient         *slack.Client
	slackDefaultChannel string
	calendarID          string
	calendarService     *calendar.Service
	currentPath         string
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

	// Get Slack DefaultChannel & Token
	slackDefaultChannel, err = getContextString(ctx, slackDefaultChannelContextKey)
	if err != nil {
		panic(err)
	}
	token, err := getContextString(ctx, slackTokenContextKey)
	if err != nil {
		panic(err)
	}
	slackClient = slack.New(token)

	// Get GoogleCalendar ID & Service
	calendarID, err = getContextString(ctx, googleCalendarContextKey)
	if err != nil {
		panic(err)
	}
	gcpCredentialPath, err := getContextString(ctx, gcpCredentialContextKey)
	if err != nil {
		panic(err)
	}
	calendarCtx := context.Background()
	calendarService, err = calendar.NewService(calendarCtx, option.WithCredentialsFile(gcpCredentialPath))

	// Get current path
	currentPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

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
			ChatInfrastructure: &slackrepo.ChatInfrastructure{
				Client:  slackClient,
				Channel: slackDefaultChannel,
				Logger:  &Logger{},
			},
			CalendarInfrastructure: &googlecalendar.CalendarInfrastructure{
				ID:      calendarID,
				Service: calendarService,
				Logger:  &Logger{},
			},
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
			ChatInfrastructure: &slackrepo.ChatInfrastructure{
				Client:  slackClient,
				Channel: slackDefaultChannel,
				Logger:  &Logger{},
			},
			CalendarInfrastructure: &googlecalendar.CalendarInfrastructure{
				ID:      calendarID,
				Service: calendarService,
				Logger:  &Logger{},
			},
			ImageInfrastructure: &image.ImageInfrastructure{
				Path: currentPath + "/images",
			},
			Logger: &Logger{},
		},
		Logger: &Logger{},
	}
}

// NewCronController is initialize Controller, Interactor and Infrastructure.
func NewCronController(ctx context.Context) *controller.CronController {
	return &controller.CronController{
		Interactor: &usecase.CronInteractor{
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
			ChatInfrastructure: &slackrepo.ChatInfrastructure{
				Client:  slackClient,
				Channel: slackDefaultChannel,
				Logger:  &Logger{},
			},
			CalendarInfrastructure: &googlecalendar.CalendarInfrastructure{
				ID:      calendarID,
				Service: calendarService,
				Logger:  &Logger{},
			},
			ImageInfrastructure: &image.ImageInfrastructure{
				Path: currentPath + "/images",
			},
			Logger: &Logger{},
		},
		Logger: &Logger{},
	}
}
