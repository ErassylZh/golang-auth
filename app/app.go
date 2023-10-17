package app

import (
	"auth/config"
	"auth/handler"
	"auth/model"
	"auth/repository"
	"auth/service"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to connect to database")
	}
	repositories, err := repository.NewRepositories(db, cfg)
	if err != nil {
		panic("error initialization Repositories " + err.Error())
	}
	services := service.NewServices(service.Deps{
		Repos: repositories,
		Cgf:   cfg,
	})
	//handlers := v1.NewHandler(services)
	handlerDelivery := handler.NewHandlerDelivery(services, "auth")
	handlerDelivery.InitAPI(r)
	r.Run(":" + cfg.Service.Port)
}
