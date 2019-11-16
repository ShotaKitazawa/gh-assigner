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
	Title     string     `db:"title"`
	UserID    uint       `db:"user_id"`
	OpenedAt  *time.Time `db:"opened_at"`
	MergedAt  *time.Time `db:"merged_at"`
	CreatedAt *time.Time `db:"created_at"`
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
