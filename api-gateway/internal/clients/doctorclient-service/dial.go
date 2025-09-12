package doctorclientservice

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/doctorpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithDoctorService(cfg config.Config) (*pb.DoctorServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.DoctorService.Host, cfg.DoctorService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewDoctorServiceClient(conn)
	return &userServiceClient, nil
}