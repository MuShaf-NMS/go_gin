package handlers

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/helper"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	GetOneTodo(ctx *gin.Context)
	GetAllTodos(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type todoController struct {
	service service.TodoService
}

func (tc *todoController) CreateTodo(ctx *gin.Context) {
	if t, ok := ctx.Get("userID"); ok {
		userID := t.(uuid.UUID)
		var todo dto.CreateTodo
		ctx.BindJSON(&todo)
		err := helper.Validate(todo)
		if err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder("Validation error", errs)
			ctx.JSON(403, res)
			return
		}
		tc.service.CreateTodo(todo, userID)
		res := helper.ResponseBuilder("Add new Todo success", nil)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized", nil)
	ctx.JSON(401, res)
}

func (tc *todoController) GetOneTodo(ctx *gin.Context) {
	todoID, err := uuid.Parse(ctx.Param("todoID"))
	if err != nil {
		res := helper.ErrorResponseBuilder("UUID format is wrong", nil)
		ctx.JSON(400, res)
		return
	}
	todo, err := tc.service.GetOneTodo(todoID)
	if err != nil {
		res := helper.ErrorResponseBuilder("Todo not found", nil)
		ctx.JSON(400, res)
		return
	}
	res := helper.ResponseBuilder("Get Todo success", todo)
	ctx.JSON(200, res)
}

func (tc *todoController) GetAllTodos(ctx *gin.Context) {
	if t, ok := ctx.Get("userID"); ok {
		userID := t.(uuid.UUID)
		todos := tc.service.GetAllTodos(userID)
		res := helper.ResponseBuilder("Get Todos success", todos)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized", nil)
	ctx.JSON(401, res)
}

func (tc *todoController) UpdateTodo(ctx *gin.Context) {
	todoID, err := uuid.Parse(ctx.Param("todoID"))
	if err != nil {
		res := helper.ErrorResponseBuilder("UUID format is wrong", nil)
		ctx.JSON(400, res)
		return
	}
	var todo dto.UpdateTodo
	ctx.BindJSON(&todo)
	err = helper.Validate(todo)
	if err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Validation error", errs)
		ctx.JSON(403, res)
		return
	}
	tc.service.UpdateTodo(todo, todoID)
	res := helper.ResponseBuilder("Update Todo success", nil)
	ctx.JSON(200, res)
}

func (tc *todoController) DeleteTodo(ctx *gin.Context) {
	todoID, err := uuid.Parse(ctx.Param("todoID"))
	if err != nil {
		res := helper.ErrorResponseBuilder("UUID format is wrong", nil)
		ctx.JSON(400, res)
		return
	}
	tc.service.DeleteTodo(todoID)
	res := helper.ResponseBuilder("Delete Todo success", nil)
	ctx.JSON(200, res)
}

func NewTodoController(service service.TodoService) TodoController {
	return &todoController{
		service: service,
	}
}
