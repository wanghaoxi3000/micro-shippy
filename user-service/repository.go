package main

import (
	"errors"
	"fmt"
	"log"

	pb "./proto/user"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Create(*pb.User) error
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}

type UserRepository struct {
	db map[string]*pb.User
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	if u, ok := repo.db[id]; ok {
		return u, nil
	}
	return nil, errors.New("user not find")
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	users := make([]*pb.User, len(repo.db))
	i := 0
	for k := range repo.db {
		users[i] = repo.db[k]
		i++
	}
	return users, nil
}

func (repo *UserRepository) Create(u *pb.User) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("created uuid error: %v", err)
	}

	u.Id = uuid.String()
	log.Printf("Create user %v", u.Id)
	repo.db[u.Id] = u
	return nil
}

func (repo *UserRepository) GetByEmailAndPassword(u *pb.User) (*pb.User, error) {
	for _, v := range repo.db {
		if v.Email == u.Email {
			return v, nil
		}
	}
	return nil, fmt.Errorf("not find user %v", u.Email)
}
