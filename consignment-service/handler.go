package main

import (
	"context"
	"log"

	pb "./proto/consignment"
	vesselPb "./proto/vessel"
)

// 微服务服务端 struct handler 必须实现 protobuf 中定义的 rpc 方法
// 实现方法的传参等可参考生成的 consignment.pb.go
type handler struct {
	repo         Repository
	vesselClient vesselPb.VesselServiceClient
}

// 托运新的货物
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	log.Printf("receive a consignment\n")

	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := h.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	// 货物被承运
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	err = h.repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}
