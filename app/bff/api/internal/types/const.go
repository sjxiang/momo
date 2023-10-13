package types

const (
	UserIdKey = "userId"
)

const (
	
	PrefixActivation        = "biz#activation#%s"          // 验证码，key
	ExpireActivation        = 60 * 30                      // 30 min 内，验证码不变 

	PrefixVerificationCount = "biz#Verification#count#%s"  // 验证码当天获取次数，key
	VerificationLimitPerDay = 10                           // 上限，1 天 10 次

)
