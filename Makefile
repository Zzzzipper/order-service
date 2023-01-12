GOPATH:=$(shell go env GOPATH)

all:
	protoc -I/usr/local/include -I. \
			-I${GOPATH}/pkg/mod \
			-I${GOPATH}/pkg/mod/gitlab.mapcard.pro/external-map-team/api-proto@v1.0.73 \
			--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. \
			--go_out=paths=source_relative:. \
			./proto/*.proto
