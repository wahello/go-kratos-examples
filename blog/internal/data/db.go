package data

import (
	"context"

	"github.com/go-kratos/examples/blog/internal/biz"
	"github.com/go-kratos/examples/blog/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Driver string
	Source string
}

type postRepo struct {
	db  *db
	log *log.Helper
}

type db struct {
	client *ent.Client
	log    *log.Helper
}

func NewDB(conf *DBConfig, logger log.Logger) *db {
	l := log.NewHelper("db", logger)
	client, err := ent.Open(conf.Driver, conf.Source)
	if err != nil {
		l.Errorf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		l.Errorf("failed creating schema resources: %v", err)
	}
	return &db{
		client: client,
		log:    l,
	}
}

func NewPostRepo(db *db, logger log.Logger) biz.PostRepo {
	return &postRepo{
		db:  db,
		log: log.NewHelper("post_repo", logger),
	}
}

// implement biz.PostRepo
func (pr *postRepo) CreatePost(post *biz.Post) error {
	return nil
}
func (pr *postRepo) UpdatePost(post *biz.Post) error {
	return nil
}
