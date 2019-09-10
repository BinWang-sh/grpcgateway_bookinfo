package main

import (
	pb "grpcT1/bookinfo-srv/proto/bookinfo"
	bkinfoprocess "grpcT1/bookinfo-srv/process"
    "net"
    "log"
	"google.golang.org/grpc"
)
const (
	PORT = ":50051"
)


func main() {
	listener, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    log.Printf("listen on: %s\n", PORT)

    server := grpc.NewServer()


    pb.RegisterBookinfoServer(server, bkinfoprocess.NewBookinfo())

    if err := server.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
	}

}