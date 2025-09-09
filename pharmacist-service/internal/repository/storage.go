package repository

import (
	"context"
	"database/sql"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	"pharmacist-service/internal/storage"
	//"pharmacist-service/logger"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type PharmaREPO struct {
	queries *storage.Queries
}

func NewPharmaSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}

func (q *PharmaREPO) CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error) {
	resp, err := q.queries.CreateMedicine(ctx, storage.CreateMedicineParams{
		Name: req.Name,
		Category: storage.MedicineCategory(req.Category),
		Description: req.Description,
		Price: ,
	})
}