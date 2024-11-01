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

type SysUserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserAddLogic {
	return &SysUserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserAddLogic) SysUserAdd(req *types.SysUserAddRequest) (resp *types.Response, err error) {

	_, err = l.svcCtx.UserRpc.SysUserAdd(l.ctx, &userclient.SysUserAddReq{
		Account:     req.Account,  // 用户名
		NickName:    req.NickName, // 姓名
		Password:    req.Password, // 密码
		State:       req.State,    // 状态 1:正常 2:停用 3:封禁
		CreatedName: "超级管理员",      // 创建人
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
