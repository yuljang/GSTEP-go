package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuljang/GSTEP-go/database"
	"github.com/yuljang/GSTEP-go/models"
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

	var progress map[string]int

	json.Unmarshal([]byte(update.Progress), &progress)

	if update.Progress != nil {
		point := 0
		step := 0

		for _, value := range progress {
			point += value + 2
		}

		switch {
		case point < 5:
			step = 0
		case point < 10:
			step = 1
		case point < 25:
			step = 2
		case point < 60:
			step = 3
		default:
			step = 4
		}

		update.Step = step
		update.Point = point
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
