# goWebApi

server module=server

go get -u google.golang.org/grpc

go build server
go run server

# Proto
# cd to proto directory
protoc greeting.proto --go_out=. --go_opt=paths=source_relative

Optional: 
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
