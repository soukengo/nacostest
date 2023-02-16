package data

import (
	"context"
	v1 "nacostest/api/user/v1"
	"nacostest/app/gateway/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	log *log.Helper
	ac  v1.UserClient
}

func NewUserRepo(ac v1.UserClient, logger log.Logger) biz.UserRepo {
	return &userRepo{
		ac:  ac,
		log: log.NewHelper(logger),
	}
}

func (r *userRepo) Find(ctx context.Context, in *biz.FindRequest) (ret *biz.FindReply, err error) {
	sret, err := r.ac.Find(ctx, &v1.FindRequest{
		Id: in.Id,
	})
	if err != nil {
		return
	}
	return &biz.FindReply{Id: sret.Id, Name: sret.Name}, nil
}
