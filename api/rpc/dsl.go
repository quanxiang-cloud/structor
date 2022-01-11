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

func (s *service) Find(ctx context.Context, req *pb.FindReq) (*pb.FindResp, error) {
	dsl, err := anyToDSL(req.GetDsl())
	if err != nil {
		return &pb.FindResp{}, err
	}

	result, err := s.dsl.Find(ctx, &dslservice.FindReq{
		TableName: req.TableName,
		DSL:       dsl,
		Page:      req.Page,
		Size:      req.Size,
		Sort:      req.Sort,
	})
	if err != nil {
		return &pb.FindResp{}, err
	}

	out, err := structpb.NewValue(result.Data)
	if err != nil {
		return &pb.FindResp{}, err
	}

	resp := &pb.FindResp{}
	resp.Data = &anypb.Any{}
	err = resp.Data.MarshalFrom(out)
	return resp, err
}

func (s *service) Count(ctx context.Context, req *pb.CountReq) (*pb.CountResp, error) {
	dsl, err := anyToDSL(req.GetDsl())
	if err != nil {
		return &pb.CountResp{}, err
	}

	result, err := s.dsl.Count(ctx, &dslservice.CountReq{
		TableName: req.TableName,
		DSL:       dsl,
	})
	if err != nil {
		return &pb.CountResp{}, err
	}

	resp := &pb.CountResp{}
	resp.Data = result.Data
	return resp, nil
}

func (s *service) Insert(ctx context.Context, req *pb.InsertReq) (*pb.InsertResp, error) {
	entirties, err := anyToEntirties(req.GetEntirties())
	if err != nil {
		return &pb.InsertResp{}, err
	}
	_, err = s.dsl.Insert(ctx, &dslservice.InsertReq{
		TableName: req.TableName,
		Entirties: entirties,
	})
	if err != nil {
		return nil, err
	}

	return &pb.InsertResp{}, nil
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

func anyToEntirties(any *anypb.Any) ([]interface{}, error) {
	out := structpb.NewNullValue()
	err := any.UnmarshalTo(out)
	if err != nil {
		return nil, err
	}

	body, err := out.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var value []interface{}
	err = json.Unmarshal(body, &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
