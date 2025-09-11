package repository

import (
	"admin-panel/internal/storage"
	"context"
	"database/sql"

	"time"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/google/uuid"
)

func (q *AdminREPO) CreatePersonal(ctx context.Context, req *pb.CreatePersonalRequest) (*pb.CreatePersonalResponse, error) {

	resp, err := q.queries.CreatePersonal(ctx, storage.CreatePersonalParams{
		Profession:  req.Profession,
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Address:     sql.NullString{String: req.Address, Valid: req.Address != ""},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreatePersonalResponse{
		Personal: &pb.Personal{
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
	}, nil
}

func (q *AdminREPO) GetPersonalById(ctx context.Context, req *pb.GetPersonalByIdRequest) (*pb.GetPersonalByIdResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetPersonalById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPersonalByIdResponse{
		Personal: &pb.Personal{
			Id:          resp.ID.String(),
			Profession:  string(resp.Profession),
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
	}, nil
}

func (q *AdminREPO) ListPersonals(ctx context.Context, req *pb.ListPersonalsRequest) (*pb.ListPersonalsResponse, error) {

	params := storage.ListPersonalsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListPersonals(ctx, params)
	if err != nil {
		return nil, err
	}

	var personals []*pb.Personal
	var totalCount int32

	for _, r := range resp {
		personals = append(personals, &pb.Personal{
			Id:          r.ID.String(),
			Profession:  string(r.Profession),
			FullName:    r.FullName,
			Email:       r.Email,
			Password:    r.Password,
			Address:     r.Address.String,
			PhoneNumber: r.PhoneNumber.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = int32(r.TotalCount)
	}

	return &pb.ListPersonalsResponse{
		Personals:  personals,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdatePersonal(ctx context.Context, req *pb.UpdatePersonalRequest) (*pb.UpdatePersonalResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdatePersonal(ctx, storage.UpdatePersonalParams{
		ID:          id,
		Profession:  req.Profession,
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Address:     sql.NullString{String: req.Address, Valid: req.Address != ""},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	})

	return &pb.UpdatePersonalResponse{
		Personal: &pb.Personal{
			Id:          resp.ID.String(),
			Profession:  string(resp.Profession),
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
	}, nil
}

func (q *AdminREPO) DeletePersonal(ctx context.Context, req *pb.DeletePersonalRequest) (*pb.DeletePersonalResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeletePersonal(ctx, storage.DeletePersonalParams{
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
