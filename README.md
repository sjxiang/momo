

<p align="center">
  <a href="https://github.com/sjxiang/momo"><img src="./doc/logo.jpg" width="150" height="150" alt="momo logo"></a>
</p>


<div align="center">

# momo

</div>


> **Note**
> 仿知乎
> 


## 功能
1. 用户注册


## 部署
### 基于 Docker 进行部署


小程序 - HTTP 请求 -> BFF - RPC 调用 -> 微服务 


"@timestamp":"2023-10-13T21:28:55.017+08:00","caller":"internal/log.go:77","content":"(/v1/user/info - 127.0.0.1:33598) interface conversion: interface {} is nil, not json.Number\ngoroutine 78 [running]:

\nruntime/debug.Stack()\n\t/home/xsj/go/go1.21.1/src/runtime/debug/stack.go:24 +0x5e\ngithub.com/zeromicro/go-zero/rest/handler.RecoverHandler.func1.1()\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/recoverhandler.go:16 +0x58\npanic({0x1b5bf80?, 0xc000553a70?})\n\t/home/xsj/go/go1.21.1/src/runtime/panic.go:914 +0x21f\nmomo/app/bff/api/internal/logic.(*UserInfoLogic).UserInfo(0xc000553a40)\n\t/home/xsj/workspace/golang/src/momo/app/bff/api/internal/logic/user_info_logic.go:30 +0x1e9\nmomo/app/bff/api/internal/handler.RegisterHandlers.UserInfoHandler.func5({0x20da2b8, 0xc00013dec0}, 0xc0004bd901?)\n\t/home/xsj/workspace/golang/src/momo/app/bff/api/internal/handler/user_info_handler.go:14 +0x4a\nnet/http.HandlerFunc.ServeHTTP(0xc00028dcb0?, {0x20da2b8?, 0xc00013dec0?}, 0xc000127890?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest/handler.Authorize.func1.1({0x20da2b8, 0xc00013dec0}, 0xc0004ed200)\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/authhandler.go:81 +0x3d1\nnet/http.HandlerFunc.ServeHTTP(0x0?, {0x20da2b8?, 0xc00013dec0?}, 0x0?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest/handler.GunzipHandler.func1({0x20da2b8, 0xc00013dec0}, 0xc0004ed200)\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/gunziphandler.go:26 +0x110\nnet/http.HandlerFunc.ServeHTTP(0x798dd1edb67a79?, {0x20da2b8?, 0xc00013dec0?}, 0xc1426f65c10d7beb?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest.(*engine).buildChainWithNativeMiddlewares.MaxBytesHandler.func8.1({0x20da2b8?, 0xc00013dec0?}, 0x307ae20?)\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/maxbyteshandler.go:24 +0xf8\nnet/http.HandlerFunc.ServeHTTP(0x70?, {0x20da2b8?, 0xc00013dec0?}, 0x7f0b6e049fc8?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest.(*engine).buildChainWithNativeMiddlewares.MetricHandler.func6.1({0x20da2b8, 0xc00013dec0}, 0x1da0b00?)\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/metrichandler.go:21 +0xae\nnet/http.HandlerFunc.ServeHTTP(0xc0000827a8?, {0x20da2b8?, 0xc00013dec0?}, 0x0?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest/handler.RecoverHandler.func1({0x20da2b8?, 0xc00013dec0?}, 0xc000318340?)\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/recoverhandler.go:21 +0x78\nnet/http.HandlerFunc.ServeHTTP(0xcfbac5?, {0x20da2b8?, 0xc00013dec0?}, 0x44739c?)\n\t/home/xsj/go/go1.21.1/src/net/http/server.go:2136 +0x29\ngithub.com/zeromicro/go-zero/rest/handler.(*timeoutHandler).ServeHTTP.func1()\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/timeouthandler.go:82 +0x71\ncreated by github.com/zeromicro/go-zero/rest/handler.(*timeoutHandler).ServeHTTP in goroutine 76\n\t/home/xsj/workspace/golang/pkg/mod/github.com/zeromicro/go-zero@v1.5.6/rest/handler/timeouthandler.go:76 +0x333\n","level":"error","span":"2caccc249852178f","trace":"2c92fc0bf477ba8cbfcdcfb4c73d344a"}
{"@timestamp":"2023-10-13T21:28:55.017+08:00","caller":"handler/loghandler.go:149","content":"[HTTP] 500 - GET /v1/user/info - 127.0.0.1:33598 - PostmanRuntime/7.33.0\nGET /v1/user/info HTTP/1.1\r\nHost: localhost:8888\r\nAccept: */*\r\nAccept-Encoding: gzip, deflate, br\r\nAuthorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc4MDU5ODcsImlhdCI6MTY5NzIwMTE4NywidXNlcklkIjoxfQ.rCNEQXfZRxjQg7jXd7jHw7zk9i1tabVBH0HDtzXPOc8\r\nConnection: keep-alive\r\nCookie: token=f05ae16e-f737-4bf8-9a08-4175e5743389\r\nPostman-Token: dbe09850-5b2d-4477-938e-c951a4763345\r\nUser-Agent: PostmanRuntime/7.33.0\r\n\r\n","duration":"0.3ms","level":"error","span":"2caccc249852178f","trace":"2c92fc0bf477ba8cbfcdcfb4c73d344a"}