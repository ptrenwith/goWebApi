# go_api

This is a skeleton for a Go gRPC API with a REST interface for external HTTP comms

project module=github.com/ptrenwith/go_api

go get -u google.golang.org/grpc

go build server
go run server

# Proto
# cd to root directory
# optionally add --go-grpc_opt=require_unimplemented_servers=false if you do not want forward compatibility
protoc --proto_path=messages --go-grpc_out=messages/pb --go-grpc_opt=paths=source_relative --go_out=./messages/pb --go_opt=paths=source_relative messages/*.proto 

Optional: 
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
