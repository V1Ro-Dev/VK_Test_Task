package tarantool

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tarantool/go-tarantool/v2"

	"poll_bot/config"
	"poll_bot/internal/models"
)

type TarantoolRepository struct {
	tarantoolConn *tarantool.Connection
}

func NewTarantoolRepository() *TarantoolRepository {
	tarantoolCfg := config.NewTarantoolConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dialer := tarantool.NetDialer{
		Address: tarantoolCfg.GetURL(),
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		log.Fatal("tarantool connection error:", err)
	}

	return &TarantoolRepository{tarantoolConn: conn}
}

func (t *TarantoolRepository) CreatePoll(poll models.Poll) error {
	jsonPoll, err := json.Marshal(poll)
	if err != nil {
		return fmt.Errorf("unable to parse struct: %v", err)
	}

	req := tarantool.NewInsertRequest("polls").Tuple([]interface{}{poll.ID, poll.ChannelID, string(jsonPoll)})

	_, err = t.tarantoolConn.Do(req).Get()
	if err != nil {
		return fmt.Errorf("unable to create poll: %v", err)
	}

	return nil
}

func (t *TarantoolRepository) UpdatePoll(poll models.Poll) error {
	jsonPoll, err := json.Marshal(poll)
	if err != nil {
		return err
	}

	req := tarantool.NewUpdateRequest("polls").Index("primary").Key([]interface{}{poll.ID, poll.ChannelID}).Operations(tarantool.NewOperations().Assign(2, string(jsonPoll)))
	_, err = t.tarantoolConn.Do(req).Get()
	if err != nil {
		return err
	}

	return nil
}

func (t *TarantoolRepository) DelPoll(poll models.Poll) error {
	req := tarantool.NewDeleteRequest("polls").Index("primary").Key([]interface{}{poll.ID, poll.ChannelID})

	_, err := t.tarantoolConn.Do(req).Get()
	if err != nil {
		fmt.Errorf("unable to delete poll: %v", err)
	}

	return nil
}

func (t *TarantoolRepository) GetPoll(channelId string, pollId string) (models.Poll, error) {
	req := tarantool.NewSelectRequest("polls").Index("primary").Iterator(tarantool.IterEq).Key([]interface{}{pollId, channelId})

	resp, err := t.tarantoolConn.Do(req).Get()
	if err != nil {
		return models.Poll{}, fmt.Errorf("get query error: %v", err)
	}

	if len(resp) == 0 {
		return models.Poll{}, fmt.Errorf("poll not found")
	}

	resp = resp[0].([]interface{})
	var poll models.Poll
	err = json.Unmarshal([]byte(resp[2].(string)), &poll)
	if err != nil {
		return models.Poll{}, fmt.Errorf("unmarshal error: %v", err)
	}

	return poll, nil
}

func (t *TarantoolRepository) Close() {
	t.tarantoolConn.Close()
}
