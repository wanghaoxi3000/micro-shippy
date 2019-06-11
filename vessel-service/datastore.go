package main

// CreateStore 创建存储 store
func CreateStore() (*VesselRepository, error) {
	store := new(VesselRepository)
	return store, nil
}
