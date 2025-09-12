package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"
)

// @Router /nurse/donors [post]
// @Summary CREATE DONOR
// @Security BearerAuth
// @Description This method creates donor
// @Tags DONORS
// @Accept json
// @Produce json
// @Param donor body models.CreateDonorRequest true "Donor"
// @Success 200 {object} models.CreateDonorResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateDonor(c *gin.Context) {

	req := pb.CreateDonorRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateDonor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /nurse/donors/{id} [get]
// @Summary GET DONOR BY ID
// @Security BearerAuth
// @Description This method gets donor by id
// @Tags DONORS
// @Accept json
// @Produce json
// @Param id path string true "Donor Id"
// @Success 200 {object} models.GetDonorByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDonorById(c *gin.Context) {

	req := pb.GetDonorByIdRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetDonorById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /nurse/donors [get]
// @Summary GET DONORS LIST
// @Security BearerAuth
// @Description This method gets donors list by filter
// @Tags DONORS
// @Accept json
// @Produce json
// @Param search query string false "Search by name/email"
// @Param gender query string false "Filter by gender"
// @Param blood_group query string false "Filter by blood group"
// @Param only_eligible query bool false "Only eligible donors"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} models.ListDonorsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListDonors(c *gin.Context) {

	req := pb.ListDonorsRequest{}
	req.Search = c.Query("search")
	req.Gender = c.Query("gender")
	req.BloodGroup = c.Query("blood_group")

	if c.Query("only_eligible") == "true" {
		req.OnlyEligible = true
	}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListDonors(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /nurse/donors/{id} [put]
// @Summary UPDATE DONOR
// @Security BearerAuth
// @Description This method updates donor
// @Tags DONORS
// @Accept json
// @Produce json
// @Param id path string true "Donor Id"
// @Param donor body models.UpdateDonorRequest true "Donor"
// @Success 200 {object} models.UpdateDonorResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateDonor(c *gin.Context) {

	req := pb.UpdateDonorRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateDonor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /nurse/donors/{id} [delete]
// @Summary DELETE DONOR
// @Security BearerAuth
// @Description This method deletes donor
// @Tags DONORS
// @Accept json
// @Produce json
// @Param id path string true "Donor Id"
// @Success 200 {object} models.DeleteDonorResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDonor(c *gin.Context) {
	
	req := pb.DeleteDonorRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeleteDonor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
