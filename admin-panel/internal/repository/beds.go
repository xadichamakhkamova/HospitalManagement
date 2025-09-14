package repository

import (
	"admin-panel/internal/storage"
	"context"
	"database/sql"
	"time"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/google/uuid"
)

func (q *AdminREPO) CreateBed(ctx context.Context, req *pb.CreateBedRequest) (*pb.CreateBedResponse, error) {

	q.log.Infof("CreateBed called with BedNumber=%d, BedType=%s", req.BedNumber, req.BedType.String())

	resp, err := q.queries.CreateBed(ctx, storage.CreateBedParams{
		BedNumber:   req.BedNumber,
		BedType:     storage.BedType(req.BedType.String()),
		Description: req.Description,
	})
	if err != nil {
		q.log.Errorf("CreateBed error: %v", err)
		return nil, err
	}

	bedType, _ := pb.BED_TYPE_value[string(resp.BedType)]
	status, _ := pb.BED_STATUS_value[string(resp.Status)]

	q.log.Infof("Bed created successfully with ID=%s", resp.ID.String())
	return &pb.CreateBedResponse{
		Bed: &pb.BedInfo{
			Id:          resp.ID.String(),
			BedNumber:   resp.BedNumber,
			BedType:     pb.BED_TYPE(bedType),
			Description: resp.Description,
			Status:      pb.BED_STATUS(status),
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) GetBedByID(ctx context.Context, req *pb.GetBedByIDRequest) (*pb.GetBedByIDResponse, error) {

	q.log.Infof("GetBedByID called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetBedByID(ctx, id)
	if err != nil {
		q.log.Errorf("GetBedByID error: %v", err)
		return nil, err
	}

	bedType, _ := pb.BED_TYPE_value[string(resp.BedType)]
	status, _ := pb.BED_STATUS_value[string(resp.Status)]

	q.log.Infof("Bed retrieved successfully with ID=%s", resp.ID.String())
	return &pb.GetBedByIDResponse{
		Bed: &pb.BedInfo{
			Id:          resp.ID.String(),
			BedNumber:   resp.BedNumber,
			BedType:     pb.BED_TYPE(bedType),
			Description: resp.Description,
			Status:      pb.BED_STATUS(status),
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) ListBedS(ctx context.Context, req *pb.ListBedSRequest) (*pb.ListBedSResponse, error) {
	
	q.log.Infof("ListBedS called with Search=%s, Status=%s, Limit=%d, Page=%d", req.Search, req.Status, req.Limit, req.Page)

	params := storage.ListBedsParams{
		Column1: req.Search,
		Column2: req.Status,
		Limit:   req.Limit,
		Column4: req.Page,
	}

	resp, err := q.queries.ListBeds(ctx, params)
	if err != nil {
		q.log.Errorf("ListBedS error: %v", err)
		return nil, err
	}

	var beds []*pb.BedInfo
	var totalCount int64
	for _, r := range resp {
		bedType, _ := pb.BED_TYPE_value[string(r.BedType)]
		status, _ := pb.BED_STATUS_value[string(r.Status)]

		beds = append(beds, &pb.BedInfo{
			Id:          r.ID.String(),
			BedNumber:   r.BedNumber,
			BedType:     pb.BED_TYPE(bedType),
			Description: r.Description,
			Status:      pb.BED_STATUS(status),
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(r.CreatedAt),
				UpdatedAt: convertNullTime(r.UpdatedAt),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListBedS returned %d beds", len(beds))
	return &pb.ListBedSResponse{
		Beds:       beds,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *AdminREPO) UpdateBed(ctx context.Context, req *pb.UpdateBedRequest) (*pb.UpdateBedResponse, error) {

	q.log.Infof("UpdateBed called with ID=%s, BedNumber=%d", req.Id, req.BedNumber)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdateBed(ctx, storage.UpdateBedParams{
		ID:          id,
		BedNumber:   req.BedNumber,
		BedType:     storage.BedType(req.BedType.String()),
		Description: req.Description,
		Status:      storage.BedStatus(req.Status.String()),
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("UpdateBed error: %v", err)
		return nil, err
	}

	bedType, _ := pb.BED_TYPE_value[string(resp.BedType)]
	status, _ := pb.BED_STATUS_value[string(resp.Status)]

	q.log.Infof("Bed updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdateBedResponse{
		Bed: &pb.BedInfo{
			Id:          resp.ID.String(),
			BedNumber:   resp.BedNumber,
			BedType:     pb.BED_TYPE(bedType),
			Description: resp.Description,
			Status:      pb.BED_STATUS(status),
			Timestamps: &pb.Timestamps1{
				CreatedAt: convertNullTime(resp.CreatedAt),
				UpdatedAt: convertNullTime(resp.UpdatedAt),
			},
		},
	}, nil
}

func (q *AdminREPO) DeleteBed(ctx context.Context, req *pb.DeleteBedRequest) (*pb.DeleteBedResponse, error) {

	q.log.Infof("DeleteBed called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeleteBed(ctx, storage.DeleteBedParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeleteBed error: %v", err)
		return nil, err
	}

	q.log.Infof("Bed deleted successfully with ID=%s", req.Id)
	return &pb.DeleteBedResponse{
		Status: 204,
	}, nil
}
