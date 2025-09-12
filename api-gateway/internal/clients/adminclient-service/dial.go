package adminclientservice 

import (
	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithAdminService(cfg config.Config) (*pb.AdminServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.AdminService.Host, cfg.AdminService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewAdminServiceClient(conn) // returns interface 
	return &userServiceClient, nil
}