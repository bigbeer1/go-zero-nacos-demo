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

type SysUserUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpLogic {
	return &SysUserUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserUpLogic) SysUserUp(req *types.SysUserUpRequest) (resp *types.Response, err error) {

	_, err = l.svcCtx.UserRpc.SysUserUpdate(l.ctx, &userclient.SysUserUpdateReq{
		Id:          req.Id,       // 用户ID
		Account:     req.Account,  // 用户名
		NickName:    req.NickName, // 姓名
		Password:    req.Password, // 密码
		State:       req.State,    // 状态 1:正常 2:停用 3:封禁
		UpdatedName: "超级管理员",      // 更新人
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
