# grpc-clean
This is my first Clean Architecture gRPC repository

## Generate client and server code

```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/grpc_clean.proto
```

Reference: https://grpc.io/docs/languages/go/basics/#generating-client-and-server-code
