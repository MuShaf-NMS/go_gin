package router

import (
	"fmt"

	"github.com/MuShaf-NMS/go_gin/app/config"
	"github.com/MuShaf-NMS/go_gin/app/handlers"
	"github.com/MuShaf-NMS/go_gin/app/middleware"
	"github.com/MuShaf-NMS/go_gin/app/repository"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutingUser(router *gin.Engine, db *gorm.DB, config config.Config) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	jwtService := service.NewJwtService(config.SecretKey)
	userController := handlers.NewUserController(userService)
	fmt.Println("Routing for User ...")
	routerAuth := router.Group("/user", middleware.Authorize(jwtService))
	{
		routerAuth.GET("/", userController.GetUser)
		routerAuth.PUT("/", userController.UpdateUser)
		routerAuth.PUT("/password", userController.UpdatePassword)
	}
}
