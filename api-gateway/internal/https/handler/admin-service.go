package handler

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"github.com/gin-gonic/gin"
)

// @Router /admin/departments [post]
// @Summary CREATE DEPARTMENT
// @Security  		BearerAuth
// @Description This method creates department
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param department body models.CreateDepartmentRequest true "Department"
// @Success 200 {object} models.CreateDepartmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateDepartment(c *gin.Context) {

	req := pb.CreateDepartmentRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var ctx context.Context
	resp, err := h.service.CreateDepartment(ctx, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Router admin/deparments/{id} [get]
// @Summary GET DEPARTMENT BY ID
// @Security  		BearerAuth
// @Description This method gets department by id
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Department Id"
// @Success 200 {object} models.GetDepartmentByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDepartmentById(c *gin.Context) {

	req := pb.GetDepartmentByIdRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.GetDepartmentById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/departments	[get]
// @Summary GET DEPARTMENTS LIST
// @Security  		BearerAuth
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

	req := pb.ListDepartmentsRequest{}
	search := c.Query("search")

	req.Search = strings.TrimSpace(search)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListDeparments(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/departments/{id} [put]
// @Summary UPDATE DEPARTMENTS
// @Security  		BearerAuth
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

	req := pb.UpdateDepartmentRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateDepartment(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/departments/{id} [delete]
// @Summary DELETE DEPARTMENT
// @Security  		BearerAuth
// @Description This method deletes department
// @Tags DEPARTMENTS
// @Accept json
// @Produce json
// @Param id path string true "Department Id"
// @Success 200 {object} models.DeleteDepartmentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDepartment(c *gin.Context) {

	req := pb.DeleteDepartmentRequest{}
	req.Id = c.Param("id")
	resp, err := h.service.DeleteDepartment(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/personals [post]
// @Summary CREATE PERSONAL
// @Security  		BearerAuth
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.CreatePersonal(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [get]
// @Summary GET PERSONAL BY ID
// @Security  		BearerAuth
// @Description This method gets personal by id
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param id path string true "Personal Id"
// @Success 200 {object} models.GetPersonalByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetPersonalById(c *gin.Context) {

	req := pb.GetPersonalByIdRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetPersonalById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/personals [get]
// @Summary GET PERSONALS LIST
// @Security  		BearerAuth
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
	search := c.Query("search")
	req.Search = strings.TrimSpace(search)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListPersonals(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [put]
// @Summary UPDATE PERSONAL
// @Security  		BearerAuth
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

	req := pb.UpdatePersonalRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdatePersonal(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/personals/{id} [delete]
// @Summary DELETE PERSONAL
// @Security  		BearerAuth
// @Description This method deletes personal
// @Tags PERSONALS
// @Accept json
// @Produce json
// @Param id path string true "Personal Id"
// @Success 200 {object} models.DeletePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeletePersonal(c *gin.Context) {
	
	req := pb.DeletePersonalRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeletePersonal(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/doctors [post]
// @Summary CREATE DOCTOR
// @Security  		BearerAuth
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.CreateDoctor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [get]
// @Summary GET DOCTOR BY ID
// @Security  		BearerAuth
// @Description This method gets doctor by id
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param id path string true "Doctor Id"
// @Success 200 {object} models.GetDoctorByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetDoctorById(c *gin.Context) {

	req := pb.GetPersonalByIdRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetDoctorById(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/doctors [get]
// @Summary GET DOCTORS LIST
// @Security  		BearerAuth
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
	search := c.Query("search")
	req.Search = strings.TrimSpace(search)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListDoctors(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [put]
// @Summary UPDATE DOCTOR
// @Security  		BearerAuth
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

	req := pb.UpdateDoctorRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateDoctor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /admin/doctors/{id} [delete]
// @Summary DELETE DOCTOR
// @Security  		BearerAuth
// @Description This method deletes doctor
// @Tags DOCTORS
// @Accept json
// @Produce json
// @Param id path string true "Doctor Id"
// @Success 200 {object} models.DeletePersonalResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) DeleteDoctor(c *gin.Context) {

	req := pb.DeletePersonalRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeleteDoctor(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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

	req := pb.CreateBedRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.CreateBed(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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

	req := pb.GetBedByIDRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.GetBedByID(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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

	req := pb.ListBedSRequest{}
	search := c.Query("search")
	req.Search = strings.TrimSpace(search)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	req.Page = int32(page)
	req.Limit = int32(limit)

	resp, err := h.service.ListBedS(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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

	req := pb.UpdateBedRequest{}
	req.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.UpdateBed(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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
	
	req := pb.DeleteBedRequest{}
	req.Id = c.Param("id")

	resp, err := h.service.DeleteBed(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
