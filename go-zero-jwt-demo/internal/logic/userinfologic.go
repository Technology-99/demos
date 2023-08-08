package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"

	"jwt-demo/internal/svc"
	"jwt-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	logc.Info(l.ctx, "userId: ", req.ID)

	// 获取 jwt 载体信息
	value := l.ctx.Value("test")
	logc.Info(l.ctx, logc.Field("test: ", value))

	resp = new(types.UserInfoResp)
	resp.Name = "huanghua"
	return resp, nil
}
