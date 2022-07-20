package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type attributeService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedAttributeServiceServer
}

func NewAttributeService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *attributeService {
	return &attributeService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *attributeService) Create(ctx context.Context, req *pb.CreateAttributeRequest) (*pb.Attribute, error) {
	id, err := s.strg.Attribute().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateAttribute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Attribute{
		Id:            id,
		Name:          req.Name,
		AttributeType: req.AttributeType,
	}, nil
}

func (s *attributeService) GetAll(ctx context.Context, req *pb.GetAllAttributeRequest) (*pb.GetAllAttributeResponse, error) {
	resp, err := s.strg.Attribute().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllAttribute", logger.Any("req", req), logger.Error(err))
		return resp, nil
	}

	return resp, nil
}

func (s *attributeService) Get(ctx context.Context, req *pb.AttributeId) (*pb.Attribute, error) {
	resp, err := s.strg.Attribute().Get(req.Id)

	if err != nil {
		s.log.Error("GetAttributeID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *attributeService) Update(ctx context.Context, req *pb.Attribute) (*pb.AttributeResult, error) {
	result, err := s.strg.Attribute().Update(req)
	if err != nil {
		s.log.Error("UpdateAttrubute", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.AttributeResult{
		Status: result,
	}, nil
}

func (s *attributeService) Delete(ctx context.Context, req *pb.AttributeId) (*pb.AttributeResult, error) {

	result, err := s.strg.Attribute().Delete(req.Id)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.AttributeResult{
		Status: result.String(),
	}, nil
}
