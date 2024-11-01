package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/service/user/model"
	"go-zero-demo/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	SysUserModel model.SysUserModel // 用户
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库的连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		SysUserModel: model.NewSysUserModel(conn, c.CacheRedis),
	}
}
