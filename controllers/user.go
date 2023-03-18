package controllers

import (
	"net/http"

	"github.com/gdsc-ys/21days-gin/database"
	"github.com/gdsc-ys/21days-gin/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func ReadUser(c *gin.Context) {
	user := models.User{}

	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	update := models.User{}
	user := models.User{}

	if err := c.ShouldBind(&update); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	database.DB.Model(&user).Updates(update)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	user := models.User{}

	if err := database.DB.Delete(&user, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
