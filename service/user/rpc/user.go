package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"go-zero-demo/service/user/rpc/internal/config"
	"go-zero-demo/service/user/rpc/internal/server"
	"go-zero-demo/service/user/rpc/internal/svc"
	"go-zero-demo/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userclient.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// register service to nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.NacosService.Host, uint64(c.NacosService.Port)),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         c.NacosService.NamespaceId,
		TimeoutMs:           uint64(c.NacosService.Timeout),
		NotLoadCacheAtStart: c.NacosService.NotLoadCacheAtStart,
		LogDir:              c.NacosService.LogDir,
		CacheDir:            c.NacosService.CacheDir,
		LogLevel:            c.NacosService.LogLevel,
	}

	opts := nacos.NewNacosConfig(c.Name, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
