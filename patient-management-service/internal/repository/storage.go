package repository

import (
	"context"
	"database/sql"
	"patient-service/internal/storage"
	"time"

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
}

func NewPatientSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *PatientREPO) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error) {

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
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
		return nil, err
	}

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

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetPatientById(ctx, id)
	if err != nil {
		return nil, err
	}

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

	params := storage.ListPatientsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListPatients(ctx, params)
	if err != nil {
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

	return &pb.ListPatientsResponse{
		Patients: patients,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *PatientREPO) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdatePatient(ctx, storage.UpdatePatientParams{
		ID: id,
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
		return nil, err
	}

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

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeletePatient(ctx, storage.DeletePatientParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}
	return &pb.DeletePatientResponse{
		Status: 204,
	}, nil
}