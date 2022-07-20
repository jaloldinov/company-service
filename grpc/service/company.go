package service

import (
	"context"
	"position_service/config"
	cb "position_service/genproto/company_service"
	"position_service/pkg/logger"
	"position_service/storage"
)

type companyService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	cb.UnimplementedCompanyServiceServer
}

func NewCompanyService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *companyService {
	return &companyService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *companyService) Create(ctx context.Context, req *cb.CreateCompany) (*cb.Company, error) {
	id, err := s.strg.Company().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateCompany", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &cb.Company{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (s *companyService) GetAll(ctx context.Context, req *cb.GetAllCompanyRequest) (*cb.GetAllCompanyResponse, error) {
	resp, err := s.strg.Company().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllCompany", logger.Any("req", req), logger.Error(err))
		return resp, nil
	}

	return resp, nil
}

func (s *companyService) Get(ctx context.Context, req *cb.CompanyId) (*cb.Company, error) {
	resp, err := s.strg.Company().Get(ctx, req.Id)

	if err != nil {
		s.log.Error("GetCompanyID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *companyService) Update(ctx context.Context, req *cb.Company) (*cb.CompanyResult, error) {
	result, err := s.strg.Company().Update(ctx, req)

	if err != nil {
		s.log.Error("UpdateCompany", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &cb.CompanyResult{
		Result: result,
	}, nil
}

func (s *companyService) Delete(ctx context.Context, req *cb.CompanyId) (*cb.CompanyResult, error) {

	result, err := s.strg.Company().Delete(ctx, req.Id)
	if err != nil {
		s.log.Error("DeleteCompany", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &cb.CompanyResult{
		Result: result.String(),
	}, nil
}
