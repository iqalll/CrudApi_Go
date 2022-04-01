package server

import (
	"task/handler"

	"github.com/gin-gonic/gin"
)

func Server() {

	router := gin.Default()
	router.GET("/:id", handler.GetId)
	router.GET("/", handler.GetAll)
	router.POST("/", handler.AddData)
	router.PUT("/", handler.UpdateData)
	router.DELETE("/", handler.DeleteData)

	router.Run(":8080")

}
