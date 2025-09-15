package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
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

	h.log.Info("CreateAppointment: request received")

	req := pb.CreateAppointmentRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.WithError(err).Warn("CreateAppointment: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateAppointment(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("CreateAppointment: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("appointment_id", resp.Appointment.Id).Info("CreateAppointment: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("GetAppointmentById: request received")

	req := pb.GetAppointmentByIdRequest{Id: id}
	resp, err := h.service.GetAppointmentById(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("GetAppointmentById: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("appointment_id", resp.Appointment.Id).Info("GetAppointmentById: success")
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

	dateStr := c.Query("date")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	h.log.WithFields(logrus.Fields{
		"date":  dateStr,
		"page":  page,
		"limit": limit,
	}).Info("ListAppointments: request received")

	req := pb.ListAppointmentsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}
	if dateStr != "" {
		parsedDate, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			h.log.WithError(err).Warn("ListAppointments: invalid date format")
			c.JSON(400, gin.H{"error": "invalid date format, use RFC3339 (e.g. 2025-09-12T15:04:05Z)"})
			return
		}
		req.Date = parsedDate.String()
	}

	resp, err := h.service.ListAppointments(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("ListAppointments: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("count", len(resp.Appointment)).Info("ListAppointments: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("UpdateAppointment: request received")

	req := pb.UpdateAppointmentRequest{Id: id}
	if err := c.BindJSON(&req); err != nil {
		h.log.WithError(err).Warn("UpdateAppointment: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateAppointment(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("UpdateAppointment: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("appointment_id", resp.Appointment.Id).Info("UpdateAppointment: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("DeleteAppointment: request received")

	req := pb.DeleteAppointmentRequest{Id: id}
	resp, err := h.service.DeleteAppointment(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("DeleteAppointment: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("appointment_id", id).Info("DeleteAppointment: success")
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

	h.log.Info("CreatePrescription: request received")

	req := pb.CreatePrescriptionRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.WithError(err).Warn("CreatePrescription: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreatePrescription(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("CreatePrescription: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("prescription_id", resp.Presc.Id).Info("CreatePrescription: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("GetPrescriptionById: request received")

	req := pb.GetPrescriptionByIdRequest{Id: id}
	resp, err := h.service.GetPrescriptionById(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("GetPrescriptionById: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("prescription_id", resp.Presc.Id).Info("GetPrescriptionById: success")
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

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	h.log.WithFields(logrus.Fields{
		"page":  page,
		"limit": limit,
	}).Info("ListPrescriptions: request received")

	req := pb.ListPrescriptionsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := h.service.ListPrescriptions(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("ListPrescriptions: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("count", len(resp.Presc)).Info("ListPrescriptions: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("UpdatePrescription: request received")

	req := pb.UpdatePrescriptionRequest{}
	req.Presc.Id = id

	if err := c.BindJSON(&req.Presc); err != nil {
		h.log.WithError(err).Warn("UpdatePrescription: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdatePrescription(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("UpdatePrescription: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("prescription_id", resp.Presc.Id).Info("UpdatePrescription: success")
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

	id := c.Param("id")
	h.log.WithField("id", id).Info("DeletePrescription: request received")

	req := pb.DeletePrescriptionRequest{Id: id}
	resp, err := h.service.DeletePrescription(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("DeletePrescription: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("prescription_id", id).Info("DeletePrescription: success")
	c.JSON(200, resp)
}
