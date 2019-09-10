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
	server := micro.NewService(
		micro.Name("go.micro.srv.bookinfoservice"),
		micro.Version("latest"),
	)

	server.Init()

    pb.RegisterBookinfoHandler(server.Server(), bkinfoprocess.NewBookinfo())

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}