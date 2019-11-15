package external

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/mysql"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

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
			GitRepository: &github.GitRepository{
				User:   ghUser,
				Token:  ghToken,
				Logger: &Logger{},
			},
			DatabaseRepository: &mysql.DatabaseRepository{
				DB:     db,
				Logger: &Logger{},
			},
			/*
				CalendarRepository: &googlecalendar.CalendarRepository{
					Credential: TODO,
				}
			*/
			Logger: &Logger{},
		},
	}
}
