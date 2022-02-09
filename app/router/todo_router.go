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

func RoutingTodo(router *gin.Engine, db *gorm.DB, config config.Config) {
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	jwtService := service.NewJwtService(config.SecretKey)
	todoController := handlers.NewTodoController(todoService)
	fmt.Println("Routing for Todo ...")
	routerAuth := router.Group("/todo", middleware.Authorize(jwtService))
	{
		routerAuth.GET("/", todoController.GetAllTodos)
		routerAuth.POST("/", todoController.CreateTodo)
		routerAuth.GET("/:todoID", todoController.GetOneTodo)
		routerAuth.PUT("/:todoID", todoController.UpdateTodo)
		routerAuth.DELETE("/:todoID", todoController.DeleteTodo)
	}
}
