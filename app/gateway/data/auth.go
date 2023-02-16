package data

import (
	"context"
	v1 "nacostest/api/user/v1"
	"nacostest/app/gateway/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	log *log.Helper
	ac  v1.AuthClient
}

func NewAuthRepo(ac v1.AuthClient, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		ac:  ac,
		log: log.NewHelper(logger),
	}
}

func (r *authRepo) Login(ctx context.Context, in *biz.LoginRequest) (ret *biz.LoginReply, err error) {
	sret, err := r.ac.Login(ctx, &v1.LoginRequest{
		Id: in.Id,
	})
	if err != nil {
		return
	}
	return &biz.LoginReply{Token: sret.Token}, nil
}
