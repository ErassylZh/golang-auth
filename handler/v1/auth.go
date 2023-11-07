package v1

import (
	"auth/middleware"
	"auth/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) initAuth(v1 *gin.RouterGroup) {
	model := v1.Group("/auth")
	model.POST("/login", middleware.GinErrorHandle(h.Login))
	model.POST("/register", middleware.GinErrorHandle(h.Register))
	model.POST("/me", middleware.GinErrorHandle(h.GetUser))
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
func (h *Handler) GetUser(c *gin.Context) error {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return fmt.Errorf("Отсутствует заголовок Authorization")
	}

	// Проверяем, начинается ли заголовок с "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fmt.Errorf("Неправильный формат заголовка Authorization")
	}

	// Извлекаем токен, удаляя "Bearer " из строки
	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.services.User.ExtractUserIDFromJWT(token)
	if err != nil {
		return err
	}
	return schema.Respond(userID, c)
}
