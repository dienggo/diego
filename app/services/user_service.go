package services

import (
	"go_base_project/app/repositories"
	_ "go_base_project/app/repositories"
	"go_base_project/app/response"
	"reflect"
	"regexp"
	"strconv"
)

type UserService struct {
	Repositories repositories.UserRepositoriesInterface
}

type UserEntity struct {
	Id   uint
	Name string
}

func NewUser(p repositories.UserRepositoriesInterface) *UserService {
	return &UserService{Repositories: p}
}

func (u UserService) Get(Id any) interface{} {
	var userId = Id

	if reflect.TypeOf(Id) == reflect.TypeOf("") {
		isInteger, _ := regexp.MatchString("[0-9]", userId.(string))

		if !isInteger {
			return response.Service().Error("Params must be in integer", nil)
		}

		convertInteger, _ := strconv.Atoi(Id.(string))
		userId = convertInteger
	}

	if _, isInteger := userId.(int); !isInteger {
		return response.Service().Error("Params must be in integer", nil)
	}

	userRepo := u.Repositories.Find(Id)

	if userRepo.Id == 0 {
		return response.Service().Error("user not found", nil)
	}

	return UserEntity{
		Id:   userRepo.Id,
		Name: userRepo.Name,
	}
}

func (u UserService) Store(name any) interface{} {

	if _, isString := name.(string); !isString {
		return response.Service().Error("Params name must be in string", nil)
	}

	userRepo := u.Repositories.Create(name)

	if userRepo.Name == "" {
		return response.Service().Error("Failed Insert Data", nil)
	}
	
	return UserEntity{
		Id:   userRepo.Id,
		Name: userRepo.Name,
	}
}
