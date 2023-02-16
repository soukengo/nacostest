package service

import (
	"context"

	v1 "nacostest/api/user/v1"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Find(ctx context.Context, in *v1.FindRequest) (*v1.FindReply, error) {
	return &v1.FindReply{Id: in.Id, Name: "用户" + in.Id}, nil
}
