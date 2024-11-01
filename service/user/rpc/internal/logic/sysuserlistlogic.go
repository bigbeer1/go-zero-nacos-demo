package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"go-zero-demo/service/user/rpc/internal/svc"
	"go-zero-demo/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserListLogic {
	return &SysUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserListLogic) SysUserList(in *userclient.SysUserListReq) (*userclient.SysUserListResp, error) {
	whereBuilder := l.svcCtx.SysUserModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 用户名
	if len(in.Account) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"account ": "%" + in.Account + "%",
		})
	}
	// 姓名
	if len(in.NickName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}
	// 密码
	if len(in.Password) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"password ": "%" + in.Password + "%",
		})
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.SysUserModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysUserModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 用户名
	if len(in.Account) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"account ": "%" + in.Account + "%",
		})
	}
	// 姓名
	if len(in.NickName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}
	// 密码
	if len(in.Password) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"password ": "%" + in.Password + "%",
		})
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	count, err := l.svcCtx.SysUserModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*userclient.SysUserListData
	for _, item := range all {
		list = append(list, &userclient.SysUserListData{
			Id:          item.Id,                         //用户ID
			Account:     item.Account,                    //用户名
			NickName:    item.NickName,                   //姓名
			Password:    item.Password,                   //密码
			State:       item.State,                      //状态 1:正常 2:停用 3:封禁
			CreatedName: item.CreatedName,                //创建人
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedName: item.UpdatedName.String,         //更新人
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
		})
	}

	return &userclient.SysUserListResp{
		Total: count,
		List:  list,
	}, nil
}
