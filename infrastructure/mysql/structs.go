package mysql

import "time"

// User is MySQL users Table
type User struct {
	ID        uint       `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

// Repository is MySQL repositories Table
type Repository struct {
	ID           uint       `db:"id"`
	Organization string     `db:"organization"`
	Repository   string     `db:"repositories"`
	created_at   *time.Time `db:"created_at"`
}

// Issue is MySQL pullrequests Table
type PullRequest struct {
	ID           uint       `db:"id"`
	UserID       uint       `db:"user_id"`
	RepositoryID uint       `db:"repository_id"`
	IssueID      uint       `db:"issue_id"`
	Title        string     `db:"title"`
	State        string     `db:"state"`
	CreatedAt    *time.Time `db:"created_at"`
	ClosedAt     *time.Time `db:"closed_at"`
}

// RequestAction is MySQL request_actions Table
type RequestAction struct {
	ID        uint       `db:"id"`
	UserID    uint       `db:"user_id"`
	PullreqID uint       `db:"pullreq_id"`
	CreatedAt *time.Time `db:"created_at"`
}

// ReviewedAction is MySQL reviewed_actions Table
type ReviewedAction struct {
	ID        uint       `db:"id"`
	UserID    uint       `db:"user_id"`
	PullreqID uint       `db:"pullreq_id"`
	CreatedAt *time.Time `db:"created_at"`
}
