package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	bj38 "bj38/proto/bj38"
)

type Bj38 struct{}

func (e *Bj38) Handle(ctx context.Context, msg *bj38.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *bj38.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
