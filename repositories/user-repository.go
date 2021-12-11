package repositories

import (
	"github.com/rhiadc/blogapi/auth"
	"github.com/rhiadc/blogapi/models"
)

type UserRepo interface {
	NewUser(user models.User) error
}

func NewUser(user models.User) error {
	db := models.Connect()
	defer db.Close()
	var err error
	user.Password, err = auth.Hash(user.Password)

	if err != nil {
		return err
	}
	return db.Create(&user).Error
}
