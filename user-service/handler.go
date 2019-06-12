package main

import (
	"context"
	"errors"

	pb "./proto/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repo         Repository
	tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	// 哈希处理用户输入的密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPwd)

	if err := h.repo.Create(req); err != nil {
		return nil
	}
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	u, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	// 进行密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}
	t, err := h.tokenService.Encode(u)
	if err != nil {
		return err
	}

	resp.Token = t
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}
