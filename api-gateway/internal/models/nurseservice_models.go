package models

import "time"

type HealthConditionType string

const (
	Healthy       HealthConditionType = "healthy"
	MinorIllness  HealthConditionType = "minor_illness"
	Chronic       HealthConditionType = "chronic"
	NotEligible   HealthConditionType = "not_eligible"
)

type PatientDonor struct {
	ID               string              `json:"id" example:"uuid"`
	FullName         string              `json:"full_name" example:"John Doe"`
	Email            string              `json:"email" example:"john@example.com"`
	Password         string              `json:"password" example:"hashed_password"`
	Address          string              `json:"address" example:"123 Main Street, City"`
	PhoneNumber      string              `json:"phone_number" example:"+998901234567"`
	Gender           string              `json:"gender" example:"Male"`
	BirthDate        string              `json:"birth_date" example:"1990-01-01"`
	BloodGroup       string              `json:"blood_group" example:"O+"`
	LastDonation     time.Time           `json:"last_donation" example:"2025-09-01T12:00:00Z"`
	DonationCount    int32               `json:"donation_count" example:"5"`
	IsEligible       bool                `json:"is_eligible" example:"true"`
	LastCheckupDate  time.Time           `json:"last_checkup_date" example:"2025-08-20T12:00:00Z"`
	Weight           float64             `json:"weight" example:"72.5"`
	HealthConditions HealthConditionType `json:"health_conditions" example:"healthy"`
	DonationLocation string              `json:"donation_location" example:"City Hospital"`
	Timestamps       Timestamps3         `json:"timestamps"`
}

type CreateDonorRequest struct {
	FullName         string              `json:"full_name" example:"John Doe"`
	Email            string              `json:"email" example:"john@example.com"`
	Password         string              `json:"password" example:"secret123"`
	Address          string              `json:"address" example:"123 Street, City"`
	PhoneNumber      string              `json:"phone_number" example:"+998901234567"`
	Gender           string              `json:"gender" example:"Male"`
	BirthDate        string              `json:"birth_date" example:"1990-01-01"`
	BloodGroup       string              `json:"blood_group" example:"O+"`
	Weight           float64             `json:"weight" example:"72.5"`
	HealthConditions HealthConditionType `json:"health_conditions" example:"healthy"`
}

type CreateDonorResponse struct {
	Donor PatientDonor `json:"donor"`
}

type GetDonorByIdRequest struct {
	ID string `json:"id" example:"uuid"`
}

type GetDonorByIdResponse struct {
	Donor PatientDonor `json:"donor"`
}

type ListDonorsRequest struct {
	Search       string `json:"search" example:"john"`
	Gender       string `json:"gender" example:"Male"`
	BloodGroup   string `json:"blood_group" example:"O+"`
	OnlyEligible bool   `json:"only_eligible" example:"true"`
	Page         int32  `json:"page" example:"1"`
	Limit        int32  `json:"limit" example:"10"`
}

type ListDonorsResponse struct {
	Patients   []PatientDonor `json:"patients"`
	TotalCount int32          `json:"total_count"`
}

type UpdateDonorRequest struct {
	ID              string              `json:"id" example:"uuid"`
	FullName        string              `json:"full_name" example:"John Doe"`
	Email           string              `json:"email" example:"john@example.com"`
	Password        string              `json:"password" example:"newSecret123"`
	Address         string              `json:"address" example:"456 New Street, City"`
	PhoneNumber     string              `json:"phone_number" example:"+998901234567"`
	Gender          string              `json:"gender" example:"Male"`
	BirthDate       string              `json:"birth_date" example:"1990-01-01"`
	BloodGroup      string              `json:"blood_group" example:"O+"`
	Weight          float64             `json:"weight" example:"75.0"`
	HealthCondition HealthConditionType `json:"health_conditions" example:"minor_illness"`
}

type UpdateDonorResponse struct {
	Donor PatientDonor `json:"donor"`
}

type DeleteDonorRequest struct {
	ID string `json:"id" example:"uuid"`
}

type DeleteDonorResponse struct {
	Status int64 `json:"status" example:"204"`
}

type RegisterDonationRequest struct {
	DonorID          string `json:"donor_id" example:"uuid"`
	DonationLocation string `json:"donation_location" example:"City Hospital"`
}

type RegisterDonationResponse struct {
	LastDonation  time.Time `json:"last_donation" example:"2025-09-01T12:00:00Z"`
	DonationCount int32     `json:"donation_count" example:"5"`
	IsEligible    bool      `json:"is_eligible" example:"true"`
}

type RegisterCheckupRequest struct {
	DonorID string `json:"donor_id" example:"uuid"`
}

type RegisterCheckupResponse struct {
	LastCheckupDate time.Time `json:"last_checkup_date" example:"2025-08-20T12:00:00Z"`
	IsEligible      bool      `json:"is_eligible" example:"true"`
}

type Timestamps3 struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
