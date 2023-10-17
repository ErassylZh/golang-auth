package handler

import (
	v1 "auth/handler/v1"
	"auth/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
	baseUrl  string
}

func NewHandlerDelivery(
	services *service.Services,
	baseUrl string,
) *Handler {
	return &Handler{
		services: services,
		baseUrl:  baseUrl,
	}
}

func (h *Handler) InitAPI(router *gin.Engine) {
	baseUrl := router.Group(h.baseUrl)

	handlerV1 := v1.NewHandler(h.services)
	api := baseUrl.Group("/api")
	{
		handlerV1.Init(api)
	}
}
