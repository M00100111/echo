package indexer

import (
	"Echo/api/internal/config"
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus2"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func NewIndexer(ctx context.Context, config config.Config, embedder *ark.Embedder) *milvus2.Indexer {
	indexer, err := milvus2.NewIndexer(ctx, &milvus2.IndexerConfig{
		ClientConfig: &milvusclient.ClientConfig{
			Address: config.MilvusAddr,
			DBName:  config.MilvusDBName,
		},
		Collection: config.MilvusCollectionName,

		Vector: &milvus2.VectorConfig{
			Dimension:    2048, // 与 embedding 模型维度匹配
			MetricType:   milvus2.COSINE,
			IndexBuilder: milvus2.NewHNSWIndexBuilder().WithM(16).WithEfConstruction(200),
		},
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	return indexer
}
