package logic

import (
	"context"
	"shorturl/quickstart-micro/greet/rpc/internal/svc"
	"shorturl/quickstart-micro/greet/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ResolveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	cachePrefix = "shorturl:"
	cacheTTL    = 24 * time.Hour
)

func NewResolveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResolveLogic {
	return &ResolveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResolveLogic) Resolve(in *pb.ResolveReq) (*pb.ResolveResp, error) {
	key := cachePrefix + in.Code

	v, err := l.svcCtx.Redis.Get(key)
	if err == nil && v != "" {
		l.Logger.Info("[cache-hit] code=%s", in.Code)
		return &pb.ResolveResp{
			LongUrl: v,
		}, nil
	}

	if err != nil && err != redis.Nil {
		l.Logger.Errorf("[cache-error] code=%s err=%v", in.Code, err)
	} else {
		l.Logger.Infof("[cache-miss] code=%s", in.Code)
	}

	row, err := l.svcCtx.ShortUrlsModel.FindOneByCode(l.ctx, in.Code)
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.Redis.Setex(key, row.LongUrl, int(cacheTTL.Seconds())); err != nil {
		l.Logger.Errorf("[cache-set-error] code=%s err=%v", in.Code, err)
	} else {
		l.Logger.Infof("[cache-set] code=%s ttl=%ds", in.Code, int(cacheTTL.Seconds()))
	}

	return &pb.ResolveResp{
		LongUrl: row.LongUrl,
	}, nil
}
