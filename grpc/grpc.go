package grpc

import (
	"position_service/config"
	"position_service/genproto/company_service"
	"position_service/genproto/position_service"
	"position_service/grpc/service"
	"position_service/pkg/logger"
	"position_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	position_service.RegisterProfessionServiceServer(grpcServer, service.NewProfessionService(cfg, log, strg))
	position_service.RegisterAttributeServiceServer(grpcServer, service.NewAttributeService(cfg, log, strg))
	company_service.RegisterCompanyServiceServer(grpcServer, service.NewCompanyService(cfg, log, strg))
	position_service.RegisterPositionServiceServer(grpcServer, service.NewPositionService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
