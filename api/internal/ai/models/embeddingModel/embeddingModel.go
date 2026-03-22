package models

import (
	"Echo/api/internal/config"
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

func NewEmbeddingModel(ctx context.Context, config config.Config) *ark.Embedder {
	apiType := ark.APITypeMultiModal
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  config.APIKey,
		Model:   config.EmbeddingModel,
		APIType: &apiType,
	})
	if err != nil {
		panic(err)
	}
	return embedder
}
