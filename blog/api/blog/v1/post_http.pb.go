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

type PostService interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error)

	DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error)

	GetPost(context.Context, *GetPostRequest) (*GetPostReply, error)

	ListPost(context.Context, *ListPostRequest) (*ListPostReply, error)

	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error)
}

func RegisterPostHTTPServer(s *http1.Server, srv PostService) {
	r := s.Route("/")

	r.POST("/api.blog.v1.Post/CreatePost", func(res http.ResponseWriter, req *http.Request) {
		in := new(CreatePostRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(PostService).CreatePost(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Post/UpdatePost", func(res http.ResponseWriter, req *http.Request) {
		in := new(UpdatePostRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(PostService).UpdatePost(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Post/DeletePost", func(res http.ResponseWriter, req *http.Request) {
		in := new(DeletePostRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(PostService).DeletePost(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Post/GetPost", func(res http.ResponseWriter, req *http.Request) {
		in := new(GetPostRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(PostService).GetPost(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

	r.POST("/api.blog.v1.Post/ListPost", func(res http.ResponseWriter, req *http.Request) {
		in := new(ListPostRequest)

		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(PostService).ListPost(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out)
	})

}
