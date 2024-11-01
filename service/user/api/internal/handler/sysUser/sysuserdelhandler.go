package sysUser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/service/user/api/internal/logic/sysUser"
	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"
)

func SysUserDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sysUser.NewSysUserDelLogic(r.Context(), svcCtx)
		resp, err := l.SysUserDel(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
