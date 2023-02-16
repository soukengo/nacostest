package data

import (
	"github.com/google/wire"
	"nacostest/app/gateway/data/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(grpc.ProviderSet, NewAuthRepo, NewUserRepo)
