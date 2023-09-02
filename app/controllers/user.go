package controllers

import (
	"github.com/dienggo/diego/app/dto_request"
	"github.com/dienggo/diego/app/dto_response"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/app/repositories"
	"github.com/dienggo/diego/app/services"
	"github.com/dienggo/diego/pkg/app"
	"github.com/dienggo/diego/pkg/helper"
	"github.com/dienggo/diego/pkg/render"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct{}

// View : to show data detail on User
// Example no effort logic
func (ctrl User) View(w http.ResponseWriter, r *http.Request) {
	err, user := repositories.User{}.Find(helper.StringToUint(mux.Vars(r)["id"]))
	if err != nil {
		render.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}

	render.Json(w, 200, map[string]any{
		"message": "Loaded",
		"user":    dto_response.User(user),
	})
}

// Upsert : to update/insert data on User
// Example execute logic in service
func (ctrl User) Upsert(w http.ResponseWriter, r *http.Request) {
	var req dto_request.User
	err := app.NewHttpProcessor(r).Cast(&req)
	if err != nil {
		render.Json(w, http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}

	service := services.NewUpsertUser(req).Do()
	if service.Error() != nil {
		render.Json(w, http.StatusBadRequest, map[string]any{
			"message": service.Error().Error(),
		})
		return
	}

	var user *models.User
	err = service.ResultParse(&user)
	if err != nil {
		render.Json(w, http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}

	render.Json(w, http.StatusOK, map[string]any{
		"message": "Inserted/Updated data",
		"user":    dto_response.User(*user),
	})
}

// Delete : to delete data on User
func (ctrl User) Delete(w http.ResponseWriter, r *http.Request) {
	err := repositories.User{}.Delete(helper.StringToUint(mux.Vars(r)["id"]))
	if err != nil {
		render.Json(w, http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
		return
	}
	render.Json(w, http.StatusOK, map[string]any{
		"message": "User deleted",
	})
}
