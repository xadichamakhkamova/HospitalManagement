package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	
)

// @Router /pharmacist/medicines [post]
// @Summary CREATE MEDICINE
// @Security BearerAuth
// @Description This method creates medicine
// @Tags MEDICINES
// @Accept json
// @Produce json
// @Param medicine body models.CreateMedicineRequest true "Medicine"
// @Success 200 {object} models.CreateMedicineResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateMedicine(c *gin.Context) {

	req := pb.CreateMedicineRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateMedicine(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /pharmacist/medicines/{id} [get]
// @Summary GET MEDICINE BY ID
// @Security BearerAuth
// @Description This method gets medicine by id
// @Tags MEDICINES
// @Accept json
// @Produce json
// @Param id path string true "Medicine Id"
// @Success 200 {object} models.GetMedicineByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetMedicineById(c *gin.Context) {

	req := pb.GetMedicineByIdRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetMedicineById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /pharmacist/medicines [get]
// @Summary GET MEDICINES LIST
// @Security BearerAuth
// @Description This method gets medicines list by filter
// @Tags MEDICINES
// @Accept json
// @Produce json
// @Param search query string false "Search by name/category/company"
// @Param status query string false "Filter by status"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} models.ListMedicinesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListMedicines(c *gin.Context) {

	req := pb.ListMedicinesRequest{}
	req.Search = c.Query("search")

	if c.Query("status") != "" {
		req.Status = pb.MedicineStatus(pb.MedicineStatus_value[c.Query("status")])
	}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListMedicines(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /pharmacist/medicines/{id} [put]
// @Summary UPDATE MEDICINE
// @Security BearerAuth
// @Description This method updates medicine
// @Tags MEDICINES
// @Accept json
// @Produce json
// @Param id path string true "Medicine Id"
// @Param medicine body models.UpdateMedicineRequest true "Medicine"
// @Success 200 {object} models.UpdateMedicineResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateMedicine(c *gin.Context) {

	req := pb.UpdateMedicineRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateMedicine(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /pharmacist/medicines/{id} [delete]
// @Summary DELETE MEDICINE
// @Security BearerAuth
// @Description This method deletes medicine
// @Tags MEDICINES
// @Accept json
// @Produce json
// @Param id path string true "Medicine Id"
// @Success 200 {object} models.DeleteMedicineResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteMedicine(c *gin.Context) {
	
	req := pb.DeleteMedicineRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeleteMedicine(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
