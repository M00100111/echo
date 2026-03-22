// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"os"
)

type Config struct {
	rest.RestConf
	ChatModel            string `yaml:"CHAT_MODEL_ID"`
	EmbeddingModel       string `yaml:"EMBEDDING_MODEL_ID"`
	APIKey               string `yaml:"ARK_API_KEY"`
	MilvusAddr           string `yaml:"MILVUS_ADDR"`
	MilvusDBName         string `yaml:"MILVUS_DB_NAME"`
	MilvusCollectionName string `yaml:"MILVUS_COLLECTION_NAME"`
}

func (c *Config) InitConfig() {
	c.ChatModel = os.Getenv("CHAT_MODEL_ID")
	c.EmbeddingModel = os.Getenv("EMBEDDING_MODEL_ID")
	c.APIKey = os.Getenv("ARK_API_KEY")
	c.MilvusAddr = os.Getenv("MILVUS_ADDR")
	c.MilvusDBName = os.Getenv("MILVUS_DB_NAME")
	c.MilvusCollectionName = os.Getenv("MILVUS_COLLECTION_NAME")
}
