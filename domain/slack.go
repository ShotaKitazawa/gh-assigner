package domain

type SlackMessage struct {
	ChannelID string
	SenderName string
	Commands   []string
}
