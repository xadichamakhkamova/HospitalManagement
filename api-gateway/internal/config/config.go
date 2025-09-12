package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	ApiGateway        ServiceConfig
	AdminService      ServiceConfig
	DoctorService     ServiceConfig
	NurseService      ServiceConfig
	PatientService    ServiceConfig
	PharmacistService ServiceConfig
	TokenKey          string
	CertFile          string
	KeyFile           string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ApiGateway: ServiceConfig{
			Host: viper.GetString("api-gateway.host"),
			Port: viper.GetInt("api-gateway.port"),
		},
		AdminService: ServiceConfig{
			Host: viper.GetString("services.admin-service.host"),
			Port: viper.GetInt("services.admin-service.port"),
		},
		DoctorService: ServiceConfig{
			Host: viper.GetString("services.doctor-service.host"),
			Port: viper.GetInt("services.doctor-service.port"),
		},
		NurseService: ServiceConfig{
			Host: viper.GetString("services.nurse-service.host"),
			Port: viper.GetInt("services.nurse-service.port"),
		},
		PatientService: ServiceConfig{
			Host: viper.GetString("services.patient-service.host"),
			Port: viper.GetInt("services.patient-service.port"),
		},
		PharmacistService: ServiceConfig{
			Host: viper.GetString("services.pharmacist-service.host"),
			Port: viper.GetInt("services.pharmacist-service.port"),
		},

		TokenKey: viper.GetString("token.key"),

		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}
	return &cfg, nil
}
