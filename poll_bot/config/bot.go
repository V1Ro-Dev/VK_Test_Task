package config

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v6/model"
	"log"
)

type Bot struct {
	Config                    MatterMost
	MattermostClient          *model.Client4
	MattermostWebSocketClient *model.WebSocketClient
	MattermostUser            *model.User
	MattermostChannel         *model.Channel
	mattermostTeam            *model.Team
}

func NewBot(cfg MatterMost) (*Bot, error) {
	bot := &Bot{
		Config: cfg,
	}

	bot.MattermostClient = model.NewAPIv4Client(bot.Config.mattermostServer)

	// Login.
	bot.MattermostClient.SetToken(bot.Config.MattermostToken)

	if user, _, err := bot.MattermostClient.GetMe(""); err != nil {
		return &Bot{}, fmt.Errorf("failed to get user from mattermost: %w", err)
	} else {
		bot.MattermostUser = user
	}

	// Find and save the bot's team to app struct.
	if team, _, err := bot.MattermostClient.GetTeamByName(bot.Config.mattermostTeamName, ""); err != nil {
		return &Bot{}, fmt.Errorf("failed to get team from mattermost: %w", err)
	} else {
		bot.mattermostTeam = team
	}

	// Find and save the talking channel to app struct.
	if channel, _, err := bot.MattermostClient.GetChannelByName(bot.Config.mattermostChannel, bot.mattermostTeam.Id, ""); err != nil {
		return &Bot{}, fmt.Errorf("failed to get channel from mattermost: %w", err)
	} else {
		bot.MattermostChannel = channel
	}

	return bot, nil
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
