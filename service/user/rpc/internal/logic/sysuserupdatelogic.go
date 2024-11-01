package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"go-zero-demo/service/user/rpc/internal/svc"
	"go-zero-demo/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpdateLogic {
	return &SysUserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserUpdateLogic) SysUserUpdate(in *userclient.SysUserUpdateReq) (*userclient.CommonResp, error) {

	res, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysUser没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("SysUser该ID已被删除：" + in.Id)
	}

	// 用户名
	if len(in.Account) > 0 {
		res.Account = in.Account
	}
	// 姓名
	if len(in.NickName) > 0 {
		res.NickName = in.NickName
	}
	// 密码
	if len(in.Password) > 0 {
		res.Password = in.Password
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysUserModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &userclient.CommonResp{}, nil
}
