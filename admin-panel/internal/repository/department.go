package repository

import (
	"admin-panel/internal/storage"
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/google/uuid"
)

type AdminREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewAdminSqlc(db *sql.DB, log *logrus.Logger) *AdminREPO {
	return &AdminREPO{
		queries: storage.New(db),
		log:     log,
	}
}
func (q *AdminREPO) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error) {

	q.log.Infof("CreateDepartment called with Name=%s, Number=%d", req.Name, req.Number)

	resp, err := q.queries.CreateDepartment(ctx, storage.CreateDepartmentParams{
		Name:        req.Name,
		Number:      int32(req.Number),
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
	})
	if err != nil {
		q.log.Errorf("CreateDepartment error: %v", err)
		return nil, err
	}

	q.log.Infof("Department created successfully with ID=%s", resp.ID.String())
	return &pb.CreateDepartmentResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *AdminREPO) GetDepartmentById(ctx context.Context, req *pb.GetDepartmentByIdRequest) (*pb.GetDepartmentByIdResponse, error) {

	q.log.Infof("GetDepartmentById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetDepartmentById(ctx, id)
	if err != nil {
		q.log.Errorf("GetDepartmentById error: %v", err)
		return nil, err
	}

	q.log.Infof("Department retrieved successfully with ID=%s", resp.ID.String())
	return &pb.GetDepartmentByIdResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *AdminREPO) ListDeparments(ctx context.Context, req *pb.ListDepartmentsRequest) (*pb.ListDepartmentsResponse, error) {

	q.log.Infof("ListDepartments called with Search=%s, Limit=%d, Page=%d", req.Search, req.Limit, req.Page)

	params := storage.ListDepartmentsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Column3: req.Page,
	}

	resp, err := q.queries.ListDepartments(ctx, params)
	if err != nil {
		q.log.Errorf("ListDepartments error: %v", err)
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
				CreatedAt: r.CreatedAt.Time.String(),
				UpdatedAt: r.UpdatedAt.Time.String(),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListDepartments returned %d departments", len(departments))
	return &pb.ListDepartmentsResponse{
		Deparments: departments,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error) {

	q.log.Infof("UpdateDepartment called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
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
		q.log.Errorf("UpdateDepartment error: %v", err)
		return nil, err
	}

	q.log.Infof("Department updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdateDepartmentResponse{
		Department: &pb.Department{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Number:      int64(resp.Number),
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps1{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *AdminREPO) DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error) {

	q.log.Infof("DeleteDepartment called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeleteDepartment(ctx, storage.DeleteDepartmentParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeleteDepartment error: %v", err)
		return nil, err
	}

	q.log.Infof("Department deleted successfully with ID=%s", req.Id)
	return &pb.DeleteDepartmentResponse{
		Status: 204,
	}, nil
}
