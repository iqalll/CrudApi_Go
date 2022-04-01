package handler

import (
	"net/http"
	"task/repository"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) {
	id := c.Param("id")
	result, _ := repository.GetId(id)
	c.JSON(http.StatusOK, result)
}

func GetAll(c *gin.Context) {
	result, _ := repository.GetAll()
	c.JSON(http.StatusOK, result)
}
func AddData(c *gin.Context) {
	id := c.PostForm("id")
	nama := c.PostForm("nama")
	alamat := c.PostForm("alamat")
	result, _ := repository.AddData(id, nama, alamat)

	c.JSON(http.StatusOK, result)
}
func UpdateData(c *gin.Context) {
	id := c.PostForm("id")
	nama := c.PostForm("nama")
	alamat := c.PostForm("alamat")
	result, _ := repository.UpdateData(id, nama, alamat)
	c.JSON(http.StatusOK, result)
}
func DeleteData(c *gin.Context) {
	id := c.PostForm("id")
	result, _ := repository.DeleteData(id)
	c.JSON(http.StatusOK, result)
}
