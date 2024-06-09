package video

import (
	"context"
	"fmt"

	"video-tools/internal/load"
	"video-tools/internal/svc"
	"video-tools/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExtractVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExtractVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExtractVideoLogic {
	return &ExtractVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExtractVideoLogic) ExtractVideo(req *types.ParseReq) (resp *types.ParseResult, err error) {
	fmt.Println(*req)
	p, err := load.LoadParser(l.svcCtx, req.URL)
	if err != nil {
		return
	}
	resp, err = p.Parse()
	return
}
