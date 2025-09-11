package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Service        ServiceConfig
	AdminService   ServiceConfig
	PatientService ServiceConfig
	Postgres       PostgresConfig
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetInt("service.port"),
		},
		AdminService: ServiceConfig{
			Host: viper.GetString("clients.admin-service.host"),
			Port: viper.GetInt("clients.patient-service.port"),
		},
		PatientService: ServiceConfig{
			Host: viper.GetString("clients.patient-service.host"),
			Port: viper.GetInt("clients.patient-service.port"),
		},
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		},
	}

	return &cfg, nil
}
