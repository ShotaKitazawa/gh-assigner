package interfaces

import (
	"context"
)

type GitHubRepository interface {
	PostMessage(context.Context, string) error
}
