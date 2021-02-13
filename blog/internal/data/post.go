package data

import (
	"github.com/go-kratos/examples/blog/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type postRepo struct {
	data *Data
	log  *log.Helper
}

// NewPostRepo .
func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.NewHelper("post_repo", logger),
	}
}

func (pr *postRepo) CreatePost(post *biz.Post) error {
	return nil
}

func (pr *postRepo) UpdatePost(post *biz.Post) error {
	return nil
}
