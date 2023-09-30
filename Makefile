TARGET=.

all: clean proto build

clean:
	rm -rf $(TARGET)/proto/pb/

build-only:
	go build -o $(TARGET)/go_api main.go

proto: clean
	mkdir $(TARGET)/proto/pb/
	protoc --proto_path=proto --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative --go_out=./proto/pb --go_opt=paths=source_relative --grpc-gateway_out ./proto/pb --grpc-gateway_opt paths=source_relative proto/*/*.proto

build: proto build-only

.PHONY: all proto clean