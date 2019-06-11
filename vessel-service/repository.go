package main

import (
	"errors"

	pb "./proto/vessel"
)

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	repo.vessels = append(repo.vessels, v)
	return nil
}

// FindAvailable 选择最近一条容量、载重都符合的货轮
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("No vessel can't be use")
}
