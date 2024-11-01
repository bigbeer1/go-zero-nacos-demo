package sysUser

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-demo/common"
	"go-zero-demo/common/msg"
	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"
	"go-zero-demo/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserInfoLogic {
	return &SysUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserInfoLogic) SysUserInfo(req *types.SysUserInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.UserRpc.SysUserFindOne(l.ctx, &userclient.SysUserFindOneReq{
		Id: req.Id, // 用户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysUserFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysUserFindOneResp struct {
	Id          string `json:"id"`           // 用户ID,
	Account     string `json:"account"`      // 用户名,
	NickName    string `json:"nick_name"`    // 姓名,
	Password    string `json:"password"`     // 密码,
	State       int64  `json:"state"`        // 状态 1:正常 2:停用 3:封禁,
	CreatedName string `json:"created_name"` // 创建人,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedName string `json:"updated_name"` // 更新人,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}
