// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http.
const _ = http1.SupportPackageIsVersion1

type CommentService interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentReply, error)

	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentReply, error)

	GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error)

	ListComment(context.Context, *ListCommentRequest) (*ListCommentReply, error)

	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentReply, error)
}

func RegisterCommentHTTPServer(s *http1.Server, srv CommentService) {
	r := s.Route("/")

	r.POST("/api.blog.v1.Comment/CreateComment", func(res http.ResponseWriter, req *http.Request) {
		in := new(CreateCommentRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(CommentService).CreateComment(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Comment/UpdateComment", func(res http.ResponseWriter, req *http.Request) {
		in := new(UpdateCommentRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(CommentService).UpdateComment(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Comment/DeleteComment", func(res http.ResponseWriter, req *http.Request) {
		in := new(DeleteCommentRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(CommentService).DeleteComment(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Comment/GetComment", func(res http.ResponseWriter, req *http.Request) {
		in := new(GetCommentRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(CommentService).GetComment(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Comment/ListComment", func(res http.ResponseWriter, req *http.Request) {
		in := new(ListCommentRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(CommentService).ListComment(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

}
