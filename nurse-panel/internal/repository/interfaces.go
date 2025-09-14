package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"
)

func NewINurseRepository(repo *NurseREPO, log *logrus.Logger) INurseRepository {
	return &NurseREPO{
		queries: repo.queries,
		log:     log,
	}
}

type INurseRepository interface {
	// Donor
	CreateDonor(ctx context.Context, req *pb.CreateDonorRequest) (*pb.CreateDonorResponse, error)
	GetDonorById(ctx context.Context, req *pb.GetDonorByIdRequest) (*pb.GetDonorByIdResponse, error)
	ListDonors(ctx context.Context, req *pb.ListDonorsRequest) (*pb.ListDonorsResponse, error)
	UpdateDonor(ctx context.Context, req *pb.UpdateDonorRequest) (*pb.UpdateDonorResponse, error)
	DeleteDonor(ctx context.Context, req *pb.DeleteDonorRequest) (*pb.DeleteDonorResponse, error)
	RegisterDonation(ctx context.Context, req *pb.RegisterDonationRequest) (*pb.RegisterDonationResponse, error)
	RegisterCheckup(ctx context.Context, req *pb.RegisterCheckupRequest) (*pb.RegisterCheckupResponse, error)
}
