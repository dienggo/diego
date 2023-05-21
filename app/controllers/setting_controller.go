package controllers

import (
	"github.com/gin-gonic/gin"
	"go_base_project/app/dto/dto_request"
	"go_base_project/app/models"
	"go_base_project/provider/database"
	"net/http"
)

type SettingController struct{}

func (controller SettingController) GetByKey(c *gin.Context) {

	var request dto_request.SettingByKey

	// bind data into dto
	errBindRequest := c.Bind(&request)
	if errBindRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to parse into dto",
		})
		return
	}

	// validate data
	errValidate := request.Validate()
	if errValidate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errValidate.Error(),
			"request": request,
		})
		return
	}

	var settingData models.Setting
	result := database.Main().Unscoped().Where("`key` = ?", request.Key).First(&settingData)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"key":   request.Key,
		"value": settingData.Value,
	})
}
