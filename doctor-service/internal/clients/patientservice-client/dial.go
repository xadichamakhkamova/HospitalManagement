package patientserviceclient

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/patientpb"
	config "doctor-service/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithPatientService(cfg config.Config) (*pb.PatientManagementServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.PatientService.Host, cfg.PatientService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewPatientManagementServiceClient(conn)
	return &userServiceClient, nil
}