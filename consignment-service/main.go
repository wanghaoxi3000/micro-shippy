package main // import "github.com/wanghaoxi3000/micro-shippy"

import (
	"log"

	pb "./proto/consignment"
	// pb "github.com/wanghaoxi3000/micro-shippy/proto/consignment"

	vesselPb "./proto/vessel"
	// vesselPb "github.com/wanghaoxi3000/micro-shippy/proto/vessel"

	"github.com/micro/go-micro"
)

func main() {
	repo, err := CreateStore()
	defer repo.Close()
	if err != nil {
		log.Fatalf("create store error: %v\n", err)
	}

	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	server.Init()
	log.Printf("micro start: go.micro.srv.consignment\n")

	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{repo, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
