.PHONY: ptm
ptm: linux_env user


# linux 环境变量
linux_env:
	set GOARCH=amd64
	go env -w GOARCH=amd64
	set GOOS=linux
	go env -w GOOS=linux


# tpmt微服务
user:
	go build -o deploy/app/out/user/rpc/user-rpc service/user/rpc/user.go
	go build -o deploy/app/out/user/api/user-api service/user/api/user.go
