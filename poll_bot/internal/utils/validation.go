package utils

import (
	"errors"
	"poll_bot/internal/models"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
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

func ValidateEndNDelNRes(post *model.Post) error {
	commands := strings.Fields(post.Message)
	if len(strings.Fields(post.Message)) < 3 {
		switch commands[1] {

		case "del":
			return errors.New(Usage("del"))

		case "end":
			return errors.New(Usage("end"))

		case "check_results":
			return errors.New(Usage("check_results"))
		}
	}

	return nil
}

func ValidateVote(poll models.Poll, post *model.Post) error {

	if !poll.IsActive {
		return errors.New("this poll is no longer active")
	}

	if _, exists := poll.MemberVotes[post.UserId]; exists {
		return errors.New("you can't vote twice")
	}

	ans := strings.Fields(post.Message)[3]
	for _, str := range poll.AnswerOptions {
		if str == ans {
			return nil
		}
	}

	return errors.New("this poll has no such option")
}

func ValidateEnd(poll models.Poll, post *model.Post) error {
	if poll.Author != post.UserId {
		return errors.New("only poll author can end this poll")
	}

	if !poll.IsActive {
		return errors.New("this poll is already closed")
	}

	return nil
}

func ValidateDel(poll models.Poll, post *model.Post) error {
	if poll.Author != post.UserId {
		return errors.New("only poll author can delete this poll")
	}

	return nil
}
