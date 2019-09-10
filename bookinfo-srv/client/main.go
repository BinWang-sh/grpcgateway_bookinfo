package main

import (
    pb "grpcT1/bookinfo-srv/proto/bookinfo"
    "google.golang.org/grpc"
    "log"
    "context"
)

const (
    ADDRESS = "rpcserver:50051"
)

func main() {
    // 连接到 gRPC 服务器
    conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("connect error: %v", err)
    }
    defer conn.Close()

    // 初始化 gRPC 客户端
    client := pb.NewBookinfoClient(conn)

    resp, err := client.Getall(context.Background(), &pb.GetallRequest{Userid:"tt"})
    if err != nil {
        log.Fatalf("Getall error: %v", err)
    }

	length := len(resp.Bookinfolist)
	for i:=0;i<length;i++ {
		log.Printf("bookName:%s\n",resp.Bookinfolist[i].BookName)
		log.Printf("author:%s\n", resp.Bookinfolist[i].Author)
		
		chaptersCount := len(resp.Bookinfolist[i].ChaptersInfo)
		for j:=0;j<chaptersCount;j++ {
			log.Printf("---Chapter.Num:%d\n", resp.Bookinfolist[i].ChaptersInfo[j].ChapterNum)
			log.Printf("---Chapter.Name:%s\n", resp.Bookinfolist[i].ChaptersInfo[j].ChapterName)
			log.Printf("---Chapter.WordsCount:%d\n", resp.Bookinfolist[i].ChaptersInfo[j].WordsCount)
			log.Println("")
		}
	}
}