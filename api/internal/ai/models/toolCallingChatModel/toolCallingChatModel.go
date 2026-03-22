package models

import (
	"Echo/api/internal/config"
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
)

func NewToolCallingChatModel(ctx context.Context, config config.Config) *ark.ChatModel {
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: config.APIKey,
		Model:  config.ChatModel,
	})
	if err != nil {
		panic(err)
	}
	return model
}
