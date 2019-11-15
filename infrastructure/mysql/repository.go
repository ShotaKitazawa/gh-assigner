package mysql

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

type DatabaseRepository struct {
	DB     *sqlx.DB
	Logger interfaces.Logger
}

func (r DatabaseRepository) CreatePullRequest(issue_id int, title, username, url string) error {
	return nil
}
func (r DatabaseRepository) CreateRequestAction(issue_id int, username string) error {
	return nil
}
func (r DatabaseRepository) CreateReviewedAction(issue_id int, username string) error {
	return nil
}
func (r DatabaseRepository) CreateUser(username string) error {
	return nil
}
func (r DatabaseRepository) GetPullRequestTTL(issue_id int) (time.Duration, error) {
	return 0, nil
}
