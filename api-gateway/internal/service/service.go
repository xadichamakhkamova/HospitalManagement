package service

import (
	"context"

	pbAdmin "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
	pbDoctor "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
	pbNurse "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"
	pbPatient "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
	pbPharmacist "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
)

type ServiceRepositoryClient struct {
	adminClient      pbAdmin.AdminServiceClient
	doctorClient     pbDoctor.DoctorServiceClient
	nurseClient      pbNurse.NurseServiceClient
	patientClient    pbPatient.PatientManagementServiceClient
	pharmacistClient pbPharmacist.PharmacistServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pbAdmin.AdminServiceClient,
	conn2 *pbDoctor.DoctorServiceClient,
	conn3 *pbNurse.NurseServiceClient,
	conn4 *pbPatient.PatientManagementServiceClient,
	conn5 *pbPharmacist.PharmacistServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		adminClient:      *conn1,
		doctorClient:     *conn2,
		nurseClient:      *conn3,
		patientClient:    *conn4,
		pharmacistClient: *conn5,
	}
}

//! Admin Service
// ------------------- Department -------------------

func (s *ServiceRepositoryClient) CreateDepartment(ctx context.Context, req *pbAdmin.CreateDepartmentRequest) (*pbAdmin.CreateDepartmentResponse, error) {
	return s.CreateDepartment(ctx, req)
}

func (s *ServiceRepositoryClient) GetDepartmentById(ctx context.Context, req *pbAdmin.GetDepartmentByIdRequest) (*pbAdmin.GetDepartmentByIdResponse, error) {
	return s.GetDepartmentById(ctx, req)
}

