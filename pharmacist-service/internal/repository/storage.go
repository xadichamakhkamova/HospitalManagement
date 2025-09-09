package repository

import (
	"context"
	"database/sql"
	"pharmacist-service/internal/storage"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	//"pharmacist-service/logger"
)

type PharmaREPO struct {
	queries *storage.Queries
}

func NewPharmaSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *PharmaREPO) CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error) {
	resp, err := q.queries.
}