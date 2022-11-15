package user_repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_base_project/app/models"
	"go_base_project/mocks"
	"testing"
)

func TestUser_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	mockRepo.EXPECT().Find(gomock.Eq(1)).Return(models.User{Id: 1, Name: "Dipa Ferdian"})

	call := mockRepo.Find(1)

	expect := models.User{
		Id:   1,
		Name: "Dipa Ferdian",
	}

	assert.Equal(t, expect, call)
}
