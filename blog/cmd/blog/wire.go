// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog/internal/biz"
	"blog/internal/data"

	"github.com/go-kratos/kratos/v2"
	grpcconf "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	httpconf "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func NewApp(hc *httpconf.Server, gc *grpcconf.Server, dc *data.DBConfig, logger log.Logger) (*kratos.App, error) {
	panic(wire.Build(NewPostRepo, postServiceProvider, httpServerProvider, grpcServerProvider, newApp))
}

func NewPostRepo(conf *data.DBConfig, logger log.Logger) (biz.PostRepo, error) {
	panic(wire.Build(data.NewDB, data.NewPostRepo))
}
