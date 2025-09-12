package pharmacistclientservice 

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/pharmacistpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithPharmacistService(cfg config.Config) (*pb.PharmacistServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.PatientService.Host, cfg.PharmacistService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewPharmacistServiceClient(conn)
	return &userServiceClient, nil
}