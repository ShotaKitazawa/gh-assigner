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
	timestampFormat = "1970-01-01 00:00:00"
	requestLabel    = "review"
	reviewedLabel   = "wip"
)

// DatabaseInfrastructure is Infrastructure
type DatabaseInfrastructure struct {
	DB     *sqlx.DB
	Logger interfaces.Logger
}

func (r DatabaseInfrastructure) CreatePullRequest(userName, organizationName, repositoryName string, issueID uint, title string) (err error) {
	userID, err := r.CreateUserIfNotExists(userName)
	if err != nil {
		return
	}
	repositoryID, err := r.CreateRepositoryIfNotExists(organizationName, repositoryName)
	if err != nil {
		return
	}

	query := `INSERT INTO pullrequests (user_id, repository_id, issue_id, title, state) VALUES (?,?,?,?,"open")`
	_, err = r.DB.Exec(query, userID, repositoryID, issueID, title)
	return
}

func (r DatabaseInfrastructure) ClosePullRequest(userName, organizationName, repositoryName string, issueID uint, title string) (err error) {
	_, err = r.CreateUserIfNotExists(userName)
	if err != nil {
		return
	}
	repositoryID, err := r.CreateRepositoryIfNotExists(organizationName, repositoryName)
	if err != nil {
		return
	}

	query := `UPDATE pullrequests SET state = "close", closed_at = NOW() WHERE repository_id = ? AND issue_id = ? AND closed_at is NULL`
	_, err = r.DB.Exec(query, repositoryID, issueID)
	return
}

func (r DatabaseInfrastructure) MergePullRequest(userName, organizationName, repositoryName string, issueID uint, title string) (err error) {
	_, err = r.CreateUserIfNotExists(userName)
	if err != nil {
		return
	}
	repositoryID, err := r.CreateRepositoryIfNotExists(organizationName, repositoryName)
	if err != nil {
		return
	}

	query := `UPDATE pullrequests SET state = "merge", closed_at = NOW() WHERE repository_id = ? AND issue_id = ? AND closed_at is NULL`
	_, err = r.DB.Exec(query, repositoryID, issueID)
	return
}

func (r DatabaseInfrastructure) CreateRequestAction(userName, organizationName, repositoryName string, issueID uint) (err error) {
	userID, err := r.CreateUserIfNotExists(userName)
	if err != nil {
		return
	}
	repositoryID, err := r.CreateRepositoryIfNotExists(organizationName, repositoryName)
	if err != nil {
		return
	}

	// Start transaction
	tx := r.DB.MustBegin()

	// Get PullRequest state is not CLOSE/MERGE
	var pullrequests []PullRequest
	query := `SELECT id, state FROM pullrequests WHERE repository_id = ? AND issue_id = ? AND NOT (state = "close" OR state = "merge") FOR UPDATE`
	err = tx.Select(&pullrequests, query, repositoryID, issueID)
	if err != nil {
		tx.Rollback()
		return
	}
	if len(pullrequests) == 0 {
		tx.Rollback()
		err = fmt.Errorf("DB Data Mismatch: None of state OPEN PullRequest")
		return
	} else if len(pullrequests) > 1 {
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

	// Create Action record
	query = `INSERT INTO actions (pullreq_id, request_user_id) VALUES (?,?)`
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

func (r DatabaseInfrastructure) CreateReviewedAction(userName, organizationName, repositoryName string, issueID uint) (err error) {
	userID, err := r.CreateUserIfNotExists(userName)
	if err != nil {
		return
	}
	repositoryID, err := r.CreateRepositoryIfNotExists(organizationName, repositoryName)
	if err != nil {
		return
	}

	// Start transaction
	tx := r.DB.MustBegin()

	// Get PullRequest state is not CLOSE/MERGE
	var pullrequests []PullRequest
	query := `SELECT id, state FROM pullrequests WHERE repository_id = ? AND issue_id = ? AND NOT (state = "close" OR state = "merge") FOR UPDATE`
	err = tx.Select(&pullrequests, query, repositoryID, issueID)
	if err != nil {
		tx.Rollback()
		return
	}
	if len(pullrequests) == 0 {
		tx.Rollback()
		err = fmt.Errorf("DB Data Mismatch: None of state OPEN PullRequest")
		return
	} else if len(pullrequests) > 1 {
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

	// Update Action record
	query = `UPDATE actions SET review_user_id = ?, reviewed_at = NOW() WHERE pullreq_id = ? AND (review_user_id IS NULL AND reviewed_at IS NULL)`
	_, err = tx.Exec(query, userID, pullrequestID)
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

func (r DatabaseInfrastructure) CreateRepositoryIfNotExists(organizationName, repositoryName string) (repositoryID uint, err error) {
	query := `INSERT INTO repositories (organization, repository) VALUES (?,?)`
	result, err := r.DB.Exec(query, organizationName, repositoryName)
	if err != nil {
		// Ignore error 1062: Duplicate entry
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				query := `SELECT id FROM repositories WHERE organization = ? AND repository = ?`
				err = r.DB.Get(&repositoryID, query, organizationName, repositoryName)
			}
		}
		return
	}
	val, err := result.LastInsertId()
	repositoryID = uint(val)
	return
}

func (r DatabaseInfrastructure) GetPullRequestTTL(organizationName, repositoryName string, issueID uint) (duration time.Duration, err error) {
	query := `
SELECT a.requested_at, a.reviewed_at
FROM repositories AS r
JOIN pullrequests AS p ON r.id = p.repository_id
JOIN actions AS a ON p.id = a.pullreq_id
WHERE r.organization = ? AND r.repository = ? AND p.issue_id = ?
`

	timestamps := []struct {
		RequestedAt *time.Time `db:"requested_at"`
		ReviewedAt  *time.Time `db:"reviewed_at"`
	}{}
	err = r.DB.Select(&timestamps, query, organizationName, repositoryName, issueID)
	if err != nil {
		return
	}
	for _, timestamp := range timestamps {
		if timestamp.RequestedAt == nil || timestamp.ReviewedAt == nil {
			continue
		}
		duration += timestamp.ReviewedAt.Sub(*timestamp.RequestedAt)
	}

	return
}

func (r DatabaseInfrastructure) SelectPullRequestTTLs(organizationName, repositoryName string, period int) (durations []time.Duration, err error) {
	query := `
SELECT p.issue_id, a.requested_at, a.reviewed_at
FROM repositories AS r
JOIN pullrequests AS p ON r.id = p.repository_id
JOIN actions AS a ON p.id = a.pullreq_id
WHERE r.organization = ? AND r.repository = ? AND p.closed_at > ?
`

	timestamps := []struct {
		IssueID     uint       `db:"issue_id"`
		RequestedAt *time.Time `db:"requested_at"`
		ReviewedAt  *time.Time `db:"reviewed_at"`
	}{}
	err = r.DB.Select(&timestamps, query, organizationName, repositoryName, time.Now().AddDate(0, 0, -1*period).Format(timestampFormat))
	if err != nil {
		fmt.Println(err)
		return
	}
	durationsMap := make(map[uint]time.Duration, 128)
	for _, timestamp := range timestamps {
		if timestamp.RequestedAt == nil || timestamp.ReviewedAt == nil {
			continue
		}
		durationsMap[timestamp.IssueID] += timestamp.ReviewedAt.Sub(*timestamp.RequestedAt)
	}
	for _, val := range durationsMap {
		durations = append(durations, val)
	}

	return
}
