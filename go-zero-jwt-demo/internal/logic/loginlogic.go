package logic

import (
	"context"
	"jwt-demo/third_party"
	"time"

	"jwt-demo/internal/svc"
	"jwt-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	// todo: 登录并查询数据库 这里示范文件 直接跳过
	// todo: 返回token
	nowTime := time.Now()
	ExpiresAt := nowTime.Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))
	token, err := third_party.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, nowTime.Unix(), ExpiresAt.Unix(), "test_123123131", "test", "1111")
	if err != nil {
		logx.Error("token gen error", err)
		return nil, err
	}
	resp.Token = token
	return resp, nil
}
