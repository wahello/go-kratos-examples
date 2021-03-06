// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/examples/blog/internal/biz"
	"github.com/go-kratos/examples/blog/internal/data"
	"github.com/go-kratos/examples/blog/internal/server"
	"github.com/go-kratos/examples/blog/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// InitApp .
func InitApp(httpConfig *server.HTTPConfig, grpcConfig *server.GRPCConfig, config *data.Config, logger log.Logger) (*kratos.App, error) {
	httpServer := server.NewHTTPServer(httpConfig, logger)
	grpcServer := server.NewGRPCServer(grpcConfig, logger)
	dataData, err := data.NewData(config, logger)
	if err != nil {
		return nil, err
	}
	postRepo := data.NewPostRepo(dataData, logger)
	postUsecase := biz.NewPostUsecase(postRepo, logger)
	postService := service.NewPostService(postUsecase, logger)
	app := newApp(logger, httpServer, grpcServer, postService)
	return app, nil
}
