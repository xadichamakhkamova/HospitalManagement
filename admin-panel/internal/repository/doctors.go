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

	q.log.Infof("CreateDoctor called with PersonalID=%s, DepartmentNumber=%d", req.PersonalId, req.DepartmentNumber)

	id, err := uuid.Parse(req.PersonalId)
	if err != nil {
		q.log.Errorf("Invalid PersonalID UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.CreateDoctor(ctx, storage.CreateDoctorParams{
		PersonalID:       id,
		DepartmentNumber: req.DepartmentNumber,
	})
	if err != nil {
		q.log.Errorf("CreateDoctor error: %v", err)
		return nil, err
	}

	q.log.Infof("Doctor created successfully with ID=%s", resp.ID.String())
	return &pb.CreateDoctorResponse{
		Id:               resp.ID.String(),
		PersonalId:       resp.PersonalID.String(),
		DepartmentNumber: resp.DepartmentNumber,
		Timestamps: &pb.Timestamps1{
			CreatedAt: resp.CreatedAt.Time.String(),
			UpdatedAt: resp.UpdatedAt.Time.String(),
		},
	}, nil
}

func (q *AdminREPO) GetDoctorById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetDoctorByIdResponse, error) {

	q.log.Infof("GetDoctorById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetDoctorById(ctx, id)
	if err != nil {
		q.log.Errorf("GetDoctorById error: %v", err)
		return nil, err
	}

	q.log.Infof("Doctor retrieved successfully with ID=%s", resp.ID.String())
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
					CreatedAt: resp.CreatedAt.Time.String(),
					UpdatedAt: resp.UpdatedAt.Time.String(),
				},
			},
			PersonalId:       resp.PersonalID.String(),
			DepartmentNumber: resp.DepartmentNumber,
		},
	}, nil
}

func (q *AdminREPO) ListDoctors(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListDoctorsResponse, error) {

	q.log.Infof("ListDoctors called with Search=%s, Limit=%d, Page=%d", req.Search, req.Limit, req.Page)

	params := storage.ListDoctorsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListDoctors(ctx, params)
	if err != nil {
		q.log.Errorf("ListDoctors error: %v", err)
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
					CreatedAt: r.CreatedAt.Time.String(),
					UpdatedAt: r.UpdatedAt.Time.String(),
				},
			},
			PersonalId:       r.PersonalID.String(),
			DepartmentNumber: r.DepartmentNumber,
		})
		totalCount = int32(r.TotalCount)
	}

	q.log.Infof("ListDoctors returned %d doctors", len(doctors))
	return &pb.ListDoctorsResponse{
		Doctors:    doctors,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdateDoctor(ctx context.Context, req *pb.UpdateDoctorRequest) (*pb.UpdateDoctorResponse, error) {

	q.log.Infof("UpdateDoctor called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdateDoctor(ctx, storage.UpdateDoctorParams{
		ID:               id,
		DepartmentNumber: req.DepartmentNumber,
		UpdatedAt:        sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("UpdateDoctor error: %v", err)
		return nil, err
	}

	q.log.Infof("Doctor updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdateDoctorResponse{
		Id:               resp.ID.String(),
		PersonalId:       resp.PersonalID.String(),
		DepartmentNumber: resp.DepartmentNumber,
		Timestamps: &pb.Timestamps1{
			CreatedAt: resp.CreatedAt.Time.String(),
			UpdatedAt: resp.UpdatedAt.Time.String(),
		},
	}, nil
}

func (q *AdminREPO) DeleteDoctor(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error) {

	q.log.Infof("DeleteDoctor called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeleteDoctor(ctx, storage.DeleteDoctorParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeleteDoctor error: %v", err)
		return nil, err
	}

	q.log.Infof("Doctor deleted successfully with ID=%s", req.Id)
	return &pb.DeletePersonalResponse{
		Status: 204,
	}, nil
}
