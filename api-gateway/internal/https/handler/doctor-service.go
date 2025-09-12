package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// @Router /doctors/appointments [post]
// @Summary CREATE APPOINTMENT
// @Security BearerAuth
// @Description This method creates appointment
// @Tags APPOINTMENTS
// @Accept json
// @Produce json
// @Param appointment body models.CreateAppointmentRequest true "Appointment"
// @Success 200 {object} models.CreateAppointmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateAppointment(c *gin.Context) {

	req := pb.CreateAppointmentRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateAppointment(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /doctors/appointments/{id} [get]
// @Summary GET APPOINTMENT BY ID
// @Security BearerAuth
// @Description This method gets appointment by id
// @Tags APPOINTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Appointment Id"
// @Success 200 {object} models.GetAppointmentByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetAppointmentById(c *gin.Context) {

	req := pb.GetAppointmentByIdRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetAppointmentById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /doctors/appointments [get]
// @Summary GET APPOINTMENTS LIST
// @Security BearerAuth
// @Description This method gets appointments list
// @Tags APPOINTMENTS
// @Accept json
// @Produce json
// @Param date query string false "Appointment date (yyyy-mm-dd)"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.ListAppointmentsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListAppointments(c *gin.Context) {

	req := pb.ListAppointmentsRequest{}

	if c.Query("date") != "" {
		parsedDate, err := time.Parse(time.RFC3339, c.Query("date"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid date format, use RFC3339 (e.g. 2025-09-12T15:04:05Z)"})
			return
		}
		req.Date = timestamppb.New(parsedDate)
	}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListAppointments(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /doctors/appointments/{id} [put]
// @Summary UPDATE APPOINTMENT
// @Security BearerAuth
// @Description This method updates appointment
// @Tags APPOINTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Appointment Id"
// @Param appointment body models.UpdateAppointmentRequest true "Appointment"
// @Success 200 {object} models.UpdateAppointmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateAppointment(c *gin.Context) {

	req := pb.UpdateAppointmentRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateAppointment(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /doctors/appointments/{id} [delete]
// @Summary DELETE APPOINTMENT
// @Security BearerAuth
// @Description This method deletes appointment
// @Tags APPOINTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Appointment Id"
// @Success 200 {object} models.DeleteAppointmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteAppointment(c *gin.Context) {

	req := pb.DeleteAppointmentRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeleteAppointment(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/prescriptions [post]
// @Summary CREATE PRESCRIPTION
// @Security BearerAuth
// @Description This method creates prescription
// @Tags PRESCRIPTIONS
// @Accept json
// @Produce json
// @Param prescription body models.CreatePrescriptionRequest true "Prescription"
// @Success 200 {object} models.CreatePrescriptionResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreatePrescription(c *gin.Context) {

	req := pb.CreatePrescriptionRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.CreatePrescription(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/prescriptions/{id} [get]
// @Summary GET PRESCRIPTION BY ID
// @Security BearerAuth
// @Description This method gets prescription by id
// @Tags PRESCRIPTIONS
// @Accept json
// @Produce json
// @Param id path string true "Prescription Id"
// @Success 200 {object} models.GetPrescriptionByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetPrescriptionById(c *gin.Context) {

	req := pb.GetPrescriptionByIdRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.GetPrescriptionById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/prescriptions [get]
// @Summary GET PRESCRIPTIONS LIST
// @Security BearerAuth
// @Description This method gets prescriptions list
// @Tags PRESCRIPTIONS
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} models.ListPrescriptionsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListPrescriptions(c *gin.Context) {

	req := pb.ListPrescriptionsRequest{}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListPrescriptions(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/prescriptions/{id} [put]
// @Summary UPDATE PRESCRIPTION
// @Security BearerAuth
// @Description This method updates prescription
// @Tags PRESCRIPTIONS
// @Accept json
// @Produce json
// @Param id path string true "Prescription id"
// @Param prescription body models.UpdatePrescriptionRequest true "Prescription"
// @Success 200 {object} models.UpdatePrescriptionResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdatePrescription(c *gin.Context) {

	req := pb.UpdatePrescriptionRequest{}
	req.Presc.Id = c.Param("id")

	if err := c.BindJSON(&req.Presc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdatePrescription(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/prescriptions/{id} [delete]
// @Summary DELETE PRESCRIPTION
// @Security BearerAuth
// @Description This method deletes prescription
// @Tags PRESCRIPTIONS
// @Accept json
// @Produce json
// @Param id path string true "Prescription Id"
// @Success 200 {object} models.DeletePrescriptionResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeletePrescription(c *gin.Context) {
	
	req := pb.DeletePrescriptionRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.DeletePrescription(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
