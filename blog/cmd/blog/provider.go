package main

import (
	"blog/api/blog/v1"
	"blog/internal/biz"
	"blog/internal/service"
	"github.com/go-kratos/kratos/v2"
	grpcconf "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	httpconf "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func newApp(svc v1.PostService, h *http.Server, g *grpc.Server) (app *kratos.App) {
	app = kratos.New(
		kratos.Server(
			h,
			g,
		),
	)
	return
}

func postServiceProvider(h *http.Server, g *grpc.Server, pr biz.PostRepo, logger log.Logger) v1.PostService {
	svc := service.NewPostService(pr, logger)
	v1.RegisterPostServer(g, svc)
	v1.RegisterPostHTTPServer(h, svc)
	return svc
}

func httpServerProvider(hc *httpconf.Server, logger log.Logger) *http.Server {
	return http.NewServer(http.Apply(hc), http.Logger(logger))

}

func grpcServerProvider(gc *grpcconf.Server, logger log.Logger) *grpc.Server {
	return grpc.NewServer(grpc.Apply(gc), grpc.Logger(logger))
}
