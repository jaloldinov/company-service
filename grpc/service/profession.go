package service

import (
	"context"
	"position_service/config"
	pb "position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type professionService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedProfessionServiceServer
}

func NewProfessionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *professionService {
	return &professionService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *professionService) Create(ctx context.Context, req *pb.CreateProfessionRequest) (*pb.Profession, error) {
	id, err := s.strg.Profession().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Profession{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (s *professionService) GetAll(ctx context.Context, req *pb.GetAllProfessionRequest) (*pb.GetAllProfessionResponse, error) {
	resp, err := s.strg.Profession().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *professionService) Get(ctx context.Context, req *pb.ProfessionId) (*pb.Profession, error) {
	resp, err := s.strg.Profession().Get(req.Id)

	if err != nil {
		s.log.Error("GetProfessionID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *professionService) Update(ctx context.Context, req *pb.Profession) (*pb.Result, error) {
	result, err := s.strg.Profession().Update(req)
	if err != nil {
		s.log.Error("UpdateProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Result{
		Result:  result,
		Message: "Profession updated",
	}, nil
}

func (s *professionService) Delete(ctx context.Context, req *pb.ProfessionId) (*pb.Result, error) {

	result, err := s.strg.Profession().Delete(req.Id)
	if err != nil {
		s.log.Error("DeleteProfession", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Result{
		Result: result.String(),
	}, nil
}
