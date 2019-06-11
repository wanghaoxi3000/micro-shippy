package main

import pb "./proto/consignment"

// Repository 仓库接口
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository 模拟数据库存放货物的仓库
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create create a consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	repo.consignments = append(repo.consignments, consignment)
	return nil
}

// GetAll get all consignment
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	return repo.consignments, nil
}

func (repo *ConsignmentRepository) Close() {
	return
}
