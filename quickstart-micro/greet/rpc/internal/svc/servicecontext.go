package svc

import (
	"shorturl/quickstart-micro/greet/rpc/internal/config"
	"shorturl/quickstart-micro/greet/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ShortUrlsModel model.ShortUrlsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:         c,
		ShortUrlsModel: model.NewShortUrlsModel(conn),
	}
}
