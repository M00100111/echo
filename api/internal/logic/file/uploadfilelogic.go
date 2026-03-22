// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package file

import (
	"Echo/api/internal/ai/indexer"
	models "Echo/api/internal/ai/models/embeddingModel"
	"Echo/api/internal/ai/transformer"
	"Echo/api/internal/config"
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"Echo/api/internal/svc"
	"Echo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile() (resp *types.UploadResp, err error) {
	// Get the HTTP request from the context
	r := l.ctx.Value("request").(*http.Request)

	// Parse multipart form
	err = r.ParseMultipartForm(32 << 20) // 32MB max
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	// Get the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}
	defer file.Close()

	// Create docs directory if it doesn't exist
	docsDir := "../docs"
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create docs directory: %w", err)
	}

	// Create the file path
	filePath := filepath.Join(docsDir, header.Filename)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy the file content
	size, err := io.Copy(dst, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	// 异步分析文件内容
	go Analyze(l.ctx, l.svcCtx.Config, header.Filename)

	// Return the response
	return &types.UploadResp{
		Filename: header.Filename,
		Size:     size,
	}, nil
}

func Analyze(ctx context.Context, config config.Config, fileName string) {
	trans := transformer.NewTransformer(ctx)
	bs, err := os.ReadFile("../docs/" + fileName)
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{{
		ID:      "fileName",
		Content: string(bs),
	},
	}
	embedder := models.NewEmbeddingModel(ctx)
	ind := indexer.NewIndexer(ctx, config, embedder)
	results, err := trans.Transform(ctx, docs)
	for index, doc := range results {
		doc.ID += strconv.Itoa(index)
	}
	ids, err := ind.Store(ctx, results)
	if err != nil {
		panic(err)
	}
	fmt.Println("分析完毕：", ids)
}
