package repository

import (
	"context"
	"database/sql"
	"pharmacist-service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	//"pharmacist-service/logger"

	"github.com/google/uuid"
)

// NullTime → *timestamppb.Timestamp converter
func convertNullTime(nt sql.NullTime) *timestamppb.Timestamp {
	if nt.Valid {
		return timestamppb.New(nt.Time)
	}
	return nil
}

type PharmaREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewPharmaSqlc(db *sql.DB, log *logrus.Logger) *PharmaREPO {
	return &PharmaREPO{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *PharmaREPO) CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error) {

	q.log.Infof("CreateMedicine called with Name=%s, Company=%s, Price=%.2f", req.Name, req.Company, req.Price)

	resp, err := q.queries.CreateMedicine(ctx, storage.CreateMedicineParams{
		Name:        req.Name,
		Category:    storage.MedicineCategory(req.Category),
		Description: req.Description,
		Price:       float64(req.Price),
		Company:     req.Company,
		Status:      storage.MedicineStatus(req.Status.String()),
	})
	if err != nil {
		q.log.Errorf("CreateMedicine DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Medicine created successfully with ID=%s", resp.ID.String())
	return &pb.CreateMedicineResponse{
		Medicine: &pb.Medicine{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Category:    pb.MedicineCategoryType(pb.MedicineCategoryType_value[string(resp.Category)]),
			Description: resp.Description,
			Price:       float32(resp.Price),
			Company:     resp.Company,
			Status:      pb.MedicineStatus(pb.MedicineStatus_value[string(resp.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PharmaREPO) GetMedicineById(ctx context.Context, req *pb.GetMedicineByIdRequest) (*pb.GetMedicineByIdResponse, error) {

	q.log.Infof("GetMedicineById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Medicine ID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetMedicineById(ctx, id)
	if err != nil {
		q.log.Errorf("GetMedicineById DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Medicine retrieved successfully with ID=%s", resp.ID.String())
	return &pb.GetMedicineByIdResponse{
		Medicine: &pb.Medicine{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Category:    pb.MedicineCategoryType(pb.MedicineCategoryType_value[string(resp.Category)]),
			Description: resp.Description,
			Price:       float32(resp.Price),
			Company:     resp.Company,
			Status:      pb.MedicineStatus(pb.MedicineStatus_value[string(resp.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PharmaREPO) ListMedicines(ctx context.Context, req *pb.ListMedicinesRequest) (*pb.ListMedicinesResponse, error) {

	q.log.Infof("ListMedicines called with Search=%s, Status=%s, Limit=%d, Page=%d", req.Search, req.Status, req.Limit, req.Page)

	params := storage.ListMedicinesParams{
		Column1: req.Search,
		Column2: req.Status,
		Limit:   req.Limit,
		Column4: req.Page,
	}

	resp, err := q.queries.ListMedicines(ctx, params)
	if err != nil {
		q.log.Errorf("ListMedicines DB error: %v", err)
		return nil, err
	}

	var medicines []*pb.Medicine
	var totalCount int64

	for _, r := range resp {
		medicines = append(medicines, &pb.Medicine{
			Id:          r.ID.String(),
			Name:        r.Name,
			Category:    pb.MedicineCategoryType(pb.MedicineCategoryType_value[string(r.Category)]),
			Description: r.Description,
			Price:       float32(r.Price),
			Company:     r.Company,
			Status:      pb.MedicineStatus(pb.MedicineStatus_value[string(r.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListMedicines returned %d medicines", len(medicines))
	return &pb.ListMedicinesResponse{
		Medicines:  medicines,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *PharmaREPO) UpdateMedicine(ctx context.Context, req *pb.UpdateMedicineRequest) (*pb.UpdateMedicineResponse, error) {

	q.log.Infof("UpdateMedicine called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Medicine ID: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdateMedicine(ctx, storage.UpdateMedicineParams{
		ID:          id,
		Name:        req.Name,
		Category:    storage.MedicineCategory(req.Category),
		Description: req.Description,
		Price:       float64(req.Price),
		Company:     req.Company,
		Status:      storage.MedicineStatus(req.Status.String()),
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("UpdateMedicine DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Medicine updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdateMedicineResponse{
		Medicine: &pb.Medicine{
			Id:          resp.ID.String(),
			Name:        resp.Name,
			Category:    pb.MedicineCategoryType(pb.MedicineCategoryType_value[string(resp.Category)]),
			Description: resp.Description,
			Price:       float32(resp.Price),
			Company:     resp.Company,
			Status:      pb.MedicineStatus(pb.MedicineStatus_value[string(resp.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PharmaREPO) DeleteMedicine(ctx context.Context, req *pb.DeleteMedicineRequest) (*pb.DeleteMedicineResponse, error) {

	q.log.Infof("DeleteMedicine called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid Medicine ID: %v", err)
		return nil, err
	}

	err = q.queries.DeleteMedicine(ctx, storage.DeleteMedicineParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeleteMedicine DB error: %v", err)
		return nil, err
	}

	q.log.Infof("Medicine deleted successfully with ID=%s", req.Id)
	return &pb.DeleteMedicineResponse{
		Status: 204,
	}, nil
}
