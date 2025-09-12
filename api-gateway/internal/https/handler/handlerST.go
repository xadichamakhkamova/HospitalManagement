package handler

import (
	"api-gateway/internal/service"
)

type HandlerST struct {
	service     *service.ServiceRepositoryClient
}

func NewApiHandler(service *service.ServiceRepositoryClient) *HandlerST {
	return &HandlerST{
		service:     service,
	}
}
