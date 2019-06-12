package main

import (
	"log"

	"github.com/micro/go-micro"

	pb "./proto/user"
)

func main() {
	repo, err := CreateStore()
	if err != nil {
		log.Fatalf("create store error: %v\n", err)
	}

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	s.Init()

	t := TokenService{repo}
	pb.RegisterUserServiceHandler(s.Server(), &handler{repo, &t})
	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}
}
