package repository

import (
	"context"
	"database/sql"
	"patient-service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
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

type PatientREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewPatientSqlc(db *sql.DB, log *logrus.Logger) *PatientREPO {
	return &PatientREPO{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *PatientREPO) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error) {

	q.log.Infof("CreatePatient called with FullName=%s, Email=%s", req.FullName, req.Email)

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		q.log.Errorf("Invalid BirthDate format: %v", err)
		return nil, err
	}

	resp, err := q.queries.CreatePatient(ctx, storage.CreatePatientParams{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Gender:      storage.GenderType(req.Gender.String()),
		BirthDate:   birthDate,
		BloodGroup:  storage.BloodType(req.BloodGroup.String()),
	})
	if err != nil {
		q.log.Errorf("CreatePatient DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Patient created successfully with ID=%s", resp.ID.String())
	gender, _ := pb.GenderType_value[string(resp.Gender)]
	blood_group, _ := pb.BloodType_value[string(resp.BloodGroup)]

	return &pb.CreatePatientResponse{
		Patient: &pb.Patient{
			Id:          resp.ID.String(),
			FullName:    resp.FullName,
			Email:       resp.Email,
			Password:    resp.Password,
			Address:     resp.Address,
			PhoneNumber: resp.PhoneNumber,
			Gender:      pb.GenderType(gender),
			BirthDate:   resp.BirthDate.Format("2006-01-02"),
			BloodGroup:  pb.BloodType(blood_group),
			Timestamps: &pb.Timestamps4{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PatientREPO) GetPatientById(ctx context.Context, req *pb.GetPatientByIdRequest) (*pb.GetPatientByIdResponse, error) {

	q.log.Infof("GetPatientById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Patient ID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetPatientById(ctx, id)
	if err != nil {
		q.log.Errorf("GetPatientById DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Patient retrieved successfully with ID=%s", resp.ID.String())
	gender, _ := pb.GenderType_value[string(resp.Gender)]
	blood_group, _ := pb.BloodType_value[string(resp.BloodGroup)]

	return &pb.GetPatientByIdResponse{
		Patient: &pb.Patient{
			Id:          resp.ID.String(),
			FullName:    resp.FullName,
			Email:       resp.Email,
			Password:    resp.Password,
			Address:     resp.Address,
			PhoneNumber: resp.PhoneNumber,
			Gender:      pb.GenderType(gender),
			BirthDate:   resp.BirthDate.Format("2006-01-02"),
			BloodGroup:  pb.BloodType(blood_group),
			Timestamps: &pb.Timestamps4{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PatientREPO) ListPatients(ctx context.Context, req *pb.ListPatientsRequest) (*pb.ListPatientsResponse, error) {

	q.log.Infof("ListPatients called with Search=%s, Limit=%d, Page=%d", req.Search, req.Limit, req.Page)

	params := storage.ListPatientsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListPatients(ctx, params)
	if err != nil {
		q.log.Errorf("ListPatients DB error: %v", err)
		return nil, err
	}

	var patients []*pb.Patient
	var totalCount int64

	for _, r := range resp {
		gender, _ := pb.GenderType_value[string(r.Gender)]
		blood_group, _ := pb.BloodType_value[string(r.BloodGroup)]

		patients = append(patients, &pb.Patient{
			Id:          r.ID.String(),
			FullName:    r.FullName,
			Email:       r.Email,
			Password:    r.Password,
			Address:     r.Address,
			PhoneNumber: r.PhoneNumber,
			Gender:      pb.GenderType(gender),
			BirthDate:   r.BirthDate.Format("2006-01-02"),
			BloodGroup:  pb.BloodType(blood_group),
			Timestamps: &pb.Timestamps4{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListPatients returned %d patients", len(patients))
	return &pb.ListPatientsResponse{
		Patients:   patients,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *PatientREPO) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error) {

	q.log.Infof("UpdatePatient called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Patient ID: %v", err)
		return nil, err
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		q.log.Errorf("Invalid BirthDate format: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdatePatient(ctx, storage.UpdatePatientParams{
		ID:          id,
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Gender:      storage.GenderType(req.Gender.String()),
		BirthDate:   birthDate,
		BloodGroup:  storage.BloodType(req.BloodGroup.String()),
	})
	if err != nil {
		q.log.Errorf("UpdatePatient DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Patient updated successfully with ID=%s", resp.ID.String())
	gender, _ := pb.GenderType_value[string(resp.Gender)]
	blood_group, _ := pb.BloodType_value[string(resp.BloodGroup)]

	return &pb.UpdatePatientResponse{
		Patient: &pb.Patient{
			Id:          resp.ID.String(),
			FullName:    resp.FullName,
			Email:       resp.Email,
			Password:    resp.Password,
			Address:     resp.Address,
			PhoneNumber: resp.PhoneNumber,
			Gender:      pb.GenderType(gender),
			BirthDate:   resp.BirthDate.Format("2006-01-02"),
			BloodGroup:  pb.BloodType(blood_group),
			Timestamps: &pb.Timestamps4{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PatientREPO) DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.DeletePatientResponse, error) {

	q.log.Infof("DeletePatient called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Patient ID: %v", err)
		return nil, err
	}

	err = q.queries.DeletePatient(ctx, storage.DeletePatientParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeletePatient DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Patient deleted successfully with ID=%s", req.Id)
	return &pb.DeletePatientResponse{
		Status: 204,
	}, nil
}
