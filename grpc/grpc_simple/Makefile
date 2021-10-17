all:
        protoc -I/usr/local/go/pkg/include -I. \
            -I${GOPATH}/pkg/mod \
            -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
            -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
            --grpc-gateway_out=logtostderr=true:./api_pb \
            --swagger_out=allow_merge=true,merge_file_name=api:. \
            --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ./api/unit/unit.proto