package main

import (
	pb "grpcT1/bookinfo-srv/proto/bookinfo_micro"
	bkinfoprocess "grpcT1/bookinfo-srv/process"

    "log"
	"github.com/micro/go-micro"
)
const (
	PORT = ":50051"
)

func main() {
	/*
	listener, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    log.Printf("listen on: %s\n", PORT)

    server := grpc.NewServer()
	*/

	server := micro.NewService(
		micro.Name("go.micro.srv.bookinfoservice"),
		micro.Version("latest"),
	)

	server.Init()

    pb.RegisterBookinfoHandler(server.Server(), bkinfoprocess.NewBookinfo())

	/*
    if err := server.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
	}
	*/

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}