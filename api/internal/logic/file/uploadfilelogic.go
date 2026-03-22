// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package file

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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

	// Return the response
	return &types.UploadResp{
		Filename: header.Filename,
		Size:     size,
	}, nil
}
