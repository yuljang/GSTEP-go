package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuljang/GSTEP-go/database"
	"github.com/yuljang/GSTEP-go/models"
	"gorm.io/datatypes"
)

func CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Milestone = datatypes.JSON([]byte(fmt.Sprintf(`{"0": "%s"}`, time.Now().String())))
	user.Progress = datatypes.JSON([]byte(`{}`))

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

		var milestone map[string]string

		json.Unmarshal([]byte(user.Milestone), &milestone)

		if _, exist := milestone[strconv.Itoa(step)]; !exist {
			fmt.Println(strconv.Itoa(step))
			milestone[strconv.Itoa(step)] = time.Now().String()
		}

		if milestone_str, err := json.Marshal(milestone); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			update.Milestone = milestone_str
		}

	}

	database.DB.Model(&user).Updates(update)

	c.JSON(http.StatusOK, user)
}

func DiscoverMission(c *gin.Context) {
	update := models.User{}
	user := models.User{}

	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var progress map[string]int

	json.Unmarshal([]byte(update.Progress), &progress)

	progress[c.Param("mission")] = 0

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

		var milestone map[string]string

		json.Unmarshal([]byte(user.Milestone), &milestone)

		if _, exist := milestone[strconv.Itoa(step)]; !exist {
			fmt.Println(strconv.Itoa(step))
			milestone[strconv.Itoa(step)] = time.Now().String()
		}

		if milestone_str, err := json.Marshal(milestone); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			update.Milestone = milestone_str
		}

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
