package internal

import (
	"poll_bot/config"
	"poll_bot/internal/delivery/handlers"
	"poll_bot/internal/repository/tarantool"
	"poll_bot/internal/usecase"
	"poll_bot/pkg/logger"
)

func Run() error {

	newVoteRepo := tarantool.NewTarantoolRepository()
	newVoteService := usecase.NewVoteService(newVoteRepo)
	newPollHandler := handlers.NewPollHandler(newVoteService)
	defer newVoteRepo.Close()

	matterMostCfg, err := config.LoadConfig("")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	bot, err := config.NewBot(*matterMostCfg)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	newPollHandler.StartListening(bot)
	logger.Info("poll bot is running")

	return nil
}
