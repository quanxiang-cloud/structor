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

func (d *ddlService) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	resp, err := d.ddl.Create(ctx, &ddlservice.ExecuteReq{
		Table:  req.TableName,
		Fields: transform(req.Fields),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResp{
		TableName: resp.Table,
	}, nil
}

func (d *ddlService) Add(ctx context.Context, req *pb.AddReq) (*pb.AddResp, error) {
	resp, err := d.ddl.Add(ctx, &ddlservice.ExecuteReq{
		Table:  req.TableName,
		Fields: transform(req.Fields),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddResp{
		TableName: resp.Table,
	}, nil
}

func (d *ddlService) Modify(ctx context.Context, req *pb.ModifyReq) (*pb.ModifyResp, error) {
	resp, err := d.ddl.Modify(ctx, &ddlservice.ExecuteReq{
		Table:  req.TableName,
		Fields: transform(req.Fields),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ModifyResp{
		TableName: resp.Table,
	}, nil
}

func (d *ddlService) Indexes(ctx context.Context, req *pb.IndexesReq) (*pb.IndexesResp, error) {
	// resp, err := d.ddl.Index(ctx, &ddlservice.IndexReq{
	// 	Option: req.Option,
	// 	Table:  req.TableName,
	// 	Fields: transformIndex(req.Titles),
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &pb.IndexesResp{
	// 	IndexName: resp.Index,
	// }, nil
	return nil, nil
}

func (d *ddlService) DropIndexes(ctx context.Context, req *pb.DropIndexesReq) (*pb.DropIndexesResp, error) {
	// _, err := d.ddl.DropIndexes(ctx, &ddlservice.DropIndexesReq{
	// 	Option: structor.DropIndexesOpt,
	// 	Table:  req.TableName,
	// 	Fields: transformIndex(req.IndexName),
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.DropIndexesResp{}, nil
}

func transform(fields []*pb.Field) []*structor.Field {
	ret := make([]*structor.Field, 0, len(fields))
	for _, f := range fields {
		ret = append(ret, &structor.Field{
			Title:   f.Title,
			Type:    f.Type,
			Max:     f.Max,
			Comment: f.Comment,
			NotNull: f.NotNull,
		})
	}
	return ret
}

func transformIndex(titles []string) []*structor.Field {
	ret := make([]*structor.Field, 0, len(titles))
	for _, title := range titles {
		ret = append(ret, &structor.Field{
			Title: title,
		})
	}
	return ret
}
