package presents

import (
	"github.com/dienggo/diego/app/models"
	"time"
)

type user struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func User(mUser models.User) user {
	return user{
		Id:        mUser.ID,
		Name:      mUser.Name,
		Email:     mUser.Email,
		CreatedAt: mUser.CreatedAt,
	}
}

func Users(mUsers []models.User) (users []user) {
	for _, mUser := range mUsers {
		users = append(users, User(mUser))
	}
	return users
}
