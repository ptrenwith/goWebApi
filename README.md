# go_api

This is a skeleton for a Go gRPC API with a REST interface for external HTTP comms

project module=github.com/ptrenwith/go_api

go get -u google.golang.org/grpc

go build server
go run server

# Proto
# cd to proto directory
protoc greeting.proto --go_out=./pb --go_opt=paths=source_relative

Optional: 
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
