package usecase

import (
	"fmt"
	"poll_bot/pkg/logger"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"

	"poll_bot/internal/models"
	"poll_bot/internal/utils"
)

type VoteRepository interface {
	UpdatePoll(poll models.Poll) error
	DelPoll(poll models.Poll) error
	CreatePoll(poll models.Poll) error
	GetPoll(channelId string, id string) (models.Poll, error)
}

type VoteService struct {
	voteRepo VoteRepository
}

func NewVoteService(voteRepo VoteRepository) *VoteService {
	return &VoteService{
		voteRepo: voteRepo,
	}
}

func (v *VoteService) Create(post *model.Post) (string, error) {
	if err := utils.ValidateCreateNVote(post); err != nil {
		return "", err
	}

	poll := models.CreatePoll(post)
	if err := v.voteRepo.CreatePoll(poll); err != nil {
		return "", fmt.Errorf("poll creation failed: %v", err)
	}

	res := fmt.Sprintf("Successfully created poll\nPollID: %s\nQuestion: %s\nAnswer Options:\n", poll.ID, poll.Question)
	for _, opt := range poll.AnswerOptions {
		res += fmt.Sprintf("\t%s\n", opt)
	}

	logger.Info(fmt.Sprintf("Created poll with ID: %s", poll.ID))
	return res, nil
}
func (v *VoteService) Vote(post *model.Post) (string, error) {
	if err := utils.ValidateCreateNVote(post); err != nil {
		return "", err
	}

	commands := strings.Fields(post.Message)
	poll, err := v.voteRepo.GetPoll(post.ChannelId, commands[2])
	if err != nil {
		return "", fmt.Errorf("vote failed: %v", err)
	}

	if err = utils.ValidateVote(poll, post); err != nil {
		return "", err
	}

	poll.MemberVotes[post.UserId] = commands[3]
	if err = v.voteRepo.UpdatePoll(poll); err != nil {
		return "", fmt.Errorf("vote failed: %v", err)
	}

	logger.Info(fmt.Sprintf("Vote for Poll with ID: %s and Option: %s was successfuly executed", poll.ID, poll.MemberVotes[post.UserId]))
	return fmt.Sprintf("You have successfully voted for Poll: %s with Option: %s", commands[2], commands[3]), nil
}
func (v *VoteService) CheckResults(post *model.Post) (string, error) {
	if err := utils.ValidateEndNDelNRes(post); err != nil {
		return "", err
	}

	poll, err := v.voteRepo.GetPoll(post.ChannelId, strings.Fields(post.Message)[2])
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf("Vote results \"%s\":\n", poll.Question)
	for _, answer := range poll.MemberVotes {
		cnt := 0
		for _, option := range poll.AnswerOptions {
			if answer == option {
				cnt++
			}
		}
		res += fmt.Sprintf(" - %s: %d votes\n", answer, cnt)
	}

	logger.Info("check_results command was successfully executed")
	return res, nil
}

func (v *VoteService) End(post *model.Post) (string, error) {
	if err := utils.ValidateEndNDelNRes(post); err != nil {
		return "", err
	}

	poll, err := v.voteRepo.GetPoll(post.ChannelId, strings.Fields(post.Message)[2])
	if err != nil {
		return "", err
	}

	if err = utils.ValidateEnd(poll, post); err != nil {
		return "", err
	}

	poll.IsActive = false
	if err = v.voteRepo.UpdatePoll(poll); err != nil {
		return "", fmt.Errorf("end poll failed: %v", err)
	}

	logger.Info(fmt.Sprintf("Poll with ID was ended: %s", poll.ID))
	return fmt.Sprintf("You have successfully end poll with PollID: %s", poll.ID), nil
}

func (v *VoteService) Del(post *model.Post) (string, error) {
	if err := utils.ValidateEndNDelNRes(post); err != nil {
		return "", err
	}

	poll, err := v.voteRepo.GetPoll(post.ChannelId, strings.Fields(post.Message)[2])
	if err != nil {
		return "", err
	}

	if err = utils.ValidateDel(poll, post); err != nil {
		return "", err
	}

	if err = v.voteRepo.DelPoll(poll); err != nil {
		return "", err
	}

	logger.Info(fmt.Sprintf("Poll with ID: %s was deleted", poll.ID))
	return fmt.Sprintf("You have successfully deleted poll with PollID: %s", poll.ID), nil
}
