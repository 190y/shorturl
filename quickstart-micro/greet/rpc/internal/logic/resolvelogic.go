package logic

import (
	"context"
	"shorturl/quickstart-micro/greet/rpc/internal/svc"
	"shorturl/quickstart-micro/greet/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResolveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResolveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResolveLogic {
	return &ResolveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResolveLogic) Resolve(in *pb.ResolveReq) (*pb.ResolveResp, error) {
	row, err := l.svcCtx.ShortUrlsModel.FindOneByCode(l.ctx, in.Code)
	if err != nil {
		return nil, err
	}

	return &pb.ResolveResp{
		LongUrl: row.LongUrl,
	}, nil
}
