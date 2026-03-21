package models

import (
	"Echo/ai/config"
	"context"
	"github.com/cloudwego/eino-ext/components/model/agenticark"
)

func NewChatModel(ctx context.Context, config config.Config) *agenticark.Model {
	am, err := agenticark.New(ctx, &agenticark.Config{
		Model:  config.ChatModel,
		APIKey: config.APIKey,
	})
	if err != nil {
		panic(err)
	}
	return am
}
