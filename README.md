# Backend API

A gRPC based api.

## Protobuf Generation

```shell
protoc --go_out=pb --go_opt=paths=source_relative \  john@wooly
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/route.proto
```