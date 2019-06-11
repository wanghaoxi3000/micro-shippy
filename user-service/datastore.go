package main

import pb "./proto/user"

// CreateStore 创建存储 store
func CreateStore() (Repository, error) {
	store := new(UserRepository)
	store.db = make(map[string]*pb.User)

	return store, nil
}
