package handlers

import (
	"encoding/json"
	"errors"
	"github.com/mattermost/mattermost-server/v6/model"
	"poll_bot/config"
	"poll_bot/internal/utils"
	"strings"
)

type PollUseCase interface {
	Create(post *model.Post) (string, error)
	Vote(post *model.Post) (string, error)
	CheckResults(post *model.Post) (string, error)
	End(post *model.Post) (string, error)
	Del(post *model.Post) (string, error)
}

type PollHandler struct {
	pollUseCase PollUseCase
}

func NewPollHandler(pollUseCase PollUseCase) *PollHandler {
	return &PollHandler{
		pollUseCase: pollUseCase,
	}
}

func (p *PollHandler) StartListening(bot *config.Bot) {
	var err error
	for {
		bot.MattermostWebSocketClient, err = model.NewWebSocketClient4(bot.MattermostClient.URL, bot.Config.MattermostToken)
		if err != nil {
			return
		}

		bot.MattermostWebSocketClient.Listen()

		for event := range bot.MattermostWebSocketClient.EventChannel {
			// Launch new goroutine for handling the actual event.
			// If required, you can limit the number of events beng processed at a time.
			go p.HandleWebSocketEvents(bot, event)
		}
	}
}

func (p *PollHandler) HandleWebSocketEvents(bot *config.Bot, event *model.WebSocketEvent) {
	// Ignore other channels.
	if event.GetBroadcast().ChannelId != bot.MattermostChannel.Id {
		return
	}

	// Ignore other types of events.
	if event.EventType() != model.WebsocketEventPosted {
		return
	}

	// Since this event is a post, unmarshal it to (*model.Post)
	post := &model.Post{}
	err := json.Unmarshal([]byte(event.GetData()["post"].(string)), &post)
	if err != nil {
		return
	}

	// Ignore messages sent by this bot itself.
	if post.UserId == bot.MattermostUser.Id {
		return
	}

	if !utils.ValidateMessage(post.Message) {
		return
	}

	commands := strings.Fields(post.Message)
	res := ""
	var usecaseErr error

	switch commands[1] {
	case "create":
		res, usecaseErr = p.pollUseCase.Create(post)

	case "vote":
		res, usecaseErr = p.pollUseCase.Vote(post)

	case "check_results":
		res, usecaseErr = p.pollUseCase.CheckResults(post)

	case "end":
		res, usecaseErr = p.pollUseCase.End(post)

	case "del":
		res, usecaseErr = p.pollUseCase.Del(post)

	default:
		usecaseErr = errors.New("Unknown command " + commands[1])
	}

	if usecaseErr != nil {
		bot.SendMessage(post.ChannelId, err.Error())
		return
	}

	bot.SendMessage(post.ChannelId, res)
}
