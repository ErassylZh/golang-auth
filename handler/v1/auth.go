package v1

import (
	"auth/middleware"
	"auth/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuth(v1 *gin.RouterGroup) {
	model := v1.Group("/auth")
	model.POST("/login", middleware.GinErrorHandle(h.Login))
	model.POST("/register", middleware.GinErrorHandle(h.Register))
	model.POST("/test", middleware.GinErrorHandle(h.Login))
}

func (h *Handler) Login(c *gin.Context) error {
	var request schema.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		return err
	}
	result, err := h.services.User.GetForLogin(request)
	if err != nil {
		return err
	}
	return schema.Respond(result, c)
}
func (h *Handler) Register(c *gin.Context) error {
	var request schema.RegisterRequest
	err := c.BindJSON(&request)
	if err != nil {
		return err
	}
	result, err := h.services.User.Create(request)
	if err != nil {
		return err
	}
	return schema.Respond(result, c)
}
