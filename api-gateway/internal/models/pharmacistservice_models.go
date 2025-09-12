package models 

import "time"

type MedicineCategory string
const (
	MedicineCategoryUnspecified  MedicineCategory = "MEDICINE_CATEGORY_UNSPECIFIED"
	MedicineCategoryAntibiotic   MedicineCategory = "MEDICINE_CATEGORY_ANTIBIOTIC"
	MedicineCategoryAnalgesic    MedicineCategory = "MEDICINE_CATEGORY_ANALGESIC"
	MedicineCategoryAntiviral    MedicineCategory = "MEDICINE_CATEGORY_ANTIVIRAL"
	MedicineCategoryVitamins     MedicineCategory = "MEDICINE_CATEGORY_VITAMINS"
	MedicineCategoryAntifungal   MedicineCategory = "MEDICINE_CATEGORY_ANTIFUNGAL"
	MedicineCategoryVaccine      MedicineCategory = "MEDICINE_CATEGORY_VACCINE"
	MedicineCategoryOther        MedicineCategory = "MEDICINE_CATEGORY_OTHER"
)

type MedicineStatus string
const (
	MedicineStatusUnspecified   MedicineStatus = "MEDICINE_STATUS_UNSPECIFIED"
	MedicineStatusAvailable     MedicineStatus = "MEDICINE_STATUS_AVAILABLE"
	MedicineStatusOutOfStock    MedicineStatus = "MEDICINE_STATUS_OUT_OF_STOCK"
	MedicineStatusExpired       MedicineStatus = "MEDICINE_STATUS_EXPIRED"
	MedicineStatusDiscontinued  MedicineStatus = "MEDICINE_STATUS_DISCONTINUED"
)

type Medicine struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Category    MedicineCategory `json:"category"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	Company     string           `json:"company"`
	Status      MedicineStatus   `json:"status"`
	Timestamps  Timestamps5      `json:"timestamps"`
}

type CreateMedicineRequest struct {
	Name        string           `json:"name"`
	Category    MedicineCategory `json:"category"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	Company     string           `json:"company"`
	Status      MedicineStatus   `json:"status"`
}
type CreateMedicineResponse struct {
	Medicine Medicine `json:"medicine"`
}

type GetMedicineByIdRequest struct {
	ID string `json:"id"`
}
type GetMedicineByIdResponse struct {
	Medicine Medicine `json:"medicine"`
}

type ListMedicinesRequest struct {
	Search string         `json:"search"` // name, category, company
	Status MedicineStatus `json:"status"` // optional filter
	Page   int32          `json:"page"`
	Limit  int32          `json:"limit"`
}
type ListMedicinesResponse struct {
	Medicines  []Medicine `json:"medicines"`
	TotalCount int32      `json:"total_count"`
}

type UpdateMedicineRequest struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Category    MedicineCategory `json:"category"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	Company     string           `json:"company"`
	Status      MedicineStatus   `json:"status"`
}
type UpdateMedicineResponse struct {
	Medicine Medicine `json:"medicine"`
}

type DeleteMedicineRequest struct {
	ID string `json:"id"`
}
type DeleteMedicineResponse struct {
	Status int64 `json:"status"` // 204 = deleted
}

type Timestamps5 struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
