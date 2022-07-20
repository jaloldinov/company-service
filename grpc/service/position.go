package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type positionService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedPositionServiceServer
}

func NewPositionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *positionService {
	return &positionService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *positionService) Create(ctx context.Context, req *pb.CreatePositionRequest) (*pb.PositionId, error) {
	id, err := s.strg.Position().Create(ctx, req)
	if err != nil {
		s.log.Error("CreatePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.PositionId{
		Id: id,
	}, nil
}

func (s *positionService) GetAll(ctx context.Context, req *pb.GetAllPositionRequest) (*pb.GetAllPositionResponse, error) {
	resp, err := s.strg.Position().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllPosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *positionService) Get(ctx context.Context, req *pb.PositionId) (*pb.GetPositionResponse, error) {
	position, err := s.strg.Position().Get(ctx, req.Id)

	if err != nil {
		s.log.Error("GetPositionID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return position, nil
}

func (s *positionService) Update(ctx context.Context, req *pb.GetPositionResponse) (*pb.PositionStatus, error) {
	status, err := s.strg.Position().Update(ctx, req)
	if err != nil {
		s.log.Error("UpdatePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.PositionStatus{
		Status: status,
	}, nil
}

func (s *positionService) Delete(ctx context.Context, req *pb.PositionId) (*pb.PositionStatus, error) {
	status, err := s.strg.Position().Delete(ctx, req.Id)
	if err != nil {
		s.log.Error("DeletePosition", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.PositionStatus{
		Status: status.String(),
	}, nil
}
