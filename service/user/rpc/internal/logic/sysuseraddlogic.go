package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go-zero-demo/service/user/model"
	"go-zero-demo/service/user/rpc/internal/svc"
	"go-zero-demo/service/user/rpc/userclient"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserAddLogic {
	return &SysUserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserAddLogic) SysUserAdd(in *userclient.SysUserAddReq) (*userclient.CommonResp, error) {
	_, err := l.svcCtx.SysUserModel.Insert(l.ctx, &model.SysUser{
		Id:          uuid.NewV4().String(), // ID
		CreatedAt:   time.Now(),            // 创建时间
		Account:     in.Account,            // 用户名
		NickName:    in.NickName,           // 姓名
		Password:    in.Password,           // 密码
		State:       in.State,              // 状态 1:正常 2:停用 3:封禁
		CreatedName: in.CreatedName,        // 创建人
	})
	if err != nil {
		return nil, err
	}

	return &userclient.CommonResp{}, nil
}
