cleanproto:
	rm bookinfo-srv/proto/bookinfo/*.go
	rm bookinfo-srv/proto/bookinfo_micro/*.go

cleanapp:
	rm ./Docker/service/servicegrpc
	rm ./Docker/client/client
	rm ./Docker/gateway/gateway

genstubgrpc:
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.1/third_party/googleapis \
	--go_out=plugins=grpc:. bookinfo-srv/proto/bookinfo/bookinfo.proto

genstubmicro:
	protoc -I/usr/local/include -I. \
	--go_out=plugins=micro:. bookinfo-srv/proto/bookinfo_micro/bookinfo.proto

gengateway:
	protoc -I/usr/local/include -I. \
	-I $(GOPATH)/src \
	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. bookinfo-srv/proto/bookinfo/bookinfo.proto

build:
	mkdir -p ./bin
	GOOS=linux GOARCH=amd64 go build -o ./Docker/service/servicegrpc ./bookinfo-srv/service_grpc
	#GOOS=linux GOARCH=amd64 go build -o ./service/servicemicro ./bookinfo-srv/service_micro
	GOOS=linux GOARCH=amd64 go build -o ./Docker/client/client ./bookinfo-srv/client
	GOOS=linux GOARCH=amd64 go build -o ./Docker/gateway/gateway ./bookinfo-srv/gateway

	docker build -t bookinfoservice ./Docker/service
	docker build -t bookinfogateway ./Docker/gateway
	docker build -t bookinfoclient ./Docker/client
