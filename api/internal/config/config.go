// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ChatModel            string `yaml:"ARK_API_KEY"`
	EmbeddingModel       string `yaml:"CHAT_MODEL_ID"`
	APIKey               string `yaml:"EMBEDDING_MODEL_ID"`
	MilvusAddr           string `yaml:"MILVUS_ADDR"`
	MilvusDBName         string `yaml:"MILVUS_DB_NAME"`
	MilvusCollectionName string `yaml:"MILVUS_COLLECTION_NAME"`
}
