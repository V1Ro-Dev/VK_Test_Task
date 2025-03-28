package internal

import (
	"poll_bot/config"
	"poll_bot/internal/delivery/handlers"
	"poll_bot/internal/repository/tarantool"
	"poll_bot/internal/usecase"
)

func Run() error {

	newVoteRepo := tarantool.NewTarantoolRepository()
	newVoteService := usecase.NewVoteService(newVoteRepo)
	newPollHandler := handlers.NewPollHandler(newVoteService)
	defer newVoteRepo.Close()

	matterMostCfg, err := config.LoadConfig("")
	if err != nil {
		return err
	}

	bot, err := config.NewBot(*matterMostCfg)
	if err != nil {
		return err
	}

	newPollHandler.StartListening(bot)

	return nil
}
