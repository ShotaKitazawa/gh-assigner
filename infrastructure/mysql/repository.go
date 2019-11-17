package mysql

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ShotaKitazawa/gh-assigner/infrastructure/interfaces"
)

const (
	requestLabel  = "review"
	reviewedLabel = "wip"
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
	// Start transaction
	tx := r.DB.MustBegin()

	// Get PullRequest state is not CLOSE/MERGE
	var pullrequests []PullRequest
	query := `SELECT id, state FROM pullrequests WHERE repository_id = ? AND issue_id = ? AND NOT (state = "close" OR state = "merge") FOR UPDATE`
	err = r.DB.Select(&pullrequests, query, repositoryID, issueID)
	if err != nil {
		tx.Rollback()
		return
	}
	if len(pullrequests) > 1 {
		tx.Rollback()
		err = fmt.Errorf("DB Data Mismatch: Multiple state OPEN PullRequest")
		return
	}

	// Return when PullRequest state has already been `requestLabel`
	if pullrequests[0].State == requestLabel {
		tx.Rollback()
		r.Logger.Debug(fmt.Sprintf("PullRequest state has already been %s", strings.ToUpper(requestLabel)))
		return
	}

	pullrequestID := pullrequests[0].ID

	// Create RequestAction record
	query = `INSERT INTO request_actions (pullreq_id, user_id) VALUES (?,?)`
	_, err = tx.Exec(query, pullrequestID, userID)
	if err != nil {
		tx.Rollback()
		return
	}

	// Update state Calumn in pullrequests Table
	query = `UPDATE pullrequests SET state = ? WHERE id = ?`
	_, err = tx.Exec(query, requestLabel, pullrequestID)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit transaction
	tx.Commit()

	return err
}

func (r DatabaseInfrastructure) CreateReviewedAction(userID, repositoryID, issueID uint) (err error) {
	// Start transaction
	tx := r.DB.MustBegin()

	// Get PullRequest state is not CLOSE/MERGE
	var pullrequests []PullRequest
	query := `SELECT id, state FROM pullrequests WHERE repository_id = ? AND issue_id = ? AND NOT (state = "close" OR state = "merge") FOR UPDATE`
	err = r.DB.Select(&pullrequests, query, repositoryID, issueID)
	if err != nil {
		tx.Rollback()
		return
	}
	if len(pullrequests) > 1 {
		tx.Rollback()
		err = fmt.Errorf("DB Data Mismatch: Multiple state OPEN PullRequest")
		return
	}

	// Return when PullRequest state has already been `reviewedLabel`
	if pullrequests[0].State == reviewedLabel {
		tx.Rollback()
		r.Logger.Debug(fmt.Sprintf("PullRequest state has already been %s", strings.ToUpper(reviewedLabel)))
		return
	}

	pullrequestID := pullrequests[0].ID

	// Create RequestAction record
	query = `INSERT INTO reviewed_actions (pullreq_id, user_id) VALUES (?,?)`
	_, err = tx.Exec(query, pullrequestID, userID)
	if err != nil {
		tx.Rollback()
		return
	}

	// Update state Calumn in pullrequests Table
	query = `UPDATE pullrequests SET state = ? WHERE id = ?`
	_, err = tx.Exec(query, reviewedLabel, pullrequestID)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit transaction
	tx.Commit()

	return err
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
