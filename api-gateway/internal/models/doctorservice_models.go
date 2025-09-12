package models

import "time"

type Appointment struct {
	ID        string      `json:"id"`
	DoctorID  string      `json:"doctor_id"`
	PatientID string      `json:"patient_id"`
	Date      time.Time   `json:"date"`
	Timestamps Timestamps2 `json:"timestamps"`
}

type Prescription struct {
	ID          string      `json:"id"`
	DoctorID    string      `json:"doctor_id"`
	PatientID   string      `json:"patient_id"`
	CaseHistory string      `json:"case_history"`
	Medication  string      `json:"medication"`
	Description string      `json:"description"`
	Timestamps  Timestamps2 `json:"timestamps"`
}

type CreateAppointmentRequest struct {
	DoctorID  string    `json:"doctor_id"`
	PatientID string    `json:"patient_id"`
	Date      time.Time `json:"date"`
}
type CreateAppointmentResponse struct {
	Appointment Appointment `json:"appointment"`
}

type GetAppointmentByIdRequest struct {
	ID string `json:"id"`
}
type GetAppointmentByIdResponse struct {
	Appointment Appointment `json:"appointment"`
}

type ListAppointmentsRequest struct {
	Date  time.Time `json:"date"`
	Page  int32     `json:"page"`
	Limit int32     `json:"limit"`
}
type ListAppointmentsResponse struct {
	Appointments []Appointment `json:"appointments"`
	TotalCount   int32         `json:"total_count"`
}

type UpdateAppointmentRequest struct {
	ID        string    `json:"id"`
	DoctorID  string    `json:"doctor_id"`
	PatientID string    `json:"patient_id"`
	Date      time.Time `json:"date"`
}
type UpdateAppointmentResponse struct {
	Appointment Appointment `json:"appointment"`
}

type DeleteAppointmentRequest struct {
	ID string `json:"id"`
}
type DeleteAppointmentResponse struct {
	Status int64 `json:"status"` // 204 = deleted
}

// -------- Prescription --------
type CreatePrescriptionRequest struct {
	DoctorID    string `json:"doctor_id"`
	PatientID   string `json:"patient_id"`
	CaseHistory string `json:"case_history"`
	Medication  string `json:"medication"`
	Description string `json:"description"`
}
type CreatePrescriptionResponse struct {
	Prescription Prescription `json:"prescription"`
}

type GetPrescriptionByIdRequest struct {
	ID string `json:"id"`
}
type GetPrescriptionByIdResponse struct {
	Prescription Prescription `json:"prescription"`
}

type ListPrescriptionsRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}
type ListPrescriptionsResponse struct {
	Prescriptions []Prescription `json:"prescriptions"`
	TotalCount    int32          `json:"total_count"`
}

type UpdatePrescriptionRequest struct {
	Prescription Prescription `json:"prescription"`
}
type UpdatePrescriptionResponse struct {
	Prescription Prescription `json:"prescription"`
}

type DeletePrescriptionRequest struct {
	ID string `json:"id"`
}
type DeletePrescriptionResponse struct {
	Status int64 `json:"status"` // 204 = deleted
}

type Timestamps2 struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
