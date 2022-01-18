package rpc

import (
	"context"
	"net"

	pb "github.com/quanxiang-cloud/structor/api/proto"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	dsl    pb.DSLServiceServer
	suffix string
}

func New(ctx context.Context, opts ...Option) (*Server, error) {
	s := &Server{}

	for _, opt := range opts {
		opt(s)
	}

	server := grpc.NewServer()

	db, err := db.New()
	if err != nil {
		return nil, err
	}

	dsl := new(ctx, &Config{
		DB:     db,
		Suffix: s.suffix,
	})

	pb.RegisterDSLServiceServer(server, dsl)
	s.Server = server
	s.dsl = dsl

	return s, nil
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	return s.Server.Serve(lis)
}

type Option func(*Server)

func WithSuffix(suffix string) Option {
	return func(s *Server) {
		s.suffix = suffix
	}
}
