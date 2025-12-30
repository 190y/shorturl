// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"shorturl/quickstart-micro/greet/rpc/pb"

	"shorturl/quickstart-micro/greet/api/internal/svc"
	"shorturl/quickstart-micro/greet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := l.svcCtx.GreetRpc.Shorten(l.ctx, &pb.ShortenReq{
		LongUrl: req.LongUrl,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{
		Code: resp.Code,
	}, nil
}
