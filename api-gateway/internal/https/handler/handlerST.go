package handler

import (
	"api-gateway/internal/service"

	"github.com/sirupsen/logrus"
)

type HandlerST struct {
	service *service.ServiceRepositoryClient
	log     *logrus.Logger
}

func NewApiHandler(service *service.ServiceRepositoryClient, log *logrus.Logger) *HandlerST {
	return &HandlerST{
		service: service,
		log:     log,
	}
}
