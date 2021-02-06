package service

import (
	"context"

	pb "github.com/go-kratos/examples/blog/api/blog/v1"
)

type CommentService struct {
	pb.UnimplementedCommentServer
}

func NewCommentService() pb.CommentServer {
	return &CommentService{}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	return &pb.CreateCommentReply{}, nil
}
func (s *CommentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentReply, error) {
	return &pb.UpdateCommentReply{}, nil
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, nil
}
func (s *CommentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	return &pb.GetCommentReply{}, nil
}
func (s *CommentService) ListComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	return &pb.ListCommentReply{}, nil
}
