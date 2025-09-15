package https

import (
	"api-gateway/internal/https/handler"
	"api-gateway/internal/service"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(service *service.ServiceRepositoryClient, port int, log *logrus.Logger) *http.Server {

	r := gin.Default()

	apiHandler := handler.NewApiHandler(service, log)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//! Admin Service
	admin := r.Group("/admin")
	{
		// Department
		admin.POST("/departments", apiHandler.CreateDepartment)
		admin.GET("/departments/:id", apiHandler.GetDepartmentById)
		admin.GET("/departments", apiHandler.ListDepartments)
		admin.PUT("/departments/:id", apiHandler.UpdateDepartment)
		admin.DELETE("/departments/:id", apiHandler.DeleteDepartment)

		// Personal
		admin.POST("/personals", apiHandler.CreatePersonal)
		admin.GET("/personals/:id", apiHandler.GetPersonalById)
		admin.GET("/personals", apiHandler.ListPersonals)
		admin.PUT("/personals/:id", apiHandler.UpdatePersonal)
		admin.DELETE("/personals/:id", apiHandler.DeletePersonal)

		// Doctor
		admin.POST("/doctors", apiHandler.CreateDoctor)
		admin.GET("/doctors/:id", apiHandler.GetDoctorById)
		admin.GET("/doctors", apiHandler.ListDoctors)
		admin.PUT("/doctors/:id", apiHandler.UpdateDoctor)
		admin.DELETE("/doctors/:id", apiHandler.DeleteDoctor)

		// Bed
		admin.POST("/beds", apiHandler.CreateBed)
		admin.GET("/beds/:id", apiHandler.GetBedById)
		admin.GET("/beds", apiHandler.ListBeds)
		admin.PUT("/beds/:id", apiHandler.UpdateBed)
		admin.DELETE("/beds/:id", apiHandler.DeleteBed)
	}

	//! Doctor Service
	doctor := r.Group("/doctor")
	{
		// Appointment
		doctor.POST("/appointments", apiHandler.CreateAppointment)
		doctor.GET("/appointments/:id", apiHandler.GetAppointmentById)
		doctor.GET("/appointments", apiHandler.ListAppointments)
		doctor.PUT("/appointments/:id", apiHandler.UpdateAppointment)
		doctor.DELETE("/appointments/:id", apiHandler.DeleteAppointment)

		// Prescription
		doctor.POST("/prescriptions", apiHandler.CreatePrescription)
		doctor.GET("/prescriptions/:id", apiHandler.GetPrescriptionById)
		doctor.GET("/prescriptions", apiHandler.ListPrescriptions)
		doctor.PUT("/prescriptions/:id", apiHandler.UpdatePrescription)
		doctor.DELETE("/prescriptions/:id", apiHandler.DeletePrescription)
	}

	//! Nurse Service
	nurse := r.Group("/nurse")
	{
		nurse.POST("/donors", apiHandler.CreateDonor)
		nurse.GET("/donors/:id", apiHandler.GetDonorById)
		nurse.GET("/donors", apiHandler.ListDonors)
		nurse.PUT("/donors/:id", apiHandler.UpdateDonor)
		nurse.DELETE("/donors/:id", apiHandler.DeleteDonor)
	}

	//! Patient Service
	patient := r.Group("/patient")
	{
		patient.POST("/patients", apiHandler.CreatePatient)
		patient.GET("/patients/:id", apiHandler.GetPatientById)
		patient.GET("/patients", apiHandler.ListPatients)
		patient.PUT("/patients/:id", apiHandler.UpdatePatient)
		patient.DELETE("/patients/:id", apiHandler.DeletePatient)
	}

	//! Pharmacist Service
	pharmacist := r.Group("/pharmacist")
	{
		pharmacist.POST("/medicines", apiHandler.CreateMedicine)
		pharmacist.GET("/medicines/:id", apiHandler.GetMedicineById)
		pharmacist.GET("/medicines", apiHandler.ListMedicines)
		pharmacist.PUT("/medicines/:id", apiHandler.UpdateMedicine)
		pharmacist.DELETE("/medicines/:id", apiHandler.DeleteMedicine)
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	address := fmt.Sprintf(":%d", port)
	srv := &http.Server{
		Addr:      address,
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}
