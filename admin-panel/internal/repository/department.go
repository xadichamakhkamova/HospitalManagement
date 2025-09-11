package repository

import (
	"admin-panel/internal/storage"
	"context"
	"database/sql"
	"time"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
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

type AdminREPO struct {
	queries *storage.Queries
}

func NewAdminSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *AdminREPO) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error) {

	resp, err := q.queries.CreateDepartment(ctx, storage.CreateDepartmentParams{
		Name:        req.Name,
		Number:      int32(req.Number),
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateDepartmentResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) GetDepartmentById(ctx context.Context, req *pb.GetDepartmentByIdRequest) (*pb.GetDepartmentByIdResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetDepartmentById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDepartmentByIdResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) ListDeparments(ctx context.Context, req *pb.ListDepartmentsRequest) (*pb.ListDepartmentsResponse, error) {

	params := storage.ListDepartmentsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListDepartments(ctx, params)
	if err != nil {
		return nil, err
	}

	var departments []*pb.Department
	var totalCount int64

	for _, r := range resp {
		departments = append(departments, &pb.Department{
			Id:          r.ID.String(),
			Name:        r.Name,
			Number:      int64(r.Number),
			Description: r.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = r.TotalCount
	}

	return &pb.ListDepartmentsResponse{
		Deparments: departments,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdateDepartment(ctx, storage.UpdateDepartmentParams{
		ID:          id,
		Name:        req.Name,
		Number:      int32(req.Number),
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDepartmentResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeleteDepartment(ctx, storage.DeleteDepartmentParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return nil, err
	}
	return &pb.DeleteDepartmentResponse{
		Status: 204,
	}, nil
}
