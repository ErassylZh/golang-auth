package middleware

import (
	"auth/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, "Unauthorized")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			cfg, _ := config.GetConfig()
			return cfg.Security.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
