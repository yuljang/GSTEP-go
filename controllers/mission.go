package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuljang/GSTEP-go/database"
	"github.com/yuljang/GSTEP-go/models"
)

func CreateMission(c *gin.Context) {
	mission := models.Mission{}

	if err := c.ShouldBind(&mission); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.Create(&mission).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, mission)
}

func ReadMission(c *gin.Context) {
	mission := models.Mission{}

	if err := database.DB.First(&mission, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, mission)
}

func UpdateMission(c *gin.Context) {
	update := models.Mission{}
	mission := models.Mission{}

	if err := c.ShouldBind(&update); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.First(&mission, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	database.DB.Model(&mission).Updates(update)

	c.JSON(http.StatusOK, mission)
}

func DeleteMission(c *gin.Context) {
	mission := models.Mission{}

	if err := database.DB.Delete(&mission, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, mission)
}
