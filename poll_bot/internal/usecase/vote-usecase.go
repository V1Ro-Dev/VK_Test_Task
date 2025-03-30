package usecase

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"

	"poll_bot/internal/models"
	"poll_bot/internal/utils"
)

type VoteRepository interface {
	GetPollResults(poll models.Poll)
	EndPoll(poll models.Poll)
	DelPoll(poll models.Poll)
	CreatePoll(poll models.Poll) error
	Vote(poll models.Poll)
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
		return "", fmt.Errorf("create poll failed: %v", err)
	}

	return fmt.Sprintf("Successfully created poll. Question: %s. AnswerOptions: %v", poll.Question, poll.AnswerOptions), nil
}
func (v *VoteService) Vote(post *model.Post) (string, error) {
	if err := utils.ValidateCreateNVote(post); err != nil {
		return "", err
	}
	return "", nil
}
func (v *VoteService) CheckResults(post *model.Post) (string, error) {
	return "", nil
}
func (v *VoteService) End(post *model.Post) (string, error) {

	return "", nil
}
func (v *VoteService) Del(post *model.Post) (string, error) {
	return "", nil
}
