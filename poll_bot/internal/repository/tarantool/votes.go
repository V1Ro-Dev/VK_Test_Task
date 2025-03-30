package tarantool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tarantool/go-tarantool/v2"

	"poll_bot/config"
	"poll_bot/internal/models"
	"poll_bot/pkg/logger"
)

type TarantoolRepository struct {
	tarantoolConn *tarantool.Connection
}

func NewTarantoolRepository() *TarantoolRepository {
	tarantoolCfg := config.NewTarantoolConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logger.Info("connecting to tarantool...")

	dialer := tarantool.NetDialer{
		Address:  tarantoolCfg.GetURL(),
		User:     tarantoolCfg.GetUser(),
		Password: tarantoolCfg.GetPass(),
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		log.Fatal("tarantool connection error:", err)
	}

	logger.Info("successfully connected to tarantool")

	return &TarantoolRepository{tarantoolConn: conn}
}

func (t *TarantoolRepository) CreatePoll(poll models.Poll) error {
	jsonPoll, err := json.Marshal(poll)
	if err != nil {
		logger.Error("unable to parse struct: ", err.Error())
		return errors.New("wrong poll structure")
	}

	req := tarantool.NewInsertRequest("polls").
		Tuple([]interface{}{poll.ID, poll.ChannelID, string(jsonPoll)})

	_, err = t.tarantoolConn.Do(req).Get()
	if err != nil {
		logger.Error("unable to insert poll: ", err.Error())
		return fmt.Errorf("unable to create poll: %v", err)
	}

	return nil
}

func (t *TarantoolRepository) UpdatePoll(poll models.Poll) error {
	jsonPoll, err := json.Marshal(poll)
	if err != nil {
		logger.Error("unable to parse struct: ", err.Error())
		return errors.New("wrong poll structure")
	}

	req := tarantool.NewUpdateRequest("polls").
		Index("primary").
		Key([]interface{}{poll.ID, poll.ChannelID}).
		Operations(tarantool.NewOperations().
			Assign(2, string(jsonPoll)))

	_, err = t.tarantoolConn.Do(req).Get()
	if err != nil {
		logger.Error("unable to update poll: ", err.Error())
		return errors.New("somthing went wrong, try again")
	}

	return nil
}

func (t *TarantoolRepository) DelPoll(poll models.Poll) error {
	req := tarantool.NewDeleteRequest("polls").
		Index("primary").
		Key([]interface{}{poll.ID, poll.ChannelID})

	_, err := t.tarantoolConn.Do(req).Get()
	if err != nil {
		logger.Error("unable to delete poll: ", err.Error())
		return errors.New("unable to delete poll, try again")
	}

	return nil
}

func (t *TarantoolRepository) GetPoll(channelId string, pollId string) (models.Poll, error) {
	req := tarantool.NewSelectRequest("polls").
		Index("primary").
		Iterator(tarantool.IterEq).
		Key([]interface{}{pollId, channelId})

	resp, err := t.tarantoolConn.Do(req).Get()

	if err != nil {
		logger.Error("unable to get poll from db: ", err.Error())
		return models.Poll{}, errors.New("unable to get poll, please try again")
	}

	if len(resp) == 0 {
		return models.Poll{}, errors.New("poll not found")
	}

	resp = resp[0].([]interface{})
	var poll models.Poll

	err = json.Unmarshal([]byte(resp[2].(string)), &poll)
	if err != nil {
		logger.Error("unmarshal error: ", err.Error())
		return models.Poll{}, errors.New("unable to get poll, please try again")
	}

	return poll, nil
}

func (t *TarantoolRepository) Close() {
	t.tarantoolConn.Close()
}
