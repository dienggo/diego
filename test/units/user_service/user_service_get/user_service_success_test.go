package user_service_get

import (
	_ "fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"go_base_project/app/models"
	"go_base_project/app/services"
	_ "go_base_project/app/services"
	"go_base_project/mocks"
	"testing"
)

func TestUserServiceGetSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoriesInterface(ctrl)

	gomock.InOrder(
		mockRepo.EXPECT().Find(gomock.Eq(1)).Return(models.User{Id: 1, Name: "Dipa Ferdian"}),
	)

	call := services.NewUser(mockRepo)
	actual := call.Get(1)

	expected := services.UserEntity{
		Id:   1,
		Name: "Dipa Ferdian",
	}

	assert.Equal(t, expected, actual)
}
