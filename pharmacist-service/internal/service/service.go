package service

import (
	"context"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	"pharmacist-service/internal/repository"
)

type PharmaService struct {
	pb.UnimplementedPharmacistServiceServer
	repo repository.IPharmacistRepository
}

func NewPharmaService(repo repository.IPharmacistRepository) *PharmaService {
	return &PharmaService{
		repo: repo,
	}
}

// ------------------- Medicine -------------------

func (s *PharmaService) CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error) {
	return s.repo.CreateMedicine(ctx, req)
}

func (s *PharmaService) GetMedicineById(ctx context.Context, req *pb.GetMedicineByIdRequest) (*pb.GetMedicineByIdResponse, error) {
	return s.repo.GetMedicineById(ctx, req)
}

func (s *PharmaService) ListMedicines(ctx context.Context, req *pb.ListMedicinesRequest) (*pb.ListMedicinesResponse, error) {
	return s.repo.ListMedicines(ctx, req)
}

func (s *PharmaService) UpdateMedicine(ctx context.Context, req *pb.UpdateMedicineRequest) (*pb.UpdateMedicineResponse, error) {
	return s.repo.UpdateMedicine(ctx, req)
}

func (s *PharmaService) DeleteMedicine(ctx context.Context, req *pb.DeleteMedicineRequest) (*pb.DeleteMedicineResponse, error) {
	return s.repo.DeleteMedicine(ctx, req)
}
