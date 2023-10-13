


# 开发日志

1. 项目介绍

微服务
高并发
深入源码
上线部署


架构设计，从单体到微服务，如何拆分
面对高并发，针对性的性能优化和架构调整



- zero-examples 所有功能全部在项目中体现
- 服务注册与发现，负载均衡
- 监控 && 告警 && 链路追踪
- 缓存异常的处理，穿透、击穿、雪崩、缓存数据库一致性保证
- 秒杀场景，超高并发的读写（模拟 10w QPS，压测）
- 请求耗时高如何优化
- 聊天室
- 问题排查定位方法
- CI/CD 完整部署上线，并可访问



准备工作
- 搭建 Go 开发环境，version 1.21.0
- 安装 goctl
- 安装 protoc，protoc-gen-go
- 可访问的 Mysql，Redis，Kafka，ElasticSearch，Etcd，对版本不做要求
- IDE，VSCode
- 调试工具 Postman
- 浏览器 Chrome



什么产品

<!-- 

项目介绍 - 仿知乎


首页
    推荐、热榜

关注
    时间线、动态

+
    发布文章、想法

会员


我的
    个人中心

小红点
    消息通知、聊天 

    reply 回复
    comment 评论
    like 赞
    
    -->



# 微服务拆分

1. DDD（难以理解）
2. 垂直功能（按照所在的职能部门，独立业务功能进行划分）



===


# 业务开发流程

1. api 代码生成
2. rpc 代码生成
3. model 代码生成（sqlx、gorm、ent）
4. 启动 rpc 服务


服务注册与发现

    - 注册中心
    - 服务调用方
    - 服务提供方

    server 启动之前的准备工作，go-zero 默默承受了太多


    查看 etcd 中是否已注册
        $ etcdctl get --prefix user.rpc
        user.rpc/7587873906211629829
        192.168.56.102:8080


    查看租约
        $ etcdctl lease list
        found 1 leases
        694d8b1d8c332b05


    查看租约剩余时间
        $ etcdctl lease timetolive 694d8b1d8c332b05
        lease 694d8b1d8c332b05 granted with TTL(10s), remaining(9s)



短信验证

    查看短信验证码
        > keys * 
        1) "biz#activation#18851762282"
        2) "biz#Verification#count#18851762282"
        
        > get biz#activation#18851762282
        "736170"
    
    查看当日获取验证码次数
        > get biz#Verification#count#18851762282
        "1"
    


注册
    加密 vs. 编码
        加密，考虑数据敏感，不可逆，e.g. md5
        编码，考虑可读性，e.g. base64

    sqlx 自动生成
        db + cache
            db
                查询，要处理 sqlx.ErrNotFound
                插入，要处理 result.LastInsertId()
    
            cache
                默认，7 天过期
                应对内存穿透，设置空缓存，1 min

                > get cache:user:id:1
                "{\"Id\":1,\"Username\":\"sjxiang\",\"Avatar\":\"\",\"Mobile\":\"MTg4NTE3NjIyODI=\",\"CreateTime\":\"2023-10-13T12:47:27+08:00\",\"UpdateTime\":\"2023-10-13T12:47:27+08:00\"}"

    sqlx 自定义   
        略        

    用户输入 mobile，可以瞎几把填，需要查询验证，故不能忽略 sqlx.ErrNotFound；
    而 JWT 携带 userId 不容易作伪，可以忽略。


登录
    JWT 
        生成 
        验证
            github.com/zeromicro/go-zero/rest/handler
            func Authorize(secret string, opts ...AuthorizeOption) func(http.Handler) http.Handler  
        
        illa-cloud
        

```go
// Authorize returns an authorization middleware.
func Authorize(secret string, opts ...AuthorizeOption) func(http.Handler) http.Handler {
	var authOpts AuthorizeOptions
	for _, opt := range opts {
		opt(&authOpts)
	}

	parser := token.NewTokenParser()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            tok, err := parser.ParseToken(r, secret, authOpts.PrevSecret)
			if err != nil {
				unauthorized(w, r, err, authOpts.Callback)
				return
			}

			if !tok.Valid {
				unauthorized(w, r, errInvalidToken, authOpts.Callback)
				return
			}

			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				unauthorized(w, r, errNoClaims, authOpts.Callback)
				return
			}

			ctx := r.Context()
			for k, v := range claims {
				switch k {
				case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
					// ignore the standard claims
				default:
					ctx = context.WithValue(ctx, k, v)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

```



# 错误处理

业务错误码

    统一返回 200，简单粗暴
