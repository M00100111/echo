package transformer

import (
	"context"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
)

func NewTransformer(ctx context.Context) document.Transformer {
	transformer, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: true,
	})
	if err != nil {
		panic(err)
	}
	return transformer
}
