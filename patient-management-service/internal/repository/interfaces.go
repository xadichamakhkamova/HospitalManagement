package repository

import (
	"patient-service/internal/storage"
	"context"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
)

func NewIPatientRepository(queries *storage.Queries) IPatientRepository {
	return &PatientREPO{
		queries: queries,
	}
}

type IPatientRepository interface {
	//Patient
	CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error)
	GetPatientById(ctx context.Context, req *pb.GetPatientByIdRequest) (*pb.GetPatientByIdResponse, error)
	ListPatients(ctx context.Context, req *pb.ListPatientsRequest) (*pb.ListPatientsResponse, error)
	UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error)
	DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.DeletePatientResponse, error)
}