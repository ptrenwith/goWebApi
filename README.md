# go_api

This is a skeleton for a Go gRPC API with a REST interface for external HTTP comms

project module=github.com/ptrenwith/go_api

go get -u google.golang.org/grpc

go build server
go run server

# Proto
# cd to root directory
# optionally add --go-grpc_opt=require_unimplemented_servers=false if you do not want forward compatibility
protoc --proto_path=proto --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative --go_out=./proto/pb --go_opt=paths=source_relative --grpc-gateway_out ./proto/pb --grpc-gateway_opt paths=source_relative proto/*/*.proto 

# To enabled REST into gRPC make use of grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
protoc --proto_path=messages --go_out=./messages/pb --go_opt=paths=source_relative --go-grpc_out=./messages/pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=./messages/pb --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_out=version=true:./messages/pb messages/*.proto

# from example
protoc messages/*.proto \
            --grpc-gateway_out . \
            --grpc-gateway_opt logtostderr=true \
            --grpc-gateway_opt paths=source_relative \
            --grpc-gateway_out=version=true:.

            --grpc-gateway_opt generate_unbound_methods=true \

Optional: 
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
