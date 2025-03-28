package utils

import (
	"errors"
	"github.com/mattermost/mattermost-server/v6/model"
	"strings"
)

func ValidateMessage(message string) bool {
	return strings.HasPrefix(message, "/poll") && len(strings.Fields(message)) >= 3
}

func ValidateCreateNVote(post *model.Post) error {
	commands := strings.Fields(post.Message)
	if len(strings.Fields(post.Message)) < 4 {
		if commands[1] == "create" {
			return errors.New(Usage("create"))
		}
		return errors.New(Usage("vote"))
	}

	return nil
}
