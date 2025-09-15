package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
)

// @Router /patient/patients [post]
// @Summary CREATE PATIENT
// @Security BearerAuth
// @Description This method creates patient
// @Tags PATIENTS
// @Accept json
// @Produce json
// @Param patient body models.CreatePatientRequest true "Patient"
// @Success 200 {object} models.CreatePatientResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreatePatient(c *gin.Context) {

	h.log.Info("CreatePatient: request received")

	req := pb.CreatePatientRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("CreatePatient: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreatePatient(context.Background(), &req)
	if err != nil {
		h.log.Error("CreatePatient: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("CreatePatient: success, id=" + resp.Patient.Id)
	c.JSON(200, resp)
}

// @Router /patient/patients/{id} [get]
// @Summary GET PATIENT BY ID
// @Security BearerAuth
// @Description This method gets patient by id
// @Tags PATIENTS
// @Accept json
// @Produce json
// @Param id path string true "Patient Id"
// @Success 200 {object} models.GetPatientByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetPatientById(c *gin.Context) {

	id := c.Param("id")
	h.log.Info("GetPatientById: request id=" + id)

	req := pb.GetPatientByIdRequest{Id: id}
	resp, err := h.service.GetPatientById(context.Background(), &req)
	if err != nil {
		h.log.Error("GetPatientById: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("GetPatientById: success, id=" + resp.Patient.Id)
	c.JSON(200, resp)
}


// @Router /patient/patients [get]
// @Summary GET PATIENTS LIST
// @Security BearerAuth
// @Description This method gets patients list by filter
// @Tags PATIENTS
// @Accept json
// @Produce json
// @Param search query string false "Search by name/email/phone"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} models.ListPatientsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListPatients(c *gin.Context) {

	h.log.Info("ListPatients: request received")

	req := pb.ListPatientsRequest{
		Search: c.Query("search"),
	}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListPatients(context.Background(), &req)
	if err != nil {
		h.log.Error("ListPatients: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("ListPatients: success, count=%d", len(resp.Patients))
	c.JSON(200, resp)
}

// @Router /patient/patients/{id} [put]
// @Summary UPDATE PATIENT
// @Security BearerAuth
// @Description This method updates patient
// @Tags PATIENTS
// @Accept json
// @Produce json
// @Param id path string true "Patient Id"
// @Param patient body models.UpdatePatientRequest true "Patient"
// @Success 200 {object} models.UpdatePatientResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdatePatient(c *gin.Context) {

	id := c.Param("id")
	h.log.Info("UpdatePatient: request id=" + id)

	req := pb.UpdatePatientRequest{Id: id}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("UpdatePatient: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdatePatient(context.Background(), &req)
	if err != nil {
		h.log.Error("UpdatePatient: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("UpdatePatient: success, id=" + resp.Patient.Id)
	c.JSON(200, resp)
}

// @Router /patient/patients/{id} [delete]
// @Summary DELETE PATIENT
// @Security BearerAuth
// @Description This method deletes patient
// @Tags PATIENTS
// @Accept json
// @Produce json
// @Param id path string true "Patient Id"
// @Success 200 {object} models.DeletePatientResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeletePatient(c *gin.Context) {
	
	id := c.Param("id")
	h.log.Info("DeletePatient: request id=" + id)

	req := pb.DeletePatientRequest{Id: id}
	_, err := h.service.DeletePatient(context.Background(), &req)
	if err != nil {
		h.log.Error("DeletePatient: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("DeletePatient: success, id=" + id)
	c.JSON(200, gin.H{"status": "deleted"})
}