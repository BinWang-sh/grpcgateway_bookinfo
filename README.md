# grpcgateway_bookinfo
A test project of using grpc gateway

## prepare
```bash
go get -u google.golang.org/grpc
```

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

```bash
go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
```

## Build
```bash
cd grpcgateway_bookinfo
```

Generate *.pb.go file using grpc plugin
```bash
make genstubgrpc 
```

Generate *.pb.go file using micro plugin
```bash
make genstubmicro  
```

Generate grpc gateway *.pb.gw.go file
```bash
make gengateway
```

Build 
```bash
make build
```


## Run with docker
```bash
cd Docker
docker-compose up
```

1. rpcserver listen 50051
2. gwserver listen 8080
3. client using grpc visit rpcserver

Using curl to test http request
```bash
curl -X POST -k http://localhost:8080/bookinfo/getall -d '{"userid":"tt"}'
```
