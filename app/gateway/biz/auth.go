package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthRepo interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

type LoginRequest struct {
	Id string
}
type LoginReply struct {
	Token string
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUseCase) Login(ctx context.Context, r *LoginRequest) (*LoginReply, error) {
	uc.log.WithContext(ctx).Infof("Login: %v", r.Id)
	return uc.repo.Login(ctx, r)
}
