package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/common"
	"go-zero-demo/common/msg"
	"net/http"
	"time"

	"go-zero-demo/service/user/api/internal/config"
	"go-zero-demo/service/user/api/internal/handler"
	"go-zero-demo/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	zero_handler "github.com/zeromicro/go-zero/rest/handler"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf,
		// token错误拦截
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.TokenErrorCode, msg.TokenError, err.Error()))
		}),
		// 请求方式错误拦截
		rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqNotAllCode, msg.ReqNotAllError, nil))
		})),
		// 路由错误拦截
		rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqRoutesErrorCode, msg.ReqRoutesError, nil))
		})),
	)

	defer server.Stop()

	// 设置日志输出 接口慢时间
	zrpc.SetClientSlowThreshold(time.Second * 2000)
	zero_handler.SetSlowThreshold(time.Second * 2000)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
