package main

import (
	"flag"
	"os"

	pb "blog/api/blog/v1"
	"blog/internal/service"

	"github.com/go-kratos/kratos/v2"
	grpcconf "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	httpconf "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/source/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/log/stdlog"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	conf := config.New(config.WithSource(
		file.NewSource(flagconf),
	))
	if err := conf.Load(); err != nil {
		panic(err)
	}

	logger := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	defer logger.Close()

	log := log.NewHelper("main", logger)

	// build transport server
	hc := new(httpconf.Server)
	gc := new(grpcconf.Server)
	if err := conf.Value("http.server").Scan(hc); err != nil {
		panic(err)
	}
	if err := conf.Value("grpc.server").Scan(gc); err != nil {
		panic(err)
	}
	httpSrv := http.NewServer(http.Apply(hc), http.Logger(logger))
	grpcSrv := grpc.NewServer(grpc.Apply(gc), grpc.Logger(logger))

	// register service
	gs := service.NewPostService()
	pb.RegisterPostServer(grpcSrv, gs)
	pb.RegisterPostHTTPServer(httpSrv, gs)

	// application lifecycle
	app := kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Errorf("start failed: %v\n", err)
	}
}
