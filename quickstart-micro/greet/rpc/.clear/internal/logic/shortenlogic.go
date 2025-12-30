package logic

import (
	"context"

	"shorturl/quickstart-micro/greet/rpc/.clear/internal/svc"
	"shorturl/quickstart-micro/greet/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *pb.ShortenReq) (*pb.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ShortenResp{}, nil
}
