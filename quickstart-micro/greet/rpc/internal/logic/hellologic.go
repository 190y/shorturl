package logic

import (
	"context"

	"shorturl/quickstart-micro/greet/rpc/internal/svc"
	"shorturl/quickstart-micro/greet/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HelloLogic) Hello(in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{
		Name: "hello " + in.Name,
	}, nil
}
