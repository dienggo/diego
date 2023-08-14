package controllers

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/dto_response"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/app/services"
	"github.com/dienggo/diego/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

// View : to show data detail on User
// Example no effort logic
func (ctrl User) View(c *gin.Context) {
	err, user := repositories.User{}.Find(helper.StringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Loaded",
		"user":    dto_response.User(user),
	})
}

// Upsert : to update/insert data on User
// Example execute logic in service
func (ctrl User) Upsert(c *gin.Context) {
	var req dto_request.User
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	service := services.NewUpsertUser(req).Do()
	if service.Error() != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": service.Error().Error(),
		})
		return
	}

	var user *models.User
	err = service.ResultParse(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Inserted/Updated data",
		"user":    dto_response.User(*user),
	})
}

// Delete : to delete data on User
func (ctrl User) Delete(c *gin.Context) {
	err := repositories.User{}.Delete(helper.StringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User deleted",
	})
}
