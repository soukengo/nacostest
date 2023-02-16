package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	Find(context.Context, *FindRequest) (*FindReply, error)
}

type FindRequest struct {
	Id string
}
type FindReply struct {
	Id   string
	Name string
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Find(ctx context.Context, r *FindRequest) (*FindReply, error) {
	uc.log.WithContext(ctx).Infof("Find: %v", r.Id)
	return uc.repo.Find(ctx, r)
}
