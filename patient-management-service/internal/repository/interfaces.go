package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
)

func NewIPatientRepository(repo *PatientREPO, log *logrus.Logger) IPatientRepository {
	return &PatientREPO{
		queries: repo.queries,
		log:     log,
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
