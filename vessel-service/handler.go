package main

import (
	"context"

	pb "./proto/vessel"
)

// 定义货船服务
type handler struct {
	repo *VesselRepository
}

// 实现微服务的服务端
func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}

// 实现服务端
func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	// 调用内部方法查找
	v, err := h.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}
