package grpc

import (
	"github.com/RichFastTrade/rft_shared/conf"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"sync"
)

var server *grpc.Server
var once sync.Once

func Get() *grpc.Server {
	once.Do(func() {
		server = grpc.NewServer(
			grpc.Address(conf.Get().Server.Grpc.Addr),
		)
	})
	return server
}
