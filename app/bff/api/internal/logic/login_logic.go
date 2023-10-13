package logic

import (
	"context"
	"errors"
	"strings"

	"momo/app/bff/api/internal/svc"
	"momo/app/bff/api/internal/types"
	"momo/app/user/rpc/user"
	"momo/pkg/cryptx"
	"momo/pkg/jwtx"

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

// 填充业务逻辑（登录有好多搭配方式，账号密码登录、短信验证登录）
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 请求参数处理（过滤空格、判断非空）
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, errors.New("login mobile cannot be empty")
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, errors.New("verification code cannot be empty")
	}
	
	// 请求参数校验（验证码）
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}

	mobile := cryptx.Encode(req.Mobile)
	// 调用 UserRPC 判断手机号是否已经注册
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u == nil || u.UserId == 0 {
		return nil, errors.New("mobile has no registered")  // 查无此人
	}

	// 生成 JWT
	token, err := jwtx.BuildTokens(jwtx.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	// 卸磨杀驴
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	return &types.LoginResponse{
		UserId: u.UserId, 
		Token:  types.Token{
			AccessToken:  token.AccessToken, 
			AccessExpire: token.AccessExpire,
		},
	}, nil
}



