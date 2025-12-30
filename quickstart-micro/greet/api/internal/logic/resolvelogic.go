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

type ResolveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResolveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResolveLogic {
	return &ResolveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResolveLogic) Resolve(req *types.ResolveReq) (resp *types.ResolveResp, err error) {
	respo, err := l.svcCtx.GreetRpc.Resolve(l.ctx, &pb.ResolveReq{
		Code: req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &types.ResolveResp{
		LongUrl: respo.LongUrl,
	}, nil
}
