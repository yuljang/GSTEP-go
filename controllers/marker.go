package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuljang/GSTEP-go/database"
	"github.com/yuljang/GSTEP-go/models"
	"gorm.io/gorm"
)

func CreateMarker(c *gin.Context) {
	marker := models.Marker{}
	target := models.Marker{}

	if err := c.ShouldBind(&marker); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.Where("sqrt(85*power(latitude - ?, 2) + 111*power(longitude - ?, 2)) < 0.05", marker.Latitude, marker.Longitude).First(&target).Error; err != gorm.ErrRecordNotFound {
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	if err := database.DB.Create(&marker).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, marker)
}

func ReadMarker(c *gin.Context) {
	marker := models.Marker{}

	if err := database.DB.First(&marker, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, marker)
}

func UpdateMarker(c *gin.Context) {
	update := models.Marker{}
	marker := models.Marker{}

	if err := c.ShouldBind(&update); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := database.DB.First(&marker, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	database.DB.Model(&marker).Updates(update)

	c.JSON(http.StatusOK, marker)
}

func DeleteMarker(c *gin.Context) {
	marker := models.Marker{}

	if err := database.DB.Delete(&marker, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, marker)
}
