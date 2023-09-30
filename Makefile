TARGET=.

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go

proto:
	protoc protoc --proto_path=proto --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative --go_out=./proto/pb --go_opt=paths=source_relative --grpc-gateway_out ./proto/pb --grpc-gateway_opt paths=source_relative proto/*/*.proto