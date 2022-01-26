package rpc

import (
	"context"

	pb "github.com/quanxiang-cloud/structor/api/proto"
	"github.com/quanxiang-cloud/structor/internal/dorm/structor"
	ddlservice "github.com/quanxiang-cloud/structor/internal/service/ddl"
)

type ddlService struct {
	ddl ddlservice.DDLService
}

func newDDL(ctx context.Context, conf *Config) pb.DDLServiceServer {
	return &ddlService{
		ddl: ddlservice.NewDDL(ctx,
			ddlservice.WithDB(conf.DB)),
	}
}

func (d *ddlService) Execute(ctx context.Context, req *pb.ExecuteReq) (*pb.ExecuteResp, error) {
	resp, err := d.ddl.Execute(ctx, &ddlservice.ExecuteReq{
		Table:  req.TableName,
		Option: req.Option,
		Fields: transform(req.Fields),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ExecuteResp{
		TableName: resp.Table,
	}, nil
}

func transform(fields []*pb.Field) []*structor.Field {
	ret := make([]*structor.Field, 0, len(fields))
	for _, f := range fields {
		ret = append(ret, &structor.Field{
			Title:   f.Title,
			Type:    f.Type,
			Comment: f.Comment,
			NotNull: f.NotNull,
		})
	}
	return ret
}
