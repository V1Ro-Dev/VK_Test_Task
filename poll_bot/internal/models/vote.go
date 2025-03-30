package models

import (
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
)

type Poll struct {
	ID            string
	Question      string
	AnswerOptions []string
	MemberVotes   map[string]string
	ChannelID     string
	Author        string
	IsActive      bool
}

func CreatePoll(post *model.Post) Poll {
	commands := strings.Fields(post.Message)
	return Poll{
		ID:            model.NewId(),
		Question:      commands[2],
		AnswerOptions: commands[3:],
		MemberVotes:   make(map[string]string),
		ChannelID:     post.ChannelId,
		Author:        post.UserId,
		IsActive:      true,
	}
}
