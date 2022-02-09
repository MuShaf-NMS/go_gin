package router

import (
	"fmt"

	"github.com/MuShaf-NMS/go_gin/app/config"
	"github.com/MuShaf-NMS/go_gin/app/handlers"
	"github.com/MuShaf-NMS/go_gin/app/repository"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutingAuth(router *gin.Engine, db *gorm.DB, config config.Config) {
	authRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepository)
	jwtService := service.NewJwtService(config.SecretKey)
	authController := handlers.NewAuthController(authService, jwtService)
	fmt.Println("Routing for Auth ...")
	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", authController.Login)
		routerAuth.POST("/register", authController.RegisterUser)
	}
}
