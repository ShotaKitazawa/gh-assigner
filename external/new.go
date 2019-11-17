package external

import (
	"context"
	"net/http"
	"time"

	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/mysql"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// NewGitHubWebhookController is initialize Controller, Interactor and Infrastructure.
func NewGitHubWebhookController(ctx context.Context) *controller.GitHubWebhookController {
	var err error

	// Get DB connection
	dsn := getContext(ctx, "dsn").(string)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to DB"))
	}

	// Get Github User & Token
	ghUser := getContext(ctx, "gh_user").(string)
	ghToken := getContext(ctx, "gh_token").(string)

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
