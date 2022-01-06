package rpc

import (
	"context"
	"encoding/json"

	pb "github.com/quanxiang-cloud/structor/api/proto"
	"github.com/quanxiang-cloud/structor/internal/dorm/db"
	"github.com/quanxiang-cloud/structor/internal/service/dsl"
	dslservice "github.com/quanxiang-cloud/structor/internal/service/dsl"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

type service struct {
	dsl dslservice.DSLService
}

type Config struct {
	DB *db.Dorm
}

func new(ctx context.Context, conf *Config) pb.DSLServiceServer {
	return &service{
		dsl: dslservice.New(ctx,
			dslservice.WithDB(conf.DB),
		),
	}
}

func (s *service) FindOne(ctx context.Context, req *pb.FindOneReq) (*pb.FindOneResp, error) {
	dsl, err := anyToDSL(req.GetDsl())
	if err != nil {
		return &pb.FindOneResp{}, err
	}

	result, err := s.dsl.FindOne(ctx, &dslservice.FindOneReq{
		TableName: req.TableName,
		DSL:       dsl,
	})

	if err != nil {
		return &pb.FindOneResp{}, err
	}

	out, err := structpb.NewValue(result.Data)
	if err != nil {
		return &pb.FindOneResp{}, err
	}

	resp := &pb.FindOneResp{}
	resp.Data = &anypb.Any{}
	err = resp.Data.MarshalFrom(out)
	return resp, err
}

func anyToDSL(any *anypb.Any) (dslservice.DSL, error) {
	out := structpb.NewNullValue()
	err := any.UnmarshalTo(out)
	if err != nil {
		return dslservice.DSL{}, err
	}

	body, err := out.MarshalJSON()
	if err != nil {
		return dslservice.DSL{}, err
	}

	dsl := dsl.DSL{}
	err = json.Unmarshal(body, &dsl)
	if err != nil {
		return dslservice.DSL{}, err
	}

	return dsl, nil
}
