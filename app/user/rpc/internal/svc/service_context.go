package svc

import (
	"momo/app/user/rpc/internal/config"
	"momo/app/user/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	// todo - 尝试换成 gorm

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	  
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
