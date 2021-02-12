package service

import (
	"context"

	pb "github.com/go-kratos/examples/blog/api/blog/v1"
	"github.com/go-kratos/examples/blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type PostService struct {
	pb.UnimplementedPostServer

	log *log.Helper

	puc *biz.PostUsecase
}

func NewPostService(pr biz.PostRepo, logger log.Logger) pb.PostServer {
	return &PostService{
		puc: biz.NewPostUsecase(pr, logger),
		log: log.NewHelper("post", logger),
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	err := s.puc.Create(&biz.Post{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.CreatePostReply{}, err
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	err := s.puc.Create(&biz.Post{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.UpdatePostReply{}, err
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	return &pb.DeletePostReply{}, nil
}
func (s *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	return &pb.GetPostReply{}, nil
}
func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	return &pb.ListPostReply{}, nil
}
