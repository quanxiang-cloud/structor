package rpc

import (
	"context"

	pb "github.com/quanxiang-cloud/structor/api/proto"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
	"google.golang.org/grpc"
)

func New(ctx context.Context) (*grpc.Server, error) {
	server := grpc.NewServer()

	db, err := db.New()
	if err != nil {
		return nil, err
	}

	dsl := new(ctx, &Config{
		DB: db,
	})
	ddl := newDDL(ctx, &Config{
		DB: db,
	})

	pb.RegisterDSLServiceServer(server, dsl)
	pb.RegisterDDLServiceServer(server, ddl)

	return server, nil
}
