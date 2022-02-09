package handlers

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/helper"
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type AuthController interface {
	RegisterUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtSerivce
}

func (ac *authController) RegisterUser(ctx *gin.Context) {
	var user dto.CreateUser
	ctx.BindJSON(&user)
	err := helper.Validate(user)
	if err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Validation error", errs)
		ctx.JSON(400, res)
		return
	}
	if ac.authService.CheckUser(user) {
		res := helper.ErrorResponseBuilder("Username or Email exist", nil)
		ctx.JSON(403, res)
		return
	}
	ac.authService.RegisterUser(user)
	res := helper.ResponseBuilder("Add new User success", nil)
	ctx.JSON(200, res)
}

func (ac *authController) Login(ctx *gin.Context) {
	var form dto.LoginDTO
	ctx.BindJSON(&form)
	err := helper.Validate(form)
	if err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Validation error", errs)
		ctx.JSON(403, res)
		return
	}
	verified := ac.authService.VerifyUser(form)
	if user, ok := verified.(models.User); ok {
		token := ac.jwtService.GenerateToken(user.UUID)
		res := helper.ResponseBuilder("Login success", gin.H{"token": token, "user": user})
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Username or Password wrong", nil)
	ctx.JSON(403, res)

}

func NewAuthController(authService service.AuthService, jwtService service.JwtSerivce) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
