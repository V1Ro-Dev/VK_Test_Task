package config

import (
	"fmt"
	"log"

	"github.com/mattermost/mattermost-server/v6/model"
)

type Bot struct {
	Config                    MatterMost
	MattermostClient          *model.Client4
	MattermostWebSocketClient *model.WebSocketClient
	MattermostUser            *model.User
}

func NewBot(cfg MatterMost) (*Bot, error) {
	client := model.NewAPIv4Client(cfg.MattermostServer)

	client.SetToken(cfg.MattermostToken)

	user, _, err := client.GetMe("")
	if err != nil {
		return &Bot{}, fmt.Errorf("failed to get user from mattermost: %w", err)
	}

	return &Bot{
		Config:                    cfg,
		MattermostClient:          client,
		MattermostWebSocketClient: &model.WebSocketClient{},
		MattermostUser:            user,
	}, nil
}

func (b *Bot) SendMessage(channelID string, text string) {
	post := &model.Post{
		ChannelId: channelID,
		Message:   text,
	}
	if _, _, err := b.MattermostClient.CreatePost(post); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
