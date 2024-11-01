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

type SysUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserListLogic {
	return &SysUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserListLogic) SysUserList(req *types.SysUserListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.UserRpc.SysUserList(l.ctx, &userclient.SysUserListReq{
		Current:  req.Current,  // 页码
		PageSize: req.PageSize, // 页数
		Account:  req.Account,  // 用户名
		NickName: req.NickName, // 姓名
		Password: req.Password, // 密码
		State:    req.State,    // 状态 1:正常 2:停用 3:封禁
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysUserListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysUserListResp struct {
	Total int64              `json:"total"`
	List  []*SysUserDataList `json:"list"`
}

type SysUserDataList struct {
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
