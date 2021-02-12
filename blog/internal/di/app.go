package di

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Service is service config.
type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewApp .
func NewApp(s *Service, hs *http.Server, gs *grpc.Server, logger log.Logger) *kratos.App {
	return kratos.New(
		kratos.Name(s.Name),
		kratos.Version(s.Version),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}
