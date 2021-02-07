package service

import(
	"context"

	pb "blog/api/blog/v1"
)

type PostService struct {
	pb.UnimplementedPostServer
}

func NewPostService() pb.PostServer {
	return &PostService{}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	return &pb.CreatePostReply{}, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	return &pb.UpdatePostReply{}, nil
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
