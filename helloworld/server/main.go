package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/log/stdlog"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "error" {
		return nil, errors.InvalidArgument("BadRequest", "invalid argument %s", in.Name)
	}
	if in.Name == "panic" {
		panic("grpc panic")
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in)}, nil
}

func main() {
	logger := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	defer logger.Close()

	log := log.NewHelper("main", logger)

	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			middleware.Chain(
				logging.HTTPServer(logger),
				recovery.Recovery(),
			),
		))
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			middleware.Chain(
				logging.GRPCServer(logger),
				status.Server(),
				recovery.Recovery(),
			),
		))

	s := &server{}
	pb.RegisterGreeterServer(grpcSrv, s)
	pb.RegisterGreeterHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
