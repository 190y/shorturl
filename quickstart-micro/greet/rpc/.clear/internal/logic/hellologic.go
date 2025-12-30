package logic

import (
	"context"

	"shorturl/quickstart-micro/greet/rpc/.clear/internal/svc"
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
	// todo: add your logic here and delete this line

	return &pb.HelloResp{}, nil
}
