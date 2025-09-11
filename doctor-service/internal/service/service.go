package service

import (
	"context"

	adminpb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
	patientpb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"

	"doctor-service/internal/repository"
)

type DoctorService struct {
	pb.UnimplementedDoctorServiceServer
	repo          repository.IDoctorRepository
	adminClient   adminpb.AdminServiceClient
	patientClient patientpb.PatientManagementServiceClient
}

func NewDoctorService(
	repo repository.IDoctorRepository,
	adminClient *adminpb.AdminServiceClient,
	patientClient *patientpb.PatientManagementServiceClient,
) *DoctorService {
	return &DoctorService{
		repo:          repo,
		adminClient:   *adminClient,
		patientClient: *patientClient,
	}
}

// ------------------- Appointment -------------------

func (s *DoctorService) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {
	// validate doctor exists in Admin service
	_, err := s.adminClient.GetDoctorById(ctx, &adminpb.GetPersonalByIdRequest{Id: req.DoctorId})
	if err != nil {
		return nil, err
	}

	// validate patient exists in Patient service
	_, err = s.patientClient.GetPatientById(ctx, &patientpb.GetPatientByIdRequest{Id: req.PatientId})
	if err != nil {
		return nil, err
	}

	return s.repo.CreateAppointment(ctx, req)
}

func (s *DoctorService) GetAppointmentById(ctx context.Context, req *pb.GetAppointmentByIdRequest) (*pb.GetAppointmentByIdResponse, error) {
	return s.repo.GetAppointmentById(ctx, req)
}

func (s *DoctorService) ListAppointments(ctx context.Context, req *pb.ListAppointmentsRequest) (*pb.ListAppointmentsResponse, error) {
	return s.repo.ListAppointments(ctx, req)
}

func (s *DoctorService) UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentRequest) (*pb.UpdateAppointmentResponse, error) {
	return s.repo.UpdateAppointment(ctx, req)
}

func (s *DoctorService) DeleteAppointment(ctx context.Context, req *pb.DeleteAppointmentRequest) (*pb.DeleteAppointmentResponse, error) {
	return s.repo.DeleteAppointment(ctx, req)
}

// ------------------- Prescription -------------------

func (s *DoctorService) CreatePrescription(ctx context.Context, req *pb.CreatePrescriptionRequest) (*pb.CreatePrescriptionResponse, error) {
	// validate doctor exists
	_, err := s.adminClient.GetDoctorById(ctx, &adminpb.GetPersonalByIdRequest{Id: req.DoctorId})
	if err != nil {
		return nil, err
	}

	// validate patient exists
	_, err = s.patientClient.GetPatientById(ctx, &patientpb.GetPatientByIdRequest{Id: req.PatientId})
	if err != nil {
		return nil, err
	}

	return s.repo.CreatePrescription(ctx, req)
}

func (s *DoctorService) GetPrescriptionById(ctx context.Context, req *pb.GetPrescriptionByIdRequest) (*pb.GetPrescriptionByIdResponse, error) {
	return s.repo.GetPrescriptionById(ctx, req)
}

func (s *DoctorService) ListPrescriptions(ctx context.Context, req *pb.ListPrescriptionsRequest) (*pb.ListPrescriptionsResponse, error) {
	return s.repo.ListPrescriptions(ctx, req)
}

func (s *DoctorService) UpdatePrescription(ctx context.Context, req *pb.UpdatePrescriptionRequest) (*pb.UpdatePrescriptionResponse, error) {
	return s.repo.UpdatePrescription(ctx, req)
}

func (s *DoctorService) DeletePrescription(ctx context.Context, req *pb.DeletePrescriptionRequest) (*pb.DeletePrescriptionResponse, error) {
	return s.repo.DeletePrescription(ctx, req)
}
