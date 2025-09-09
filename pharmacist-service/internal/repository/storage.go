package repository

import (
	"context"
	"database/sql"
	"pharmacist-service/internal/storage"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"pharmacist-service/logger"
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
		Status: storage.MedicineStatusMEDICINESTATUSAVAILABLE,
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
			Status: pb.MedicineStatus(req.Status),
			Timestamps: &pb.Timestamps5{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}