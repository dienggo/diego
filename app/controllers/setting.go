package controllers

import (
	"github.com/dienggo/diego/app/dto/dto_request"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/services"
	"github.com/dienggo/diego/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Setting struct{}

func (controller Setting) GetByKeyOnService(c *gin.Context) {
	var request dto_request.SettingByKey

	// bind data into dto
	errBindRequest := c.Bind(&request)
	if errBindRequest != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Fail to parse into dto",
		})
		return
	}

	var settingData models.Setting
	err := services.NewExample(request).Do().ResultParse(&settingData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"key":   request.Key,
		"value": settingData.Value,
	})
}

func (controller Setting) GetByKey(c *gin.Context) {

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
