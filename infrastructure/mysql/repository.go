package mysql

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

// DatabaseInfrastructure is Infrastructure
type DatabaseInfrastructure struct {
	DB     *sqlx.DB
	Logger interfaces.Logger
}

func (r DatabaseInfrastructure) CreatePullRequest(userID, repositoryID, issueID uint, title string) (err error) {
	query := `INSERT INTO pullrequests (user_id, repository_id, issue_id, title, state) VALUES (?,?,?,?,"open")`
	_, err = r.DB.Exec(query, userID, repositoryID, issueID, title)
	return
}

func (r DatabaseInfrastructure) CreateRequestAction(userID, repositoryID, issueID uint) (err error) {
	pullrequestID, err := r.GetPullRequestIDStateOpen(repositoryID, issueID)
	if err != nil {
		return
	}

	query := `INSERT INTO request_actions (pullreq_id, user_id) VALUES (?,?)`
	_, err = r.DB.Exec(query, pullrequestID, userID)
	return err
}

func (r DatabaseInfrastructure) CreateReviewedAction(userID, repositoryID, issueID uint) (err error) {
	pullrequestID, err := r.GetPullRequestIDStateOpen(repositoryID, issueID)
	if err != nil {
		return
	}

	query := `INSERT INTO reviewed_actions (pullreq_id, user_id) VALUES (?,?)`
	_, err = r.DB.Exec(query, pullrequestID, userID)
	return err
}

func (r DatabaseInfrastructure) GetPullRequestIDStateOpen(repositoryID, issueID uint) (pullrequestID uint, err error) {
	var pullrequestIDs []uint

	query := `SELECT id FROM pullrequests WHERE repository_id = ? AND issue_id = ? AND state = "open"`
	err = r.DB.Select(&pullrequestIDs, query, repositoryID, issueID)
	if err != nil {
		return
	}
	if len(pullrequestIDs) > 1 {
		err = fmt.Errorf("DB Data Mismatch: Multiple state OPEN PullRequest")
		return
	}
	pullrequestID = pullrequestIDs[0]
	return
}

func (r DatabaseInfrastructure) CreateUserIfNotExists(username string) (userID uint, err error) {
	query := `INSERT INTO users (name) VALUES (?)`
	result, err := r.DB.Exec(query, username)
	if err != nil {
		// Ignore error 1062: Duplicate entry
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				query := `SELECT id FROM users WHERE name = ?`
				err = r.DB.Get(&userID, query, username)
			}
		}
		return
	}
	val, err := result.LastInsertId()
	userID = uint(val)
	return
}

func (r DatabaseInfrastructure) CreateRepositoryIfNotExists(organization, repository string) (repositoryID uint, err error) {
	query := `INSERT INTO repositories (organization, repository) VALUES (?,?)`
	result, err := r.DB.Exec(query, organization, repository)
	if err != nil {
		// Ignore error 1062: Duplicate entry
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				query := `SELECT id FROM repositories WHERE organization = ? AND repository = ?`
				err = r.DB.Get(&repositoryID, query, organization, repository)
			}
		}
		return
	}
	val, err := result.LastInsertId()
	repositoryID = uint(val)
	return
}

func (r DatabaseInfrastructure) GetPullRequestTTL(issueID uint) (time.Duration, error) {
	// TODO
	return 0, nil
}
