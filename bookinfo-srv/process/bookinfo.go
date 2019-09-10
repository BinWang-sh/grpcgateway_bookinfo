package bookinfoprocess

import (
	bookinfo "grpcT1/bookinfo-srv/proto/bookinfo"
	"context"
	"fmt"
)

type bookinfosrvc struct {}

func NewBookinfo() bookinfo.BookinfoServer {
	return &bookinfosrvc{}
}

func (s *bookinfosrvc) Getall(ctx context.Context, p *bookinfo.GetallRequest) (*bookinfo.GetallResponse, error) {
	res := &bookinfo.GetallResponse{}
	fmt.Println("bookinfo.getall")
	cpsInfo := make([]*bookinfo.ChapterInfo,4)
	cpsInfo[0] = &bookinfo.ChapterInfo{ChapterNum:1, ChapterName:"序言", WordsCount:3259}
	cpsInfo[1] = &bookinfo.ChapterInfo{ChapterNum:2, ChapterName:"布尔", WordsCount:4559}
	cpsInfo[2] = &bookinfo.ChapterInfo{ChapterNum:3, ChapterName:"字符", WordsCount:7559}
	cpsInfo[3] = &bookinfo.ChapterInfo{ChapterNum:4, ChapterName:"函数", WordsCount:7859}

	bkInfo := &bookinfo.BookInfo{BookName:"Ugly language", Author:"Bill", 
						ChaptersInfo:cpsInfo}

	res.Bookinfolist = append(res.Bookinfolist, bkInfo)

	return res, nil
}


func (s *bookinfosrvc) Add(ctx context.Context, p *bookinfo.AddRequest) (*bookinfo.AddResponse, error) {

	return nil, nil
}