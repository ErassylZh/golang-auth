package v1

import (
	"auth/middleware"
	"auth/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuth(v1 *gin.RouterGroup) {
	model := v1.Group("")
	model.POST("/login", middleware.GinErrorHandle(h.Login))
	model.POST("/register", middleware.GinErrorHandle(h.Register))
}

func (h *Handler) Login(c *gin.Context) error {
	return nil
}
func (h *Handler) Register(c *gin.Context) error {
	test := "nice"
	return schema.Respond(test, c)
}
