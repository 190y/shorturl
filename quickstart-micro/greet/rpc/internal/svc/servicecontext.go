package svc

import (
	"shorturl/quickstart-micro/greet/rpc/internal/config"
	"shorturl/quickstart-micro/greet/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ShortUrlsModel model.ShortUrlsModel
	Redis          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	r := redis.MustNewRedis(c.CacheRedis)

	return &ServiceContext{
		Config:         c,
		ShortUrlsModel: model.NewShortUrlsModel(conn),
		Redis:          r,
	}
}
