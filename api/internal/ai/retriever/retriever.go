package retriever

import (
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus2"
	"github.com/cloudwego/eino-ext/components/retriever/milvus2/search_mode"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"os"
)

func NewRetriever(ctx context.Context, embedder *ark.Embedder) *milvus2.Retriever {
	retriever, err := milvus2.NewRetriever(ctx, &milvus2.RetrieverConfig{
		ClientConfig: &milvusclient.ClientConfig{
			Address: os.Getenv("MILVUS_ADDR"),
			DBName:  os.Getenv("MILVUS_DB_NAME"),
		},
		Collection: os.Getenv("MILVUS_COLLECTION_NAME"),
		TopK:       2,
		SearchMode: search_mode.NewApproximate(milvus2.COSINE),
		Embedding:  embedder,
	})
	if err != nil {
		panic(err)
	}
	return retriever
}
