package service

import (
	"context"
	"github.com/google/uuid"
	v1 "nacostest/api/user/v1"
)

// AuthService is a auth service.
type AuthService struct {
	v1.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{Token: uuid.New().String()}, nil
}
