package handler

import (
	"context"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/gin-gonic/gin"
)

// @Router /admin/departments [post]
// @Summary CREATE DEPARTMENT
// @Security BearerAuth
// @Description This method creates department
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param department body models.CreateDepartmentRequest true "Department"
// @Success 200 {object} models.CreateDepartmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateDepartment(c *gin.Context) {

	h.log.Infof("CreateDepartment called")

	req := pb.CreateDepartmentRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Errorf("CreateDepartment invalid request: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateDepartment(context.Background(), &req)
	if err != nil {
		h.log.Errorf("CreateDepartment service error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("Department created successfully with Name=%s", req.Name)
	c.JSON(200, resp)
}

// @Router /admin/departments/{id} [get]
// @Summary GET DEPARTMENT BY ID
// @Security BearerAuth
// @Description This method gets department by id
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Department Id"
// @Success 200 {object} models.GetDepartmentByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDepartmentById(c *gin.Context) {

	id := c.Param("id")
	h.log.Infof("GetDepartmentById called with ID=%s", id)

	req := pb.GetDepartmentByIdRequest{Id: id}
	resp, err := h.service.GetDepartmentById(context.Background(), &req)
	if err != nil {
		h.log.Errorf("GetDepartmentById service error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("Department retrieved successfully with ID=%s", id)
	c.JSON(200, resp)
}

// @Router /admin/departments [get]
// @Summary GET DEPARTMENTS LIST
// @Security BearerAuth
// @Description This method gets departments list by filter
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param search query string false "Product search"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.ListDepartmentsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListDepartments(c *gin.Context) {

	search := strings.TrimSpace(c.Query("search"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	h.log.Infof("ListDepartments called with Search=%s, Page=%d, Limit=%d", search, page, limit)

	req := pb.ListDepartmentsRequest{
		Search: search,
		Page:   int32(page),
		Limit:  int32(limit),
	}

	resp, err := h.service.ListDeparments(context.Background(), &req)
	if err != nil {
		h.log.Errorf("ListDepartments service error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("ListDepartments returned %d departments", len(resp.Deparments))
	c.JSON(200, resp)
}

// @Router /admin/departments/{id} [put]
// @Summary UPDATE DEPARTMENTS
// @Security BearerAuth
// @Description This method updates departments
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Department id"
// @Param product body models.UpdateDepartmentRequest true "Department"
// @Success 200 {object} models.UpdateDepartmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateDepartment(c *gin.Context) {

	id := c.Param("id")
	h.log.Infof("UpdateDepartment called with ID=%s", id)

	req := pb.UpdateDepartmentRequest{Id: id}
	if err := c.BindJSON(&req); err != nil {
		h.log.Errorf("UpdateDepartment invalid request: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateDepartment(context.Background(), &req)
	if err != nil {
		h.log.Errorf("UpdateDepartment service error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("Department updated successfully with ID=%s", id)
	c.JSON(200, resp)
}

// @Router /admin/departments/{id} [delete]
// @Summary DELETE DEPARTMENT
// @Security BearerAuth
// @Description This method deletes department
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Department Id"
// @Success 200 {object} models.DeleteDepartmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDepartment(c *gin.Context) {

	id := c.Param("id")
	h.log.Infof("DeleteDepartment called with ID=%s", id)

	req := pb.DeleteDepartmentRequest{Id: id}
	resp, err := h.service.DeleteDepartment(context.Background(), &req)
	if err != nil {
		h.log.Errorf("DeleteDepartment service error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("Department deleted successfully with ID=%s", id)
	c.JSON(200, resp)
}

// @Router /admin/personals [post]
// @Summary CREATE PERSONAL
// @Security BearerAuth
// @Description This method creates personal
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param personal body models.CreatePersonalRequest true "Personal"
// @Success 200 {object} models.CreatePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreatePersonal(c *gin.Context) {

	req := pb.CreatePersonalRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("failed to bind request in CreatePersonal", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.log.Info("CreatePersonal request received", "request", req.Email)

	resp, err := h.service.CreatePersonal(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to create personal", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("CreatePersonal successful", "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [get]
// @Summary GET PERSONAL BY ID
// @Security BearerAuth
// @Description This method gets personal by id
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param id path string true "Personal Id"
// @Success 200 {object} models.GetPersonalByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetPersonalById(c *gin.Context) {

	req := pb.GetPersonalByIdRequest{Id: c.Param("id")}
	h.log.Info("GetPersonalById request received", "id", req.Id)

	resp, err := h.service.GetPersonalById(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to get personal by id", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("GetPersonalById successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/personals [get]
// @Summary GET PERSONALS LIST
// @Security BearerAuth
// @Description This method gets personals list by filter
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param search query string false "Search by profession or full_name"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.ListPersonalsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListPersonals(c *gin.Context) {

	req := pb.ListPersonalsRequest{}
	req.Search = strings.TrimSpace(c.Query("search"))

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	h.log.Info("ListPersonals request received", "search", req.Search, "page", req.Page, "limit", req.Limit)

	resp, err := h.service.ListPersonals(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to list personals", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("ListPersonals successful", "total", len(resp.Personals))
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [put]
// @Summary UPDATE PERSONAL
// @Security BearerAuth
// @Description This method updates personal
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param id path string true "Personal Id"
// @Param personal body models.UpdatePersonalRequest true "Personal"
// @Success 200 {object} models.UpdatePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdatePersonal(c *gin.Context) {

	req := pb.UpdatePersonalRequest{Id: c.Param("id")}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("failed to bind request in UpdatePersonal", "id", req.Id, "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.log.Info("UpdatePersonal request received", "id", req.Id, "request", req.Email)

	resp, err := h.service.UpdatePersonal(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to update personal", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("UpdatePersonal successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [delete]
// @Summary DELETE PERSONAL
// @Security BearerAuth
// @Description This method deletes personal
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param id path string true "Personal Id"
// @Success 200 {object} models.DeletePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeletePersonal(c *gin.Context) {

	req := pb.DeletePersonalRequest{Id: c.Param("id")}
	h.log.Info("DeletePersonal request received", "id", req.Id)

	resp, err := h.service.DeletePersonal(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to delete personal", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("DeletePersonal successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/doctors [post]
// @Summary CREATE DOCTOR
// @Security BearerAuth
// @Description This method creates doctor
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param doctor body models.CreateDoctorRequest true "Doctor"
// @Success 200 {object} models.CreateDoctorResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateDoctor(c *gin.Context) {

	req := pb.CreateDoctorRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("failed to bind request in CreateDoctor", "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.log.Info("CreateDoctor request received", "request", req.PersonalId)

	resp, err := h.service.CreateDoctor(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to create doctor", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("CreateDoctor successful", "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [get]
// @Summary GET DOCTOR BY ID
// @Security BearerAuth
// @Description This method gets doctor by id
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param id path string true "Doctor Id"
// @Success 200 {object} models.GetDoctorByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDoctorById(c *gin.Context) {

	req := pb.GetPersonalByIdRequest{Id: c.Param("id")}
	h.log.Info("GetDoctorById request received", "id", req.Id)

	resp, err := h.service.GetDoctorById(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to get doctor by id", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("GetDoctorById successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/doctors [get]
// @Summary GET DOCTORS LIST
// @Security BearerAuth
// @Description This method gets doctors list by filter
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param search query string false "Search by name, specialization, etc"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.ListDoctorsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListDoctors(c *gin.Context) {

	req := pb.ListPersonalsRequest{}
	req.Search = strings.TrimSpace(c.Query("search"))

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	h.log.Info("ListDoctors request received", "search", req.Search, "page", req.Page, "limit", req.Limit)

	resp, err := h.service.ListDoctors(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to list doctors", "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("ListDoctors successful", "count", len(resp.Doctors))
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [put]
// @Summary UPDATE DOCTOR
// @Security BearerAuth
// @Description This method updates doctor
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param id path string true "Doctor Id"
// @Param doctor body models.UpdateDoctorRequest true "Doctor"
// @Success 200 {object} models.UpdateDoctorResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateDoctor(c *gin.Context) {

	req := pb.UpdateDoctorRequest{Id: c.Param("id")}
	if err := c.BindJSON(&req); err != nil {
		h.log.Error("failed to bind request in UpdateDoctor", "id", req.Id, "error", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.log.Info("UpdateDoctor request received", "id", req.Id, "request", req.Id)

	resp, err := h.service.UpdateDoctor(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to update doctor", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("UpdateDoctor successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [delete]
// @Summary DELETE DOCTOR
// @Security BearerAuth
// @Description This method deletes doctor
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param id path string true "Doctor Id"
// @Success 200 {object} models.DeletePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDoctor(c *gin.Context) {

	req := pb.DeletePersonalRequest{Id: c.Param("id")}
	h.log.Info("DeleteDoctor request received", "id", req.Id)

	resp, err := h.service.DeleteDoctor(context.Background(), &req)
	if err != nil {
		h.log.Error("failed to delete doctor", "id", req.Id, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("DeleteDoctor successful", "id", req.Id, "response", resp)
	c.JSON(200, resp)
}

// @Router /admin/beds [post]
// @Summary CREATE BED
// @Security  		BearerAuth
// @Description This method creates bed
// @Tags BEDS
// @Accept json
// @Produce json
// @Param bed body models.CreateBedRequest true "Bed"
// @Success 200 {object} models.CreateBedResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateBed(c *gin.Context) {

	h.log.Info("CreateBed: request received")

	req := pb.CreateBedRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.log.WithError(err).Warn("CreateBed: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateBed(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("CreateBed: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("bed_id", resp.Bed.Id).Info("CreateBed: success")
	c.JSON(200, resp)
}

// @Router /admin/beds/{id} [get]
// @Summary GET BED BY ID
// @Security  		BearerAuth
// @Description This method gets bed by id
// @Tags BEDS
// @Accept json
// @Produce json
// @Param id path string true "Bed Id"
// @Success 200 {object} models.GetBedByIDResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetBedById(c *gin.Context) {

	id := c.Param("id")
	h.log.WithField("id", id).Info("GetBedById: request received")

	req := pb.GetBedByIDRequest{Id: id}
	resp, err := h.service.GetBedByID(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("GetBedById: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("bed_id", resp.Bed.Id).Info("GetBedById: success")
	c.JSON(200, resp)
}

// @Router /admin/beds [get]
// @Summary GET BEDS LIST
// @Security  		BearerAuth
// @Description This method gets beds list by filter
// @Tags BEDS
// @Accept json
// @Produce json
// @Param search query string false "Search by room number, status, etc"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.ListBedsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) ListBeds(c *gin.Context) {

	search := strings.TrimSpace(c.Query("search"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	h.log.WithFields(logrus.Fields{
		"search": search,
		"page":   page,
		"limit":  limit,
	}).Info("ListBeds: request received")

	req := pb.ListBedSRequest{
		Search: search,
		Page:   int32(page),
		Limit:  int32(limit),
	}

	resp, err := h.service.ListBedS(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("ListBeds: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("count", len(resp.Beds)).Info("ListBeds: success")
	c.JSON(200, resp)
}

// @Router /admin/beds/{id} [put]
// @Summary UPDATE BED
// @Security  		BearerAuth
// @Description This method updates bed
// @Tags BEDS
// @Accept json
// @Produce json
// @Param id path string true "Bed Id"
// @Param bed body models.UpdateBedRequest true "Bed"
// @Success 200 {object} models.UpdateBedResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateBed(c *gin.Context) {

	id := c.Param("id")
	h.log.WithField("id", id).Info("UpdateBed: request received")

	req := pb.UpdateBedRequest{Id: id}
	if err := c.BindJSON(&req); err != nil {
		h.log.WithError(err).Warn("UpdateBed: invalid request body")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateBed(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("UpdateBed: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("bed_id", resp.Bed.Id).Info("UpdateBed: success")
	c.JSON(200, resp)
}

// @Router /admin/beds/{id} [delete]
// @Summary DELETE BED
// @Security  		BearerAuth
// @Description This method deletes bed
// @Tags BEDS
// @Accept json
// @Produce json
// @Param id path string true "Bed Id"
// @Success 200 {object} models.DeleteBedResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteBed(c *gin.Context) {
	id := c.Param("id")
	h.log.WithField("id", id).Info("DeleteBed: request received")

	req := pb.DeleteBedRequest{Id: id}
	resp, err := h.service.DeleteBed(context.Background(), &req)
	if err != nil {
		h.log.WithError(err).Error("DeleteBed: service error")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.log.WithField("bed_id", id).Info("DeleteBed: success")
	c.JSON(200, resp)
}
