package user_service_store

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_base_project/app/models"
	"go_base_project/app/services"
	"go_base_project/mocks"
	"testing"
)

func TestUserServiceStoreSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	gomock.InOrder(
		mockRepo.EXPECT().Create(gomock.Eq("Dipa Ferdian")).Return(models.User{Id: 1, Name: "Dipa Ferdian"}),
	)

	call := services.NewUser(mockRepo)
	actual := call.Store("Dipa Ferdian")

	expected := services.UserEntity{
		Id:   1,
		Name: "Dipa Ferdian",
	}

	assert.Equal(t, expected, actual)
}
