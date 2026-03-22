package retriever

import (
	"Echo/api/internal/config"
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus2"
	"github.com/cloudwego/eino-ext/components/retriever/milvus2/search_mode"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func NewRetriever(ctx context.Context, config config.Config, embedder *ark.Embedder) *milvus2.Retriever {
	retriever, err := milvus2.NewRetriever(ctx, &milvus2.RetrieverConfig{
		ClientConfig: &milvusclient.ClientConfig{
			Address: config.MilvusAddr,
			DBName:  config.MilvusDBName,
		},
		Collection: config.MilvusCollectionName,
		TopK:       2,
		SearchMode: search_mode.NewApproximate(milvus2.COSINE),
		Embedding:  embedder,
	})
	if err != nil {
		panic(err)
	}
	return retriever
}
