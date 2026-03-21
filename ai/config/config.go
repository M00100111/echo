package config

import "os"

type Config struct {
	ChatModel            string
	EmbeddingModel       string
	APIKey               string
	MilvusAddr           string
	MilvusDBName         string
	MilvusCollectionName string
}

func NewConfig() Config {
	return Config{
		ChatModel:            os.Getenv("CHAT_MODEL_ID"),
		EmbeddingModel:       os.Getenv("EMBEDDING_MODEL_ID"),
		APIKey:               os.Getenv("ARK_API_KEY"),
		MilvusAddr:           os.Getenv("MILVUS_ADDR"),
		MilvusDBName:         os.Getenv("MILVUS_DB_NAME"),
		MilvusCollectionName: os.Getenv("MILVUS_COLLECTION_NAME"),
	}
}
