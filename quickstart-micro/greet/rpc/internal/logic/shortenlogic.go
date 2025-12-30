package logic

import (
	"context"
	"math/rand"
	"shorturl/quickstart-micro/greet/rpc/internal/model"
	"time"

	"shorturl/quickstart-micro/greet/rpc/internal/svc"
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
	// 1. 生成一个简单的 code（先够用）
	rand.Seed(time.Now().UnixNano())
	code := randomString(6)

	// 2. 写入数据库
	_, err := l.svcCtx.ShortUrlsModel.Insert(l.ctx, &model.ShortUrls{
		Code:    code,
		LongUrl: in.LongUrl,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ShortenResp{
		Code: code,
	}, nil
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
