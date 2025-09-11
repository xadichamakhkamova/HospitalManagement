package service

import (
	"context"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
	"admin-panel/internal/repository"
)

type AdminService struct {
	pb.UnimplementedAdminServiceServer
	repo repository.IAdminRepository
}

func NewAdminService(repo repository.IAdminRepository) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

// ------------------- Department -------------------

func (s *AdminService) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error) {
	return s.repo.CreateDepartment(ctx, req)
}

func (s *AdminService) GetDepartmentById(ctx context.Context, req *pb.GetDepartmentByIdRequest) (*pb.GetDepartmentByIdResponse, error) {
	return s.repo.GetDepartmentById(ctx, req)
}

func (s *AdminService) ListDeparments(ctx context.Context, req *pb.ListDepartmentsRequest) (*pb.ListDepartmentsResponse, error) {
	return s.repo.ListDeparments(ctx, req)
}

func (s *AdminService) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error) {
	return s.repo.UpdateDepartment(ctx, req)
}

func (s *AdminService) DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error) {
	return s.repo.DeleteDepartment(ctx, req)
}

// ------------------- Personal -------------------

func (s *AdminService) CreatePersonal(ctx context.Context, req *pb.CreatePersonalRequest) (*pb.CreatePersonalResponse, error) {
	return s.repo.CreatePersonal(ctx, req)
}

func (s *AdminService) GetPersonalById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetPersonalByIdResponse, error) {
	return s.repo.GetPersonalById(ctx, req)
}

func (s *AdminService) ListPersonals(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListPersonalsResponse, error) {
	return s.repo.ListPersonals(ctx, req)
}

func (s *AdminService) UpdatePersonal(ctx context.Context, req *pb.UpdatePersonalRequest) (*pb.UpdatePersonalResponse, error) {
	return s.repo.UpdatePersonal(ctx, req)
}

func (s *AdminService) DeletePersonal(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error) {
	return s.repo.DeletePersonal(ctx, req)
}

// ------------------- Doctor -------------------

func (s *AdminService) CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*pb.CreateDoctorResponse, error) {
	return s.repo.CreateDoctor(ctx, req)
}

func (s *AdminService) GetDoctorById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetDoctorByIdResponse, error) {
	return s.repo.GetDoctorById(ctx, req)
}

func (s *AdminService) ListDoctors(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListDoctorsResponse, error) {
	return s.repo.ListDoctors(ctx, req)
}

func (s *AdminService) UpdateDoctor(ctx context.Context, req *pb.UpdateDoctorRequest) (*pb.UpdateDoctorResponse, error) {
	return s.repo.UpdateDoctor(ctx, req)
}

func (s *AdminService) DeleteDoctor(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error) {
	return s.repo.DeleteDoctor(ctx, req)
}

// ------------------- Bed -------------------

func (s *AdminService) CreateBed(ctx context.Context, req *pb.CreateBedRequest) (*pb.CreateBedResponse, error) {
	return s.repo.CreateBed(ctx, req)
}

func (s *AdminService) GetBedByID(ctx context.Context, req *pb.GetBedByIDRequest) (*pb.GetBedByIDResponse, error) {
	return s.repo.GetBedByID(ctx, req)
}

func (s *AdminService) ListBedS(ctx context.Context, req *pb.ListBedSRequest) (*pb.ListBedSResponse, error) {
	return s.repo.ListBedS(ctx, req)
}

func (s *AdminService) UpdateBed(ctx context.Context, req *pb.UpdateBedRequest) (*pb.UpdateBedResponse, error) {
	return s.repo.UpdateBed(ctx, req)
}

func (s *AdminService) DeleteBed(ctx context.Context, req *pb.DeleteBedRequest) (*pb.DeleteBedResponse, error) {
	return s.repo.DeleteBed(ctx, req)
}
