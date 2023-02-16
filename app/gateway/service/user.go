package service

import (
	"context"
	v1 "nacostest/api/gateway/v1"
	"nacostest/app/gateway/biz"
)

// UserService is a greeter service.
type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Find(ctx context.Context, in *v1.FindRequest) (out *v1.FindReply, err error) {
	sret, err := s.uc.Find(ctx, &biz.FindRequest{Id: in.Id})
	if err != nil {
		return
	}
	return &v1.FindReply{Id: sret.Id, Name: sret.Name}, nil
}
