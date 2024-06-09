package video

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"video-tools/internal/logic/video"
	"video-tools/internal/svc"
	"video-tools/internal/types"
)

func ExtractVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ParseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewExtractVideoLogic(r.Context(), svcCtx)
		resp, err := l.ExtractVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
