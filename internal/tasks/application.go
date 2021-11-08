package tasks

import (
	"authe/internal/config"
	"authe/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Server() {

	config.CreateDatabaseConnection()
	route := gin.Default()
	route.Use(func(ctx *gin.Context) {
		defer LoggingAndErrorMapping(ctx)
	})
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, model.Response{404, "Not Found"})
	})

}

func LoggingAndErrorMapping(ctx *gin.Context) {
	if rec := recover(); rec != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{500, "Internal server error"})
	} else {

	}

}
