package models

import "time"


type Timestamps1 struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Department struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Number      int64       `json:"number"`
	Description string      `json:"description"`
	Timestamps  Timestamps1 `json:"timestamps"`
}

type Personal struct {
	ID          string      `json:"id"`
	Profession  string      `json:"profession"`
	FullName    string      `json:"full_name"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	Address     string      `json:"address"`
	PhoneNumber string      `json:"phone_number"`
	Timestamps  Timestamps1 `json:"timestamps"`
}

type Doctor struct {
	Info             Personal    `json:"info"`
	PersonalID       string      `json:"personal_id"`
	DepartmentNumber int64       `json:"department_number"`
	Timestamps       Timestamps1 `json:"timestamps"`
}


// -------- Department --------
type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Number      int64  `json:"number"`
	Description string `json:"description"`
}
type CreateDepartmentResponse struct {
	Department Department `json:"department"`
}

type GetDepartmentByIdRequest struct {
	ID string `json:"id"`
}
type GetDepartmentByIdResponse struct {
	Department Department `json:"department"`
}

type ListDepartmentsRequest struct {
	Search string `json:"search"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}
type ListDepartmentsResponse struct {
	Departments []Department `json:"departments"`
	TotalCount  int32        `json:"total_count"`
}

type UpdateDepartmentRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Number      int64  `json:"number"`
	Description string `json:"description"`
}
type UpdateDepartmentResponse struct {
	Department Department `json:"department"`
}

type DeleteDepartmentRequest struct {
	ID string `json:"id"`
}
type DeleteDepartmentResponse struct {
	Status int64 `json:"status"` // 204 = deleted
}

// -------- Personal & Doctor --------
type CreatePersonalRequest struct {
	Profession  string `json:"profession"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
type CreatePersonalResponse struct {
	Personal Personal `json:"personal"`
}

type CreateDoctorRequest struct {
	PersonalID       string `json:"personal_id"`
	DepartmentNumber int64  `json:"department_number"`
}
type CreateDoctorResponse struct {
	ID               string      `json:"id"`
	PersonalID       string      `json:"personal_id"`
	DepartmentNumber int64       `json:"department_number"`
	Timestamps       Timestamps1 `json:"timestamps"`
}

type GetPersonalByIdRequest struct {
	ID string `json:"id"`
}
type GetPersonalByIdResponse struct {
	Personal Personal `json:"personal"`
}

type ListPersonalsRequest struct {
	Search string `json:"search"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}
type ListPersonalsResponse struct {
	Personals  []Personal `json:"personals"`
	TotalCount int32      `json:"total_count"`
}

type GetDoctorByIdResponse struct {
	Doctor Doctor `json:"doctor"`
}
type ListDoctorsResponse struct {
	Doctors    []Doctor `json:"doctors"`
	TotalCount int32    `json:"total_count"`
}

type UpdatePersonalRequest struct {
	ID          string `json:"id"`
	Profession  string `json:"profession"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
type UpdatePersonalResponse struct {
	Personal Personal `json:"personal"`
}

type UpdateDoctorRequest struct {
	ID               string `json:"id"`
	DepartmentNumber int64  `json:"department_number"`
}
type UpdateDoctorResponse struct {
	ID               string      `json:"id"`
	PersonalID       string      `json:"personal_id"`
	DepartmentNumber int64       `json:"department_number"`
	Timestamps       Timestamps1 `json:"timestamps"`
}

type DeletePersonalRequest struct {
	ID string `json:"id"`
}
type DeletePersonalResponse struct {
	Status int64 `json:"status"`
}
