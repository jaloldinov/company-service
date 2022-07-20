package storage

import (
	"context"
	"errors"
	cb "position_service/genproto/company_service"
	pb "position_service/genproto/position_service"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	Profession() ProfessionI
	Attribute() AttributeI
	Company() CompanyI
	Position() PositionI
}

type ProfessionI interface {
	Create(ctx context.Context, entity *pb.CreateProfessionRequest) (id string, err error)
	GetAll(ctx context.Context, req *pb.GetAllProfessionRequest) (*pb.GetAllProfessionResponse, error)
	Get(id string) (*pb.Profession, error)
	Update(req *pb.Profession) (string, error)
	Delete(id string) (*pb.Result, error)
}

type AttributeI interface {
	Create(ctx context.Context, entity *pb.CreateAttributeRequest) (id string, err error)
	GetAll(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error)
	Get(id string) (*pb.Attribute, error)
	Update(req *pb.Attribute) (string, error)
	Delete(id string) (*pb.AttributeResult, error)
}

type CompanyI interface {
	Create(ctx context.Context, entity *cb.CreateCompany) (id string, err error)
	GetAll(ctx context.Context, req *cb.GetAllCompanyRequest) (*cb.GetAllCompanyResponse, error)
	Get(ctx context.Context, id string) (*cb.Company, error)
	Update(ctx context.Context, req *cb.Company) (string, error)
	Delete(ctx context.Context, id string) (*cb.CompanyResult, error)
}

type PositionI interface {
	Create(ctx context.Context, req *pb.CreatePositionRequest) (id string, err error)
	GetAll(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error)
	Get(ctx context.Context, id string) (*pb.GetPositionResponse, error)
	Update(ctx context.Context, req *pb.GetPositionResponse) (string, error)
	Delete(ctx context.Context, id string) (*pb.PositionStatus, error)
}
