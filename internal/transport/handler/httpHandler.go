package handler

import (
	"frappuccino/internal/service"
	"frappuccino/internal/service/usecase"
	"frappuccino/pkg/logger"
)

type HttpCustomHandler struct {
	service service.ServiceModule
	Logger  *logger.CustomLogger
}

func NewHttpHandler(s *usecase.CustomService, l *logger.CustomLogger) *HttpCustomHandler {
	return &HttpCustomHandler{service: s}
}
