package models 

import "time"

type GenderType string
const (
	GenderUnspecified GenderType = "GENDER_TYPE_UNSPECIFIED"
	GenderMale        GenderType = "MALE"
	GenderFemale      GenderType = "FEMALE"
)

type BloodType string
const (
	BloodUnspecified BloodType = "BLOOD_TYPE_UNSPECIFIED"
	APositive        BloodType = "A_POSITIVE"
	ANegative        BloodType = "A_NEGATIVE"
	BPositive        BloodType = "B_POSITIVE"
	BNegative        BloodType = "B_NEGATIVE"
	ABPositive       BloodType = "AB_POSITIVE"
	ABNegative       BloodType = "AB_NEGATIVE"
	OPositive        BloodType = "O_POSITIVE"
	ONegative        BloodType = "O_NEGATIVE"
)

type Patient struct {
	ID         string      `json:"id"`
	FullName   string      `json:"full_name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone_number"`
	Gender     GenderType  `json:"gender"`
	BirthDate  string      `json:"birth_date"`
	BloodGroup BloodType   `json:"blood_group"`
	Timestamps Timestamps4 `json:"timestamps"`
}

type CreatePatientRequest struct {
	FullName   string     `json:"full_name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Address    string     `json:"address"`
	Phone      string     `json:"phone_number"`
	Gender     GenderType `json:"gender"`
	BirthDate  string     `json:"birth_date"`
	BloodGroup BloodType  `json:"blood_group"`
}
type CreatePatientResponse struct {
	Patient Patient `json:"patient"`
}

type GetPatientByIdRequest struct {
	ID string `json:"id"`
}
type GetPatientByIdResponse struct {
	Patient Patient `json:"patient"`
}

type ListPatientsRequest struct {
	Search string `json:"search"` // name, email, gender
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}
type ListPatientsResponse struct {
	Patients   []Patient `json:"patients"`
	TotalCount int32     `json:"total_count"`
}

type UpdatePatientRequest struct {
	ID         string      `json:"id"`
	FullName   string      `json:"full_name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone_number"`
	Gender     GenderType  `json:"gender"`
	BirthDate  string      `json:"birth_date"`
	BloodGroup BloodType   `json:"blood_group"`
	Timestamps Timestamps4 `json:"timestamps"`
}
type UpdatePatientResponse struct {
	Patient Patient `json:"patient"`
}

type DeletePatientRequest struct {
	ID string `json:"id"`
}
type DeletePatientResponse struct {
	Status int64 `json:"status"` // 204 = deleted
}


type Timestamps4 struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
