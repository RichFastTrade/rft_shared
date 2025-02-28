API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
			--proto_path=./third_party \
			--go_out=paths=source_relative:./api \
			--go-http_out=paths=source_relative:./api \
			--go-grpc_out=paths=source_relative:./api \
			$(API_PROTO_FILES)