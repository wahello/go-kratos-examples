package data

import (
	"context"

	"github.com/go-kratos/examples/blog/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPostRepo)

// Config is data config.
type Config struct {
	Driver string
	Source string
}

// Data .
type Data struct {
	*ent.Client
}

// NewData .
func NewData(conf *Config, logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)
	client, err := ent.Open(conf.Driver, conf.Source)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}
	return &Data{
		Client: client,
	}, nil
}
