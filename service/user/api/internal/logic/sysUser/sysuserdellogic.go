package sysUser

import (
	"context"
	"go-zero-demo/common"
	"go-zero-demo/common/msg"
	"go-zero-demo/service/user/rpc/userclient"

	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserDelLogic {
	return &SysUserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserDelLogic) SysUserDel(req *types.SysUserDelRequest) (resp *types.Response, err error) {

	_, err = l.svcCtx.UserRpc.SysUserDelete(l.ctx, &userclient.SysUserDeleteReq{
		Id:          req.Id,  // 用户ID
		DeletedName: "超级管理员", // 删除人
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
