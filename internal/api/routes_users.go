package api

import (
	"authe/internal/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {

	userRoute := router.Group("/users")

	userRoute.POST("/create", func(ctx *gin.Context) {
		firstName := ctx.PostForm("first_name")
		lastName := ctx.PostForm("last_name")
		email := ctx.PostForm("email")
		phoneNumber := ctx.PostForm("msisdn")
		service.CreateUser(firstName, lastName, phoneNumber, email)

	})

}
