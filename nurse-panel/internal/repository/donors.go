package repository

import (
	"context"
	"database/sql"
	"nurse-service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"

	"github.com/google/uuid"
)

type NurseREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewNurseSqlc(db *sql.DB, log *logrus.Logger) *NurseREPO {
	return &NurseREPO{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *NurseREPO) CreateDonor(ctx context.Context, req *pb.CreateDonorRequest) (*pb.CreateDonorResponse, error) {

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.CreateDonor(ctx, storage.CreateDonorParams{
		FullName:        req.FullName,
		Email:           req.Email,
		Password:        req.Password,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		Gender:          storage.GenderType(req.Gender),
		BirthDate:       birthDate,
		BloodGroup:      storage.BloodType(req.BloodGroup),
		Weight:          int16(req.Weight),
		HealthCondition: storage.HealthConditionType(req.HealthConditions.String()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateDonorResponse{
		Donor: &pb.PatientDonor{
			Id:               resp.ID.String(),
			FullName:         resp.FullName,
			Email:            resp.Email,
			Password:         resp.Password,
			Address:          resp.Address,
			PhoneNumber:      resp.PhoneNumber,
			Gender:           string(resp.Gender),
			BirthDate:        resp.BirthDate.Format("2006-01-02"),
			BloodGroup:       string(resp.BloodGroup),
			Weight:           float64(resp.Weight),
			HealthConditions: pb.HealthConditionType(pb.HealthConditionType_value[string(resp.HealthCondition)]),
			Timestamps: &pb.Timestamps3{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *NurseREPO) GetDonorById(ctx context.Context, req *pb.GetDonorByIdRequest) (*pb.GetDonorByIdResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.GetDonorById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDonorByIdResponse{
		Donor: &pb.PatientDonor{
			Id:               resp.ID.String(),
			FullName:         resp.FullName,
			Email:            resp.Email,
			Password:         resp.Password,
			Address:          resp.Address,
			PhoneNumber:      resp.PhoneNumber,
			Gender:           string(resp.Gender),
			BirthDate:        resp.BirthDate.Format("2006-01-02"),
			BloodGroup:       string(resp.BloodGroup),
			Weight:           float64(resp.Weight),
			HealthConditions: pb.HealthConditionType(pb.HealthConditionType_value[string(resp.HealthCondition)]),
			Timestamps: &pb.Timestamps3{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil

}

func (q *NurseREPO) ListDonors(ctx context.Context, req *pb.ListDonorsRequest) (*pb.ListDonorsResponse, error) {

	params := storage.ListDonorsParams{
		Column1: req.Search,
		Column2: req.OnlyEligible,
		Limit:   req.Limit,
		Column4: req.Page,
	}

	resp, err := q.queries.ListDonors(ctx, params)
	if err != nil {
		return nil, err
	}

	var donors []*pb.PatientDonor
	var totalCount int64

	for _, r := range resp {
		donors = append(donors, &pb.PatientDonor{
			Id:               r.ID.String(),
			FullName:         r.FullName,
			Email:            r.Email,
			Password:         r.Password,
			Address:          r.Address,
			PhoneNumber:      r.PhoneNumber,
			Gender:           string(r.Gender),
			BirthDate:        r.BirthDate.Format("2006-01-02"),
			BloodGroup:       string(r.BloodGroup),
			LastDonation:     r.LastDonation.Time.String(),
			DonationCount:    r.DonationCount.Int32,
			IsEligible:       r.IsEligible.Bool,
			LastCheckupDate:  r.LastCheckupDate.Time.String(),
			Weight:           float64(r.Weight),
			HealthConditions: pb.HealthConditionType(pb.HealthConditionType_value[string(r.HealthCondition)]),
			DonationLocation: r.DonationLocation.String,
			Timestamps: &pb.Timestamps3{
				CreatedAt: r.CreatedAt.Time.String(),
				UpdatedAt: r.UpdatedAt.Time.String(),
			},
		})
		totalCount = r.TotalCount
	}

	return &pb.ListDonorsResponse{
		Patients:   donors,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *NurseREPO) UpdateDonor(ctx context.Context, req *pb.UpdateDonorRequest) (*pb.UpdateDonorResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.UpdateDonor(ctx, storage.UpdateDonorParams{
		ID:              id,
		FullName:        req.FullName,
		Email:           req.Email,
		Password:        req.Password,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		Gender:          storage.GenderType(req.Gender),
		BirthDate:       birthDate,
		BloodGroup:      storage.BloodType(req.BloodGroup),
		Weight:          int16(req.Weight),
		HealthCondition: storage.HealthConditionType(req.HealthConditions.String()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDonorResponse{
		Donor: &pb.PatientDonor{
			Id:               resp.ID.String(),
			FullName:         resp.FullName,
			Email:            resp.Email,
			Password:         resp.Password,
			Address:          resp.Address,
			PhoneNumber:      resp.PhoneNumber,
			Gender:           string(resp.Gender),
			BirthDate:        resp.BirthDate.Format("2006-01-02"),
			BloodGroup:       string(resp.BloodGroup),
			Weight:           float64(resp.Weight),
			HealthConditions: pb.HealthConditionType(pb.HealthConditionType_value[string(resp.HealthCondition)]),
			Timestamps: &pb.Timestamps3{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *NurseREPO) DeleteDonor(ctx context.Context, req *pb.DeleteDonorRequest) (*pb.DeleteDonorResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = q.queries.DeleteDonor(ctx, storage.DeleteDonorParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteDonorResponse{
		Status: 204,
	}, nil
}

func (q *NurseREPO) RegisterDonation(ctx context.Context, req *pb.RegisterDonationRequest) (*pb.RegisterDonationResponse, error) {

	id, err := uuid.Parse(req.DonorId)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.RegisterDonation(ctx, storage.RegisterDonationParams{
		ID:               id,
		DonationLocation: sql.NullString{String: req.DonationLocation, Valid: req.DonationLocation != ""},
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterDonationResponse{
		LastDonation:  resp.LastDonation.Time.String(),
		DonationCount: resp.DonationCount.Int32,
		IsEligible:    resp.IsEligible.Bool,
	}, nil
}

func (q *NurseREPO) RegisterCheckup(ctx context.Context, req *pb.RegisterCheckupRequest) (*pb.RegisterCheckupResponse, error) {

	id, err := uuid.Parse(req.DonorId)
	if err != nil {
		return nil, err
	}

	resp, err := q.queries.RegisterCheckup(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterCheckupResponse{
		LastCheckupDate: resp.LastCheckupDate.Time.String(),
		IsEligible:      resp.IsEligible.Bool,
	}, nil
}
