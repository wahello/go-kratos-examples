package main

import (
	"flag"
	"os"

	"github.com/go-kratos/examples/blog/internal/di"
	"github.com/go-kratos/examples/blog/internal/server"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/log/stdlog"
	"gopkg.in/yaml.v2"
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
	logger := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	defer logger.Close()

	conf := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
		config.WithLogger(logger),
	)
	if err := conf.Load(); err != nil {
		panic(err)
	}

	l := log.NewHelper("main", logger)

	var (
		sc di.Service
		hc server.HTTPConfig
		gc server.GRPCConfig
	)

	if err := conf.Value("service").Scan(&sc); err != nil {
		panic(err)
	}
	if err := conf.Value("http.server").Scan(&hc); err != nil {
		panic(err)
	}
	if err := conf.Value("grpc.server").Scan(&gc); err != nil {
		panic(err)
	}

	// application lifecycle
	app, err := di.InitApp(&sc, &hc, &gc, logger)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		l.Errorf("start failed: %v\n", err)
	}
}
