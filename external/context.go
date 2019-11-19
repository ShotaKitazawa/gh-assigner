package external

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type contextKey string

const (
	hostContextKey                contextKey = "host"
	dsnContextKey                 contextKey = "dsn"
	ghUserContextKey              contextKey = "gh_user"
	ghTokenContextKey             contextKey = "gh_token"
	slackChannelsContextKey       contextKey = "slack_channels"
	slackBotUserContextKey        contextKey = "slack_botuser"
	slackDefaultChannelContextKey contextKey = "slack_default_channel"
	slackTokenContextKey          contextKey = "slack_token"
	gcpCredentialContextKey       contextKey = "gcp_credential"
	googleCalendarContextKey      contextKey = "google_calendar"
)

func getContextString(ctx context.Context, key contextKey) (result string, err error) {
	inter := ctx.Value(key)
	if inter == nil {
		err = errors.New(fmt.Sprintf("context not in value %s", key))
		return
	}
	result, ok := inter.(string)
	if !ok {
		err = errors.New(fmt.Sprintf("value %s is not string", key))
		return
	}
	return
}
