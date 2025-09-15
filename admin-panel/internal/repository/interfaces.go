package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
)

func NewIAdminRepository(repo *AdminREPO, log *logrus.Logger) IAdminRepository {
	return &AdminREPO{
		queries: repo.queries,
		log:     log,
	}
}

type IAdminRepository interface {
	//Department
	CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error)
	GetDepartmentById(ctx context.Context, req *pb.GetDepartmentByIdRequest) (*pb.GetDepartmentByIdResponse, error)
	ListDeparments(ctx context.Context, req *pb.ListDepartmentsRequest) (*pb.ListDepartmentsResponse, error)
	UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error)
	DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error)

	//Personal
	CreatePersonal(ctx context.Context, req *pb.CreatePersonalRequest) (*pb.CreatePersonalResponse, error)
	GetPersonalById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetPersonalByIdResponse, error)
	ListPersonals(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListPersonalsResponse, error)
	UpdatePersonal(ctx context.Context, req *pb.UpdatePersonalRequest) (*pb.UpdatePersonalResponse, error)
	DeletePersonal(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error)

	//Doctor
	CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*pb.CreateDoctorResponse, error)
	GetDoctorById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetDoctorByIdResponse, error)
	ListDoctors(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListDoctorsResponse, error)
	UpdateDoctor(ctx context.Context, req *pb.UpdateDoctorRequest) (*pb.UpdateDoctorResponse, error)
	DeleteDoctor(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error)
}
