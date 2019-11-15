package mysql

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

// DatabaseRepository is Repository
type DatabaseRepository struct {
	DB     *sqlx.DB
	Logger interfaces.Logger
}

func (r DatabaseRepository) CreatePullRequest(issueID int, title, username, url string) error {
	// TODO
	return nil
}
func (r DatabaseRepository) CreateRequestAction(issueID int, username string) error {
	// TODO
	return nil
}
func (r DatabaseRepository) CreateReviewedAction(issueID int, username string) error {
	// TODO
	return nil
}
func (r DatabaseRepository) CreateUser(username string) error {
	// TODO
	return nil
}
func (r DatabaseRepository) GetPullRequestTTL(issueID int) (time.Duration, error) {
	// TODO
	return 0, nil
}
