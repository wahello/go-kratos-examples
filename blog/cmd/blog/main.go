package main

import (
	"blog/internal/data"
	"flag"
	"os"

	grpcconf "github.com/go-kratos/kratos/v2/api/kratos/config/grpc"
	httpconf "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/source/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/log/stdlog"
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

	l := log.NewHelper("main", logger)

	// build transport server
	hc := new(httpconf.Server)
	gc := new(grpcconf.Server)
	dc := new(data.DBConfig)

	if err := conf.Value("http.server").Scan(hc); err != nil {
		panic(err)
	}
	if err := conf.Value("grpc.server").Scan(gc); err != nil {
		panic(err)
	}
	if err := conf.Value("db").Scan(dc); err != nil {
		panic(err)
	}

	// application lifecycle
	app, err := NewApp(hc, gc, dc, logger)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		l.Errorf("start failed: %v\n", err)
	}
}
