#/bin/bash
# This is how we want to name the binary output

# 网关模块--gin
GW_SRC_PROTO=cmd/main.go
# db 模块
DB_SRC_PROTO=module/db/main.go
# public 模块
PUBLIC_SRC_PROTO=module/public/main.go


# db模块 输出
DB_OUTPUT=module/bin/server_db
# 公共微服务
PUBLIC_OUTPUT=module/bin/server_public
# 网关输出
GW_OUTPUT=bin/gin_micro

# These are the values we want to pass for Version and BuildTime
GIT_TAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`



# Setup the -ldflags option for go build here, interpolate the variable values
LDF_LAGS=-ldflags "-X main.Version=${GIT_TAG} -X main.Build_Time=${BUILD_TIME} -s -w"

local:
	rm -f ./bin/gin_*
	rm -f ./module/bin/server_*
	go build ${LDF_LAGS} -o ${GW_OUTPUT} ${GW_SRC_PROTO}
	go build ${LDF_LAGS} -o ${DB_OUTPUT} ${DB_SRC_PROTO}
	go build ${LDF_LAGS} -o ${PUBLIC_OUTPUT} ${PUBLIC_SRC_PROTO}

proto:
	protoc --proto_path=module/db/proto   --go_out=module/db/proto --micro_out=module/db/proto module/db/proto/db.proto
	protoc --proto_path=module/public/proto   --go_out=module/public/proto --micro_out=module/public/proto module/public/proto/public.proto

web-server:
	rm -f ./bin/gin_*
	rm -f ./module/bin/server_*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDF_LAGS} -o ${GW_OUTPUT} ${GW_SRC_PROTO}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDF_LAGS} -o ${DB_OUTPUT} ${DB_SRC_PROTO}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter ${LDF_LAGS} -o ${PUBLIC_OUTPUT} ${PUBLIC_SRC_PROTO}

clean:
	rm -f ./bin/gin_*
	rm -f ./module/bin/server_*
