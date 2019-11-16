package mysql

import "time"

// User is MySQL users Table
type User struct {
	ID        uint       `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

// Issue is MySQL pullrequests Table
type Issue struct {
	ID        uint       `db:"id"`
	IssueID   uint       `db:"issue_id"`
	UserID    uint       `db:"user_id"`
	Title     string     `db:"title"`
	URL       string     `db:"url"`
	State     string     `db:"state"`
	CreatedAt *time.Time `db:"created_at"`
	ClosedAt  *time.Time `db:"closed_at"`
}

// RequestAction is MySQL request_actions Table
type RequestAction struct {
	ID        uint       `db:"id"`
	PullreqID uint       `db:"pullreq_id"`
	UserID    uint       `db:"user_id"`
	CreatedAt *time.Time `db:"created_at"`
}

// ReviewedAction is MySQL reviewed_actions Table
type ReviewedAction struct {
	ID        uint       `db:"id"`
	PullreqID uint       `db:"pullreq_id"`
	UserID    uint       `db:"user_id"`
	CreatedAt *time.Time `db:"created_at"`
}
