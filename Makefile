LOCAL_BIN:=$(CURDIR)/bin


.PHONY: deps
deps:
	ls go.mod || go mod init github.com/genridarkbkru/registration
	GOBIN=$(LOCAL_BIN) go get -d github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -d github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -d github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -d google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: build
build:
	go build main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: clearbd
clearbd:
	migrate -path migrations -database "postgres://postgres:1@localhost:5432/grpc_dev?sslmode=disable" down
	migrate -path migrations -database "postgres://postgres:1@localhost:5432/grpc_dev?sslmode=disable" up


.PHONY: generate
generate:
	GOBIN=$(LOCAL_BIN) protoc -I internal/app/grpc_server \
    			--go_out=internal/app/grpc_server --go_opt=paths=import \
	   			--go-grpc_out=internal/app/grpc_server --go-grpc_opt=paths=import \
	   			--grpc-gateway_out=internal/app/grpc_server \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
                --swagger_out=allow_merge=true,merge_file_name=api:swagger \
				internal/app/grpc_server/grpc_server.proto

