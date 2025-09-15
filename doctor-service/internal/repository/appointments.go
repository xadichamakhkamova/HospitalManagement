package repository

import (
	"context"
	"database/sql"
	"doctor-service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"

	"github.com/google/uuid"
)

type DoctorREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewDoctorSqlc(db *sql.DB, log *logrus.Logger) *DoctorREPO {
	return &DoctorREPO{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *DoctorREPO) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {

	q.log.Infof("CreateAppointment called with DoctorId=%s, PatientId=%s, Date=%v", req.DoctorId, req.PatientId, req.Date)

	doctor_id, err := uuid.Parse(req.DoctorId)
	if err != nil {
		q.log.Errorf("Invalid Doctor UUID: %v", err)
		return nil, err
	}
	patient_id, err := uuid.Parse(req.PatientId)
	if err != nil {
		q.log.Errorf("Invalid Patient UUID: %v", err)
		return nil, err
	}

	layout := "2006-01-02 15:04:05" 
	parsedDate, err := time.Parse(layout, req.Date)
	if err != nil {
		q.log.Errorf("Invalid date format: %v", err)
		return nil, err
	}

	resp, err := q.queries.CreateAppointment(ctx, storage.CreateAppointmentParams{
		DoctorID:        doctor_id,
		PatientID:       patient_id,
		AppointmentDate: parsedDate,
	})
	if err != nil {
		q.log.Errorf("CreateAppointment error: %v", err)
		return nil, err
	}

	q.log.Infof("Appointment created successfully with ID=%s", resp.ID.String())
	return &pb.CreateAppointmentResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      resp.AppointmentDate.String(),
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) GetAppointmentById(ctx context.Context, req *pb.GetAppointmentByIdRequest) (*pb.GetAppointmentByIdResponse, error) {

	q.log.Infof("GetAppointmentById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetAppointmentById(ctx, id)
	if err != nil {
		q.log.Errorf("GetAppointmentById error: %v", err)
		return nil, err
	}

	q.log.Infof("Appointment retrieved successfully with ID=%s", resp.ID.String())
	return &pb.GetAppointmentByIdResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      resp.AppointmentDate.String(),
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) ListAppointments(ctx context.Context, req *pb.ListAppointmentsRequest) (*pb.ListAppointmentsResponse, error) {

	q.log.Infof("ListAppointments called with Date=%v, Limit=%d, Page=%d", req.Date, req.Limit, req.Page)

	layout := "2006-01-02 15:04:05" 
	parsedDate, err := time.Parse(layout, req.Date)
	if err != nil {
		q.log.Errorf("Invalid date format: %v", err)
		return nil, err
	}

	params := storage.ListAppointmentsParams{
		Column1: parsedDate,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListAppointments(ctx, params)
	if err != nil {
		q.log.Errorf("ListAppointments error: %v", err)
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
				CreatedAt: r.CreatedAt.Time.String(),
				UpdatedAt: r.UpdatedAt.Time.String(),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListAppointments returned %d appointments", len(appointments))
	return &pb.ListAppointmentsResponse{
		Appointment: appointments,
		TotalCount:  int32(totalCount),
	}, nil
}

func (q *DoctorREPO) UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentRequest) (*pb.UpdateAppointmentResponse, error) {

	q.log.Infof("UpdateAppointment called with ID=%s, DoctorId=%s, PatientId=%s, Date=%v", req.Id, req.DoctorId, req.PatientId, req.Date)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}
	doctor_id, err := uuid.Parse(req.DoctorId)
	if err != nil {
		q.log.Errorf("Invalid Doctor UUID: %v", err)
		return nil, err
	}
	patient_id, err := uuid.Parse(req.PatientId)
	if err != nil {
		q.log.Errorf("Invalid Patient UUID: %v", err)
		return nil, err
	}

	layout := "2006-01-02 15:04:05" 
	parsedDate, err := time.Parse(layout, req.Date)
	if err != nil {
		q.log.Errorf("Invalid date format: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdateAppointment(ctx, storage.UpdateAppointmentParams{
		ID:              id,
		DoctorID:        doctor_id,
		PatientID:       patient_id,
		AppointmentDate: parsedDate,
		UpdatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("UpdateAppointment error: %v", err)
		return nil, err
	}

	q.log.Infof("Appointment updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdateAppointmentResponse{
		Appointment: &pb.Appointment{
			Id:        resp.ID.String(),
			DoctorId:  resp.DoctorID.String(),
			PatientId: resp.PatientID.String(),
			Date:      resp.AppointmentDate.String(),
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) DeleteAppointment(ctx context.Context, req *pb.DeleteAppointmentRequest) (*pb.DeleteAppointmentResponse, error) {

	q.log.Infof("DeleteAppointment called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeleteAppointment(ctx, storage.DeleteAppointmentParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeleteAppointment error: %v", err)
		return nil, err
	}

	q.log.Infof("Appointment deleted successfully with ID=%s", req.Id)
	return &pb.DeleteAppointmentResponse{
		Status: 204,
	}, nil
}
