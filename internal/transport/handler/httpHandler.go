package handler

import (
	"order-providing-system/internal/service"
	"order-providing-system/internal/service/usecase"
	"order-providing-system/pkg/logger"
)

type HttpCustomHandler struct {
	service service.ServiceModule
	Logger  *logger.CustomLogger
}

func NewHttpHandler(s *usecase.CustomService, l *logger.CustomLogger) *HttpCustomHandler {
	return &HttpCustomHandler{service: s}
}
