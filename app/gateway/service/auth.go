package service

import (
	"context"
	v1 "nacostest/api/gateway/v1"
	"nacostest/app/gateway/biz"
)

// AuthService is a auth service.
type AuthService struct {
	uc *biz.AuthUseCase
}

func NewAuthService(uc *biz.AuthUseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (out *v1.LoginReply, err error) {
	sret, err := s.uc.Login(ctx, &biz.LoginRequest{Id: in.Id})
	if err != nil {
		return
	}
	return &v1.LoginReply{Token: sret.Token}, nil
}
