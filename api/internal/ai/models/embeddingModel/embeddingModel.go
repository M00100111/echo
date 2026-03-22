package models

import (
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"os"
)

func NewEmbeddingModel(ctx context.Context) *ark.Embedder {
	apiType := ark.APITypeMultiModal
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDING_MODEL_ID"),
		APIType: &apiType,
	})
	if err != nil {
		panic(err)
	}
	return embedder
}
