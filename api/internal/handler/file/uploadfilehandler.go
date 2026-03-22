// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package file

import (
	"context"
	"net/http"

	"Echo/api/internal/logic/file"
	"Echo/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with the request object
		ctx := context.WithValue(r.Context(), "request", r)
		l := file.NewUploadFileLogic(ctx, svcCtx)
		resp, err := l.UploadFile()
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
		} else {
			httpx.OkJsonCtx(ctx, w, resp)
		}
	}
}
