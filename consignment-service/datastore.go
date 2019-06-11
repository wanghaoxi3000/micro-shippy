package main

// CreateStore 创建存储 store
func CreateStore() (*ConsignmentRepository, error) {
	store := new(ConsignmentRepository)
	return store, nil
}