func (s *ServiceRepositoryClient) ListDeparments(ctx context.Context, req *pbAdmin.ListDepartmentsRequest) (*pbAdmin.ListDepartmentsResponse, error) {
	return s.ListDeparments(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateDepartment(ctx context.Context, req *pbAdmin.UpdateDepartmentRequest) (*pbAdmin.UpdateDepartmentResponse, error) {
	return s.UpdateDepartment(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteDepartment(ctx context.Context, req *pbAdmin.DeleteDepartmentRequest) (*pbAdmin.DeleteDepartmentResponse, error) {
	return s.DeleteDepartment(ctx, req)
}

// ------------------- Personal -------------------

func (s *ServiceRepositoryClient) CreatePersonal(ctx context.Context, req *pbAdmin.CreatePersonalRequest) (*pbAdmin.CreatePersonalResponse, error) {
	return s.CreatePersonal(ctx, req)
}

func (s *ServiceRepositoryClient) GetPersonalById(ctx context.Context, req *pbAdmin.GetPersonalByIdRequest) (*pbAdmin.GetPersonalByIdResponse, error) {
	return s.GetPersonalById(ctx, req)
}

func (s *ServiceRepositoryClient) ListPersonals(ctx context.Context, req *pbAdmin.ListPersonalsRequest) (*pbAdmin.ListPersonalsResponse, error) {
	return s.ListPersonals(ctx, req)
}

func (s *ServiceRepositoryClient) UpdatePersonal(ctx context.Context, req *pbAdmin.UpdatePersonalRequest) (*pbAdmin.UpdatePersonalResponse, error) {
	return s.UpdatePersonal(ctx, req)
}

func (s *ServiceRepositoryClient) DeletePersonal(ctx context.Context, req *pbAdmin.DeletePersonalRequest) (*pbAdmin.DeletePersonalResponse, error) {
	return s.DeletePersonal(ctx, req)
}

// ------------------- Doctor -------------------

func (s *ServiceRepositoryClient) CreateDoctor(ctx context.Context, req *pbAdmin.CreateDoctorRequest) (*pbAdmin.CreateDoctorResponse, error) {
	return s.CreateDoctor(ctx, req)
}

func (s *ServiceRepositoryClient) GetDoctorById(ctx context.Context, req *pbAdmin.GetPersonalByIdRequest) (*pbAdmin.GetDoctorByIdResponse, error) {
	return s.GetDoctorById(ctx, req)
}

func (s *ServiceRepositoryClient) ListDoctors(ctx context.Context, req *pbAdmin.ListPersonalsRequest) (*pbAdmin.ListDoctorsResponse, error) {
	return s.ListDoctors(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateDoctor(ctx context.Context, req *pbAdmin.UpdateDoctorRequest) (*pbAdmin.UpdateDoctorResponse, error) {
	return s.UpdateDoctor(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteDoctor(ctx context.Context, req *pbAdmin.DeletePersonalRequest) (*pbAdmin.DeletePersonalResponse, error) {
	return s.DeleteDoctor(ctx, req)
}

// ------------------- Bed -------------------

func (s *ServiceRepositoryClient) CreateBed(ctx context.Context, req *pbAdmin.CreateBedRequest) (*pbAdmin.CreateBedResponse, error) {
	return s.CreateBed(ctx, req)
}

func (s *ServiceRepositoryClient) GetBedByID(ctx context.Context, req *pbAdmin.GetBedByIDRequest) (*pbAdmin.GetBedByIDResponse, error) {
	return s.GetBedByID(ctx, req)
}

func (s *ServiceRepositoryClient) ListBedS(ctx context.Context, req *pbAdmin.ListBedSRequest) (*pbAdmin.ListBedSResponse, error) {
	return s.ListBedS(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateBed(ctx context.Context, req *pbAdmin.UpdateBedRequest) (*pbAdmin.UpdateBedResponse, error) {
	return s.UpdateBed(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteBed(ctx context.Context, req *pbAdmin.DeleteBedRequest) (*pbAdmin.DeleteBedResponse, error) {
	return s.DeleteBed(ctx, req)
}

//! Doctor Service
// ------------------- Appointment -------------------

func (s *ServiceRepositoryClient) CreateAppointment(ctx context.Context, req *pbDoctor.CreateAppointmentRequest) (*pbDoctor.CreateAppointmentResponse, error) {
	return s.doctorClient.CreateAppointment(ctx, req)
}

func (s *ServiceRepositoryClient) GetAppointmentById(ctx context.Context, req *pbDoctor.GetAppointmentByIdRequest) (*pbDoctor.GetAppointmentByIdResponse, error) {
	return s.doctorClient.GetAppointmentById(ctx, req)
}

func (s *ServiceRepositoryClient) ListAppointments(ctx context.Context, req *pbDoctor.ListAppointmentsRequest) (*pbDoctor.ListAppointmentsResponse, error) {
	return s.doctorClient.ListAppointments(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateAppointment(ctx context.Context, req *pbDoctor.UpdateAppointmentRequest) (*pbDoctor.UpdateAppointmentResponse, error) {
	return s.doctorClient.UpdateAppointment(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteAppointment(ctx context.Context, req *pbDoctor.DeleteAppointmentRequest) (*pbDoctor.DeleteAppointmentResponse, error) {
	return s.doctorClient.DeleteAppointment(ctx, req)
}

// ------------------- Prescription -------------------

func (s *ServiceRepositoryClient) CreatePrescription(ctx context.Context, req *pbDoctor.CreatePrescriptionRequest) (*pbDoctor.CreatePrescriptionResponse, error) {
	return s.doctorClient.CreatePrescription(ctx, req)
}

func (s *ServiceRepositoryClient) GetPrescriptionById(ctx context.Context, req *pbDoctor.GetPrescriptionByIdRequest) (*pbDoctor.GetPrescriptionByIdResponse, error) {
	return s.doctorClient.GetPrescriptionById(ctx, req)
}

func (s *ServiceRepositoryClient) ListPrescriptions(ctx context.Context, req *pbDoctor.ListPrescriptionsRequest) (*pbDoctor.ListPrescriptionsResponse, error) {
	return s.doctorClient.ListPrescriptions(ctx, req)
}

func (s *ServiceRepositoryClient) UpdatePrescription(ctx context.Context, req *pbDoctor.UpdatePrescriptionRequest) (*pbDoctor.UpdatePrescriptionResponse, error) {
	return s.doctorClient.UpdatePrescription(ctx, req)
}

func (s *ServiceRepositoryClient) DeletePrescription(ctx context.Context, req *pbDoctor.DeletePrescriptionRequest) (*pbDoctor.DeletePrescriptionResponse, error) {
	return s.doctorClient.DeletePrescription(ctx, req)
}

//! Nurse Service
// ------------------- Donor -------------------

func (s *ServiceRepositoryClient) CreateDonor(ctx context.Context, req *pbNurse.CreateDonorRequest) (*pbNurse.CreateDonorResponse, error) {
	return s.nurseClient.CreateDonor(ctx, req)
}

func (s *ServiceRepositoryClient) GetDonorById(ctx context.Context, req *pbNurse.GetDonorByIdRequest) (*pbNurse.GetDonorByIdResponse, error) {
	return s.nurseClient.GetDonorById(ctx, req)
}

func (s *ServiceRepositoryClient) ListDonors(ctx context.Context, req *pbNurse.ListDonorsRequest) (*pbNurse.ListDonorsResponse, error) {
	return s.nurseClient.ListDonors(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateDonor(ctx context.Context, req *pbNurse.UpdateDonorRequest) (*pbNurse.UpdateDonorResponse, error) {
	return s.nurseClient.UpdateDonor(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteDonor(ctx context.Context, req *pbNurse.DeleteDonorRequest) (*pbNurse.DeleteDonorResponse, error) {
	return s.nurseClient.DeleteDonor(ctx, req)
}

//! Patient Service 
// ------------------- Patient -------------------

func (s *ServiceRepositoryClient) CreatePatient(ctx context.Context, req *pbPatient.CreatePatientRequest) (*pbPatient.CreatePatientResponse, error) {
	return s.patientClient.CreatePatient(ctx, req)
}

func (s *ServiceRepositoryClient) GetPatientById(ctx context.Context, req *pbPatient.GetPatientByIdRequest) (*pbPatient.GetPatientByIdResponse, error) {
	return s.patientClient.GetPatientById(ctx, req)
}

func (s *ServiceRepositoryClient) ListPatients(ctx context.Context, req *pbPatient.ListPatientsRequest) (*pbPatient.ListPatientsResponse, error) {
	return s.patientClient.ListPatient(ctx, req)
}

func (s *ServiceRepositoryClient) UpdatePatient(ctx context.Context, req *pbPatient.UpdatePatientRequest) (*pbPatient.UpdatePatientResponse, error) {
	return s.patientClient.UpdatePatient(ctx, req)
}

func (s *ServiceRepositoryClient) DeletePatient(ctx context.Context, req *pbPatient.DeletePatientRequest) (*pbPatient.DeletePatientResponse, error) {
	return s.patientClient.DeletePatient(ctx, req)
}

//! Pharmacist Service
// ------------------- Medicine -------------------

func (s *ServiceRepositoryClient) CreateMedicine(ctx context.Context, req *pbPharmacist.CreateMedicineRequest) (*pbPharmacist.CreateMedicineResponse, error) {
	return s.pharmacistClient.CreateMedicine(ctx, req)
}

func (s *ServiceRepositoryClient) GetMedicineById(ctx context.Context, req *pbPharmacist.GetMedicineByIdRequest) (*pbPharmacist.GetMedicineByIdResponse, error) {
	return s.pharmacistClient.GetMedicineById(ctx, req)
}

func (s *ServiceRepositoryClient) ListMedicines(ctx context.Context, req *pbPharmacist.ListMedicinesRequest) (*pbPharmacist.ListMedicinesResponse, error) {
	return s.pharmacistClient.ListMedicines(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateMedicine(ctx context.Context, req *pbPharmacist.UpdateMedicineRequest) (*pbPharmacist.UpdateMedicineResponse, error) {
	return s.pharmacistClient.UpdateMedicine(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteMedicine(ctx context.Context, req *pbPharmacist.DeleteMedicineRequest) (*pbPharmacist.DeleteMedicineResponse, error) {
	return s.pharmacistClient.DeleteMedicine(ctx, req)
}
