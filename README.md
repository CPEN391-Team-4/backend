# Backend API

A gRPC based api.

## Protobuf Generation

```shell
export GO111MODULE=on 
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/route.proto
```

## DB

Setup MySQL/MariaDB

Create user and database:

```sql
CREATE USER 'cpen391'@'localhost' IDENTIFIED BY '******';
CREATE DATABASE cpen391_backend;
GRANT ALL PRIVILEGES ON cpen391_backend.* TO 'cpen391'@'localhost';

USE cpen391_backend;
create table table_name
(
    name varchar(255) null,
    image_id varchar(255) null,
    restricted bool null
);
```