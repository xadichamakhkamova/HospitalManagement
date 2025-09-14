package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
)

func NewIDoctorRepository(repo *DoctorREPO, log *logrus.Logger) IDoctorRepository {
	return &DoctorREPO{
		queries: repo.queries,
		log:     log,
	}
}

type IDoctorRepository interface {
	// Appointment
	CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error)
	GetAppointmentById(ctx context.Context, req *pb.GetAppointmentByIdRequest) (*pb.GetAppointmentByIdResponse, error)
	ListAppointments(ctx context.Context, req *pb.ListAppointmentsRequest) (*pb.ListAppointmentsResponse, error)
	UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentRequest) (*pb.UpdateAppointmentResponse, error)
	DeleteAppointment(ctx context.Context, req *pb.DeleteAppointmentRequest) (*pb.DeleteAppointmentResponse, error)

	// Prescription
	CreatePrescription(ctx context.Context, req *pb.CreatePrescriptionRequest) (*pb.CreatePrescriptionResponse, error)
	GetPrescriptionById(ctx context.Context, req *pb.GetPrescriptionByIdRequest) (*pb.GetPrescriptionByIdResponse, error)
	ListPrescriptions(ctx context.Context, req *pb.ListPrescriptionsRequest) (*pb.ListPrescriptionsResponse, error)
	UpdatePrescription(ctx context.Context, req *pb.UpdatePrescriptionRequest) (*pb.UpdatePrescriptionResponse, error)
	DeletePrescription(ctx context.Context, req *pb.DeletePrescriptionRequest) (*pb.DeletePrescriptionResponse, error)
}
