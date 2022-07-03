package http_example_service

import (
	"go_base_project/app/models"
	"go_base_project/app/repositories"
	"testing"
)

func TestHttpExample_Do(t *testing.T) {
	service := HttpExample{repositories.HttpExampleRepository{}}.Do()

	if service.Status == true {
		for _, value := range service.Data.(models.Posts) {
			println("title :", value.Title)
		}
	}
	println(service.Message)
}