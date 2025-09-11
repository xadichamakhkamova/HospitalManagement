package repository

import (
	"admin-panel/internal/storage"
	"context"
	"database/sql"

	"time"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/google/uuid"
)

func (q *AdminREPO) CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*pb.CreateDoctorResponse, error) {

	id, err := uuid.Parse(req.PersonalId)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.CreateDoctor(ctx, storage.CreateDoctorParams{
		PersonalID:       id,
		DepartmentNumber: req.DepartmentNumber,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateDoctorResponse{
		Id:               resp.ID.String(),
		PersonalId:       resp.PersonalID.String(),
		DepartmentNumber: resp.DepartmentNumber,
		Timestamps: &pb.Timestamps1{
			CreatedAt: convertNullTime(resp.CreatedAt),
			UpdatedAt: convertNullTime(resp.UpdatedAt),
		},
	}, nil
}

func (q *AdminREPO) GetDoctorById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetDoctorByIdResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetDoctorById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDoctorByIdResponse{
		Doctor: &pb.Doctor{
			Info: &pb.Personal{
				Id:          resp.ID.String(),
				Profession:  resp.Profession,
				FullName:    resp.FullName,
				Email:       resp.Email,
				Password:    resp.Password,
				Address:     resp.Address.String,
				PhoneNumber: resp.PhoneNumber.String,
				Timestamps: &pb.Timestamps1{
					CreatedAt: convertNullTime(resp.CreatedAt),
					UpdatedAt: convertNullTime(resp.UpdatedAt),
				},
			},
			PersonalId:       resp.PersonalID.String(),
			DepartmentNumber: resp.DepartmentNumber,
		},
	}, nil
}

func (q *AdminREPO) ListDoctors(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListDoctorsResponse, error) {

	params := storage.ListDoctorsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListDoctors(ctx, params)
	if err != nil {
		return nil, err
	}

	var doctors []*pb.Doctor
	var totalCount int32

	for _, r := range resp {
		doctors = append(doctors, &pb.Doctor{
			Info: &pb.Personal{
				Id:          r.ID.String(),
				Profession:  r.Profession,
				FullName:    r.FullName,
				Email:       r.Email,
				Password:    r.Password,
				Address:     r.Address.String,
				PhoneNumber: r.PhoneNumber.String,
				Timestamps: &pb.Timestamps1{
					CreatedAt: convertNullTime(r.CreatedAt),
					UpdatedAt: convertNullTime(r.UpdatedAt),
				},
			},
			PersonalId:       r.PersonalID.String(),
			DepartmentNumber: r.DepartmentNumber,
		})
		totalCount = int32(r.TotalCount)
	}

	return &pb.ListDoctorsResponse{
		Doctors:    doctors,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdateDoctor(ctx context.Context, req *pb.UpdateDoctorRequest) (*pb.UpdateDoctorResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdateDoctor(ctx, storage.UpdateDoctorParams{
		ID:               id,
		DepartmentNumber: req.DepartmentNumber,
		UpdatedAt:        sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDoctorResponse{
		Id:               resp.ID.String(),
		PersonalId:       resp.PersonalID.String(),
		DepartmentNumber: resp.DepartmentNumber,
		Timestamps: &pb.Timestamps1{
			CreatedAt: convertNullTime(resp.CreatedAt),
			UpdatedAt: convertNullTime(resp.UpdatedAt),
		},
	}, nil
}

func (q *AdminREPO) DeleteDoctor(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeleteDoctor(ctx, storage.DeleteDoctorParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}
	return &pb.DeletePersonalResponse{
		Status: 204,
	}, nil
}
