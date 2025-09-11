package repository

import (
	"context"
	"database/sql"
	"pharmacist-service/internal/storage"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"pharmacist-service/logger"

	"github.com/google/uuid"
)

type PharmaREPO struct {
	queries *storage.Queries
}

func NewPharmaSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

// NullTime → *timestamppb.Timestamp converter
func convertNullTime(nt sql.NullTime) *timestamppb.Timestamp {
    if nt.Valid {
        return timestamppb.New(nt.Time)
    }
    return nil
}

func (q *PharmaREPO) CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error) {

	resp, err := q.queries.CreateMedicine(ctx, storage.CreateMedicineParams{
		Name: req.Name,
		Category: storage.MedicineCategory(req.Category),
		Description: req.Description,
		Price: float64(req.Price),
		Company: req.Company,
		Status: storage.MedicineStatus(req.Status.String()),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateMedicineResponse{
		Medicine: &pb.Medicine{
			Id: resp.ID.String(),
			Name: resp.Name,
			Category: string(resp.Category),
			Description: resp.Description,
			Price: float32(resp.Price),
			Company: resp.Company,
			Status: pb.MedicineStatus(pb.MedicineStatus_value[string(resp.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *PharmaREPO) GetMedicineById(ctx context.Context, req *pb.GetMedicineByIdRequest) (*pb.GetMedicineByIdResponse, error) {

	
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetMedicineById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetMedicineByIdResponse{
		Medicine: &pb.Medicine{
			Id: resp.ID.String(),
			Name: resp.Name,
			Category: string(resp.Category),
			Description: resp.Description,
			Price: float32(resp.Price),
			Company: resp.Company,
			Status: pb.MedicineStatus(pb.MedicineStatus_value[string(resp.Status)]),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil

}
