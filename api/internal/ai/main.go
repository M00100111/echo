package main

func main() {
	//ctx := context.Background()
	//
	//// 加载 .env 文件
	//err := godotenv.Load("./ai/.env")
	//if err != nil {
	//	panic(err)
	//}
	//config := config.NewConfig()

	//embedder := models.NewEmbeddingModel(ctx, config)
	//vectors, err := embedder.EmbedStrings(ctx, []string{"hello", "how are you"})
	//if err != nil {
	//	log.Fatalf("EmbedStrings of Ark failed, err=%v", err)
	//}
	//log.Printf("vectors : %v", vectors)

	//bs, err := os.ReadFile("./docs/test.md")
	//if err != nil {
	//	panic(err)
	//}
	//docs := []*schema.Document{
	//	{
	//		ID:      "doc1",
	//		Content: string(bs),
	//	},
	//}
	//trans := transformer.NewTransformer(ctx)
	//splitDocs, err := trans.Transform(ctx, docs)
	//if err != nil {
	//	log.Fatalf("转换失败: %v", err)
	//}
	//for i, doc := range splitDocs {
	//	doc.ID = docs[0].ID + "_" + strconv.Itoa(i)
	//}

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
	//ids, err := indexer.Store(ctx, splitDocs)
	//if err != nil {
	//	log.Fatalf("Failed to store: %v", err)
	//	return
	//}
	//log.Printf("Store success, ids: %v", ids)

}
