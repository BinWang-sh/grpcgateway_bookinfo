package main

import (
  "flag"
  "net/http"
   
  "github.com/golang/glog"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
   	
  gw "grpcT1/bookinfo-srv/proto/bookinfo"
)
   
var (
  echoEndpoint = flag.String("getall_endpoint", "rpcserver:50051", "endpoint of BookInfoService")
)
   
func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()
   
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterBookinfoHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
  if err != nil {
    return err
  }
   
  return http.ListenAndServe(":8080", mux)
}
   
func main() {
  flag.Parse()
  defer glog.Flush()
   
  if err := run(); err != nil {
    glog.Fatal(err)
  }
}