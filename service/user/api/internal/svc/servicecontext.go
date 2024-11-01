package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"go-zero-demo/service/user/api/internal/config"
	"go-zero-demo/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	// 日志服务
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
