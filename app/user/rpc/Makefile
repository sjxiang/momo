
# rpc 代码生成
.PHONY: rpc
rpc:
	goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero


# model 代码生成
.PHONY: model
model:
	goctl model mysql ddl -src ./internal/model/*.sql -dir ./internal/model -cache=true --style=go_zero


# 启动 user-rpc 服务
.PHONY: run
run:
	go run user.go 
