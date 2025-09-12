package service

import (
	"context"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"
	"nurse-service/internal/repository"
)

type NurseService struct {
	pb.UnimplementedNurseServiceServer
	repo repository.INurseRepository
}

func NewNurseService(repo repository.INurseRepository) *NurseService {
	return &NurseService{
		repo: repo,
	}
}

// ------------------- Nurse -------------------

func (s *NurseService) CreateDonor(ctx context.Context, req *pb.CreateDonorRequest) (*pb.CreateDonorResponse, error) {
	return s.repo.CreateDonor(ctx, req)
}

func (s *NurseService) GetNurseById(ctx context.Context, req *pb.GetDonorByIdRequest) (*pb.GetDonorByIdResponse, error) {
	return s.repo.GetDonorById(ctx, req)
}

func (s *NurseService) ListDonors(ctx context.Context, req *pb.ListDonorsRequest) (*pb.ListDonorsResponse, error) {
	return s.repo.ListDonors(ctx, req)
}

func (s *NurseService) UpdateDonor(ctx context.Context, req *pb.UpdateDonorRequest) (*pb.UpdateDonorResponse, error) {
	return s.repo.UpdateDonor(ctx, req)
}

func (s *NurseService) DeleteDonor(ctx context.Context, req *pb.DeleteDonorRequest) (*pb.DeleteDonorResponse, error) {
	return s.repo.DeleteDonor(ctx, req)
}
