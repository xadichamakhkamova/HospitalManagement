package nurseclientservice 

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/nursepb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithNurseService(cfg config.Config) (*pb.NurseServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.NurseService.Host, cfg.NurseService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewNurseServiceClient(conn)
	return &userServiceClient, nil
}