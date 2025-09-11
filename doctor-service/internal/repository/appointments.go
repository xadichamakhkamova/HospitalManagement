package repository

import (
	"context"
	"database/sql"
	"doctor-service/internal/storage"
	"time"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
)

// NullTime → *timestamppb.Timestamp converter
func convertNullTime(nt sql.NullTime) *timestamppb.Timestamp {
	if nt.Valid {
		return timestamppb.New(nt.Time)
	}
	return nil
}

type DoctorREPO struct {
	queries *storage.Queries
}

func NewDoctorSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *DoctorREPO) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {

	doctor_id, err := uuid.Parse(req.DoctorId)
	if err != nil {
		return nil, err
	}
	patient_id, err := uuid.Parse(req.PatientId)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.CreateAppointment(ctx, storage.CreateAppointmentParams{
		DoctorID:        doctor_id,
		PatientID:       patient_id,
		AppointmentDate: req.Date.AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAppointmentResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      convertNullTime(sql.NullTime{Time: resp.AppointmentDate, Valid: true}),
			Timestamps: &pb.Timestamps2{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *DoctorREPO) GetAppointmentById(ctx context.Context, req *pb.GetAppointmentByIdRequest) (*pb.GetAppointmentByIdResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetAppointmentById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAppointmentByIdResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      convertNullTime(sql.NullTime{Time: resp.AppointmentDate, Valid: true}),
			Timestamps: &pb.Timestamps2{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *DoctorREPO) ListAppointments(ctx context.Context, req *pb.ListAppointmentsRequest) (*pb.ListAppointmentsResponse, error) {

	params := storage.ListAppointmentsParams{
		Column1: req.Date.AsTime(),
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListAppointments(ctx, params)
	if err != nil {
		return nil, err
	}

	var appointments []*pb.Appointment
	var totalCount int64

	for _, r := range resp {
		appointments = append(appointments, &pb.Appointment{
			Id:        r.ID.String(),
			DoctorId:  r.DoctorID.String(),
			PatientId: r.PatientID.String(),
			Date:      req.Date,
			Timestamps: &pb.Timestamps2{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = r.TotalCount
	}

	return &pb.ListAppointmentsResponse{
		Appointment: appointments,
		TotalCount:  int32(totalCount),
	}, nil
}

func (q *DoctorREPO) UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentRequest) (*pb.UpdateAppointmentResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	doctor_id, err := uuid.Parse(req.DoctorId)
	if err != nil {
		return nil, err
	}
	patient_id, err := uuid.Parse(req.PatientId)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdateAppointment(ctx, storage.UpdateAppointmentParams{
		ID:              id,
		DoctorID:        doctor_id,
		PatientID:       patient_id,
		AppointmentDate: req.Date.AsTime(),
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAppointmentResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      convertNullTime(sql.NullTime{Time: resp.AppointmentDate, Valid: true}),
			Timestamps: &pb.Timestamps2{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *DoctorREPO) DeleteAppointment(ctx context.Context, req *pb.DeleteAppointmentRequest) (*pb.DeleteAppointmentResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeleteAppointment(ctx, storage.DeleteAppointmentParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}
	return &pb.DeleteAppointmentResponse{
		Status: 204,
	}, nil
}
