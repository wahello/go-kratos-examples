// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/go-kratos/examples/blog/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// InitApp .
func InitApp(*Service, *server.HTTPConfig, *server.GRPCConfig, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, NewApp))
}
