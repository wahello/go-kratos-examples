package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Post struct {
	Id      int64
	Title   string
	Content string
	CTime   time.Time
}

type PostRepo interface {
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
}

type PostUsecase struct {
	repo PostRepo
}

func NewPostUsecase(repo PostRepo, logger log.Logger) *PostUsecase {
	return &PostUsecase{repo: repo}
}

func (uc *PostUsecase) Create(post *Post) error {
	return uc.repo.CreatePost(post)
}

func (uc *PostUsecase) Update(post *Post) error {
	return uc.repo.UpdatePost(post)
}
