package repository

import (
	"context"
	"pharmacist-service/internal/storage"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
)

func NewIPharmacistRepository(queries *storage.Queries) IPharmacistRepository {
	return &PharmaREPO{
		queries: queries,
	}
}

type IPharmacistRepository interface {
	//Medicines
	CreateMedicine(ctx context.Context, req *pb.CreateMedicineRequest) (*pb.CreateMedicineResponse, error)
	GetMedicineById(ctx context.Context, req *pb.GetMedicineByIdRequest) (*pb.GetMedicineByIdResponse, error)
	ListMedicines(ctx context.Context, req *pb.ListMedicinesRequest) (*pb.ListMedicinesResponse, error)
	UpdateMedicine(ctx context.Context, req *pb.UpdateMedicineRequest) (*pb.UpdateMedicineResponse, error)
	DeleteMedicine(ctx context.Context, req *pb.DeleteMedicineRequest) (*pb.DeleteMedicineResponse, error)
}
