package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"momo/app/bff/api/internal/svc"
	"momo/app/bff/api/internal/types"
	"momo/app/user/rpc/user"
	"momo/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)


type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	
	// 统计，验证码当天获取次数
	count, err := l.getVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("getVerificationCount mobile: %s, error: %v", req.Mobile, err)
	}
	if count > types.VerificationLimitPerDay {
		return nil, err 
	}

	// 查询 & 生成验证码
	code, err := getActivationCache(req.Mobile, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("getActivationCache  mobile: %s, error: %v", req.Mobile, err)
	}
	if len(code) == 0 {
		// 好几种情况
		// 1). 30 min，过期清除，需要重新生成
		// 2). 还没存过，需要生成
		code = util.RandomNumberic(6)
	}

	// 调用 UserRPC，给用户发送短信
	l.svcCtx.UserRPC.SendSms(l.ctx, &user.SendSmsRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		l.Logger.Errorf("sendSms mobile: %s, error: %v", req.Mobile, err)
		return nil, err
	}

	// 保存
	err = saveActivationCache(req.Mobile, code, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("saveActivationCache mobile: %s, error: %v", req.Mobile, err)
		return nil, err
	}

	// 累计 +1
	err = l.incrVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("incrVerificationCount mobile: %s, error: %v", req.Mobile, err)
	}

	return &types.VerificationResponse{}, nil
}


// 验证码
func (l *VerificationLogic) getVerificationCount(mobile string) (int, error) {
	key := fmt.Sprintf(types.PrefixVerificationCount, mobile)
	val, err := l.svcCtx.BizRedis.Get(key)
	if err != nil {
		return 0, err
	}
	if len(val) == 0 {
		return 0, nil
	}

	return strconv.Atoi(val)
}

func (l *VerificationLogic) incrVerificationCount(mobile string) error {
	key := fmt.Sprintf(types.PrefixVerificationCount, mobile)
	_, err := l.svcCtx.BizRedis.Incr(key)
	if err != nil {
		return err
	}

	return l.svcCtx.BizRedis.Expireat(key, util.EndOfDay(time.Now()).Unix())
}

// 激活码
func getActivationCache(mobile string, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(types.PrefixActivation, mobile)
	return rds.Get(key)
} 

func saveActivationCache(mobile, code string, rds *redis.Redis) error {
	key := fmt.Sprintf(types.PrefixActivation, mobile)
	return rds.Setex(key, code, types.ExpireActivation)  // 30 min，不过期
}

func delActivationCache(mobile, code string, rds *redis.Redis) error {
	key := fmt.Sprintf(types.PrefixActivation, mobile)
	_, err := rds.Del(key)
	return err
}

  