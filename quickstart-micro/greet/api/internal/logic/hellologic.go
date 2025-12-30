// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"shorturl/quickstart-micro/greet/rpc/greet"

	"shorturl/quickstart-micro/greet/api/internal/svc"
	"shorturl/quickstart-micro/greet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloReq) (resp *types.HelloResp, err error) {
	r, err := l.svcCtx.GreetRpc.Hello(l.ctx, &greet.HelloReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.HelloResp{
		Name: r.Name,
	}, nil
}
