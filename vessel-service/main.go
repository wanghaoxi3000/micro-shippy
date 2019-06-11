package main // import "github.com/wanghaoxi3000/micro-shippy/vessel-service"

import (
	"log"

	pb "./proto/vessel"
	// pb "github.com/wanghaoxi3000/micro-shippy/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

func main() {
	repo, err := CreateStore()
	if err != nil {
		log.Fatalf("create store error: %v\n", err)
	}

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	// 将实现服务端的 API 注册到服务端
	pb.RegisterVesselServiceHandler(server.Server(), &handler{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
