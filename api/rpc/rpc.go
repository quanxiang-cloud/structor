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

func New(ctx context.Context) (*Server, error) {
	server := grpc.NewServer()

	db, err := db.New()
	if err != nil {
		return nil, err
	}

	dsl := new(ctx, &Config{
		DB: db,
	})

	pb.RegisterDSLServiceServer(server, dsl)

	return &Server{
		Server: server,
		dsl:    dsl,
	}, nil
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	return s.Server.Serve(lis)
}
