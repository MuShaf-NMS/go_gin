package handlers

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/helper"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (uc *userController) GetUser(ctx *gin.Context) {
	if t, ok := ctx.Get("userID"); ok {
		userID := t.(uuid.UUID)
		user := uc.userService.GetUser(userID)
		res := helper.ResponseBuilder("Get User success", user)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized", nil)
	ctx.JSON(401, res)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	if t, ok := ctx.Get("userID"); ok {
		userID := t.(uuid.UUID)
		var user dto.UpdateUser
		ctx.BindJSON(&user)
		err := helper.Validate(user)
		if err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder("Validation error", errs)
			ctx.JSON(403, res)
			return
		}
		uc.userService.UpdateUser(userID, user)
		res := helper.ResponseBuilder("Update User success", nil)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized", nil)
	ctx.JSON(401, res)
}

func (uc *userController) UpdatePassword(ctx *gin.Context) {
	if t, ok := ctx.Get("userID"); ok {
		userID := t.(uuid.UUID)
		var userPass dto.UpdatePasswordUser
		ctx.BindJSON(&userPass)
		err := helper.Validate(userPass)
		if err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder("Validation error", errs)
			ctx.JSON(403, res)
			return
		}
		uc.userService.UpdatePassword(userID, userPass)
		res := helper.ResponseBuilder("Update password success", nil)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized", nil)
	ctx.JSON(401, res)
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}
