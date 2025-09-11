package service

import (
	"context"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
	"patient-service/internal/repository"
)

type PatientService struct {
	pb.UnimplementedPatientManagementServiceServer
	repo repository.IPatientRepository
}

func NewPatientService(repo repository.IPatientRepository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

// ------------------- Patient -------------------

func (s *PatientService) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error) {
	return s.repo.CreatePatient(ctx, req)
}

func (s *PatientService) GetPatientById(ctx context.Context, req *pb.GetPatientByIdRequest) (*pb.GetPatientByIdResponse, error) {
	return s.repo.GetPatientById(ctx, req)
}

func (s *PatientService) ListDeparments(ctx context.Context, req *pb.ListPatientsRequest) (*pb.ListPatientsResponse, error) {
	return s.repo.ListPatients(ctx, req)
}

func (s *PatientService) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error) {
	return s.repo.UpdatePatient(ctx, req)
}

func (s *PatientService) DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.DeletePatientResponse, error) {
	return s.repo.DeletePatient(ctx, req)
}
