package repository

import (
	"context"
	"database/sql"
	"doctor-service/internal/storage"
	"time"

	"github.com/google/uuid"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
)

func (q *DoctorREPO) CreatePrescription(ctx context.Context, req *pb.CreatePrescriptionRequest) (*pb.CreatePrescriptionResponse, error) {

	q.log.Infof("CreatePrescription called with DoctorId=%s, PatientId=%s", req.DoctorId, req.PatientId)

	doctor_id, err := uuid.Parse(req.DoctorId)
	if err != nil {
		q.log.Errorf("Invalid Doctor UUID: %v", err)
		return nil, err
	}
	patient_id, err := uuid.Parse(req.PatientId)
	if err != nil {
		q.log.Errorf("Invalid Patient UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.CreatePrescription(ctx, storage.CreatePrescriptionParams{
		DoctorID:    doctor_id,
		PatientID:   patient_id,
		CaseHistory: req.CaseHistory,
		Medication:  req.Medication,
		Description: sql.NullString{String: req.Description, Valid: true},
	})
	if err != nil {
		q.log.Errorf("CreatePrescription error: %v", err)
		return nil, err
	}

	q.log.Infof("Prescription created successfully with ID=%s", resp.ID.String())
	return &pb.CreatePrescriptionResponse{
		Presc: &pb.Prescription{
			Id:          resp.ID.String(),
			DoctorId:    resp.DoctorID.String(),
			PatientId:   resp.PatientID.String(),
			CaseHistory: resp.CaseHistory,
			Medication:  resp.Medication,
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) GetPrescriptionById(ctx context.Context, req *pb.GetPrescriptionByIdRequest) (*pb.GetPrescriptionByIdResponse, error) {

	q.log.Infof("GetPrescriptionById called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.GetPrescriptionById(ctx, id)
	if err != nil {
		q.log.Errorf("GetPrescriptionById error: %v", err)
		return nil, err
	}

	q.log.Infof("Prescription retrieved successfully with ID=%s", resp.ID.String())
	return &pb.GetPrescriptionByIdResponse{
		Presc: &pb.Prescription{
			Id:          resp.ID.String(),
			DoctorId:    resp.DoctorID.String(),
			PatientId:   resp.PatientID.String(),
			CaseHistory: resp.CaseHistory,
			Medication:  resp.Medication,
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) ListPrescriptions(ctx context.Context, req *pb.ListPrescriptionsRequest) (*pb.ListPrescriptionsResponse, error) {

	q.log.Infof("ListPrescriptions called with Limit=%d, Page=%d", req.Limit, req.Page)

	params := storage.ListPrescriptionsParams{
		Limit:   req.Limit,
		Column2: req.Page,
	}

	resp, err := q.queries.ListPrescriptions(ctx, params)
	if err != nil {
		q.log.Errorf("ListPrescriptions error: %v", err)
		return nil, err
	}

	var prescs []*pb.Prescription
	var totalCount int64
	for _, r := range resp {
		prescs = append(prescs, &pb.Prescription{
			Id:          r.ID.String(),
			DoctorId:    r.DoctorID.String(),
			PatientId:   r.PatientID.String(),
			CaseHistory: r.CaseHistory,
			Medication:  r.Medication,
			Description: r.Description.String,
			Timestamps: &pb.Timestamps2{
				CreatedAt: r.CreatedAt.Time.String(),
				UpdatedAt: r.UpdatedAt.Time.String(),
			},
		})
		totalCount = r.TotalCount
	}

	q.log.Infof("ListPrescriptions returned %d prescriptions", len(prescs))
	return &pb.ListPrescriptionsResponse{
		Presc:      prescs,
		TotalCount: int32(totalCount),
	}, nil
}

func (q *DoctorREPO) UpdatePrescription(ctx context.Context, req *pb.UpdatePrescriptionRequest) (*pb.UpdatePrescriptionResponse, error) {

	q.log.Infof("UpdatePrescription called with ID=%s, DoctorId=%s, PatientId=%s", req.Presc.Id, req.Presc.DoctorId, req.Presc.PatientId)

	id, err := uuid.Parse(req.Presc.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}
	doctor_id, err := uuid.Parse(req.Presc.DoctorId)
	if err != nil {
		q.log.Errorf("Invalid Doctor UUID: %v", err)
		return nil, err
	}
	patient_id, err := uuid.Parse(req.Presc.PatientId)
	if err != nil {
		q.log.Errorf("Invalid Patient UUID: %v", err)
		return nil, err
	}

	resp, err := q.queries.UpdatePrescription(ctx, storage.UpdatePrescriptionParams{
		ID:          id,
		DoctorID:    doctor_id,
		PatientID:   patient_id,
		CaseHistory: req.Presc.CaseHistory,
		Medication:  req.Presc.Medication,
		Description: sql.NullString{String: req.Presc.Description, Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("UpdatePrescription error: %v", err)
		return nil, err
	}

	q.log.Infof("Prescription updated successfully with ID=%s", resp.ID.String())
	return &pb.UpdatePrescriptionResponse{
		Presc: &pb.Prescription{
			Id:          resp.ID.String(),
			DoctorId:    resp.DoctorID.String(),
			PatientId:   resp.PatientID.String(),
			CaseHistory: resp.CaseHistory,
			Medication:  resp.Medication,
			Description: resp.Description.String,
			Timestamps: &pb.Timestamps2{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
			},
		},
	}, nil
}

func (q *DoctorREPO) DeletePrescription(ctx context.Context, req *pb.DeletePrescriptionRequest) (*pb.DeletePrescriptionResponse, error) {

	q.log.Infof("DeletePrescription called with ID=%s", req.Id)

	id, err := uuid.Parse(req.Id)
	if err != nil {
		q.log.Errorf("Invalid UUID: %v", err)
		return nil, err
	}

	err = q.queries.DeletePrescription(ctx, storage.DeletePrescriptionParams{
		ID:        id,
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		q.log.Errorf("DeletePrescription error: %v", err)
		return nil, err
	}

	q.log.Infof("Prescription deleted successfully with ID=%s", req.Id)
	return &pb.DeletePrescriptionResponse{
		Status: 204,
	}, nil
}
