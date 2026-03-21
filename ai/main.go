package main

import (
	"Echo/ai/transformer"
	"context"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	// 加载 .env 文件
	err := godotenv.Load("./ai/.env")
	if err != nil {
		panic(err)
	}
	//config := config.NewConfig()

	//embedder := models.NewEmbeddingModel(ctx, config)
	//vectors, err := embedder.EmbedStrings(ctx, []string{"hello", "how are you"})
	//if err != nil {
	//	log.Fatalf("EmbedStrings of Ark failed, err=%v", err)
	//}
	//log.Printf("vectors : %v", vectors)

	// 创建索引器
	//indexer := indexer.NewIndexer(ctx, config, embedder)
	//// 存储文档
	//docs := []*schema.Document{
	//	{
	//		ID:      "1",
	//		Content: "你说得对",
	//		MetaData: map[string]any{
	//			"category": "md",
	//			"year":     2021,
	//		},
	//	},
	//	{
	//		ID:      "2",
	//		Content: "原神启动",
	//	},
	//}
	//ids, err := indexer.Store(ctx, docs)
	//if err != nil {
	//	log.Fatalf("Failed to store: %v", err)
	//	return
	//}
	//log.Printf("Store success, ids: %v", ids)

	// 检索文档
	//retriever := retriever.NewRetriever(ctx, config, embedder)
	//documents, err := retriever.Retrieve(ctx, "原神")
	//if err != nil {
	//	log.Fatalf("Failed to retrieve: %v", err)
	//	return
	//}
	//
	//// 打印文档
	//for i, doc := range documents {
	//	fmt.Printf("Document %d:\n", i)
	//	fmt.Printf("  ID: %s\n", doc.ID)
	//	fmt.Printf("  Content: %s\n", doc.Content)
	//	fmt.Printf("  Score: %v\n", doc.Score())
	//}

	//content, err := os.OpenFile("./docs/test.md", os.O_CREATE|os.O_RDONLY, 0755)
	//if err != nil {
	//	panic(err)
	//}
	//defer content.Close()
	bs, err := os.ReadFile("./docs/test.md")
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{
			ID:      "1",
			Content: string(bs),
		},
	}
	trans := transformer.NewTransformer(ctx)
	splitDocs, err := trans.Transform(ctx, docs)
	if err != nil {
		log.Fatalf("转换失败: %v", err)
	}
	for _, doc := range splitDocs {
		log.Printf("内容: %s, 元数据: %v\n", doc.Content, doc.MetaData)
	}
}
