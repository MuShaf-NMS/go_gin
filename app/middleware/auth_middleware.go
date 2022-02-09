package middleware

import (
	"strings"

	"github.com/MuShaf-NMS/go_gin/app/helper"
	"github.com/MuShaf-NMS/go_gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Authorize(jwtService service.JwtSerivce) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			res := helper.ErrorResponseBuilder("Unauthorized", nil)
			ctx.JSON(401, res)
			return
		}
		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) < 2 {
			res := helper.ErrorResponseBuilder("Unauthorized", nil)
			ctx.JSON(401, res)
			return
		}
		if splitAuth[0] != "Bearer" {
			res := helper.ErrorResponseBuilder("Unauthorized", nil)
			ctx.JSON(401, res)
			return
		}
		tokenString := splitAuth[1]
		token, er := jwtService.ValidateToken(tokenString)
		if er != nil || !token.Valid {
			res := helper.ErrorResponseBuilder("Unauthorized", nil)
			ctx.JSON(401, res)
			return
		}
		claim := token.Claims.(jwt.MapClaims)
		userID := uuid.MustParse(claim["userID"].(string))
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
