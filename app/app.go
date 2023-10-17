package app

import (
	"auth/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	r := gin.Default()
	dsn := "user=" + cfg.Database.DBUser + " password=" + cfg.Database.DBPassword +
		" dbname=" + cfg.Database.DBName + " sslmode=" + cfg.Database.SSLMODE
	fmt.Println(dsn)
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	r.Run(":" + cfg.Service.Port)
}
