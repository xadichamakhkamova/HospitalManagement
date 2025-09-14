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

	q.log.Infof("CreatePersonal called with FullName=%s, Email=%s", req.FullName, req.Email)

	resp, err := q.queries.CreatePersonal(ctx, storage.CreatePersonalParams{
		Profession:  req.Profession,
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		Address:     sql.NullString{String: req.Address, Valid: req.Address != ""},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
	})
	if err != nil {
		q.log.Errorf("CreatePersonal error: %v", err)
		return nil, err
	}

	q.log.Infof("Personal created successfully with ID=%s", resp.ID.String())
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

	q.log.Infof("GetPersonalById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetPersonalById(ctx, id)
	if err != nil {
		q.log.Errorf("GetPersonalById error: %v", err)
		return nil, err
	}

	q.log.Infof("Personal retrieved successfully with ID=%s", resp.ID.String())
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

	q.log.Infof("ListPersonals called with Search=%s, Limit=%d, Page=%d", req.Search, req.Limit, req.Page)

	params := storage.ListPersonalsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListPersonals(ctx, params)
	if err != nil {
		q.log.Errorf("ListPersonals error: %v", err)
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

	q.log.Infof("ListPersonals returned %d personals", len(personals))
	return &pb.ListPersonalsResponse{
		Personals:  personals,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdatePersonal(ctx context.Context, req *pb.UpdatePersonalRequest) (*pb.UpdatePersonalResponse, error) {

	q.log.Infof("UpdatePersonal called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
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
	if err != nil {
		q.log.Errorf("UpdatePersonal error: %v", err)
		return nil, err
	}

	q.log.Infof("Personal updated successfully with ID=%s", resp.ID.String())
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

	q.log.Infof("DeletePersonal called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeletePersonal(ctx, storage.DeletePersonalParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeletePersonal error: %v", err)
		return nil, err
	}

	q.log.Infof("Personal deleted successfully with ID=%s", req.Id)
	return &pb.DeletePersonalResponse{
		Status: 204,
	}, nil
}
