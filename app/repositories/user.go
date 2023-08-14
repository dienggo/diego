package repositories

import (
	"errors"
	"github.com/dienggo/diego/app/models"
	"github.com/dienggo/diego/pkg/database"
)

type IUser interface {
	Upsert(user *models.User) error
	Delete(id uint) error
	Find(id uint) (err error, user models.User)
}

type User struct{}

func (User) Upsert(user *models.User) error {
	tx := database.Main().Where("id = ?", user.ID)
	if user.ID == 0 {
		return tx.Create(&user).Error
	} else {
		return tx.Updates(&user).Error
	}
}

func (User) Delete(id uint) error {
	tx := database.Main().Model(models.User{}).Delete("id = ?", id)
	if tx.RowsAffected == 0 {
		tx.Error = errors.New("no data affected")
	}
	return tx.Error
}

func (User) Find(id uint) (err error, user models.User) {
	err = database.Main().Where("id = ?", id).First(&user).Error
	return err, user
}
