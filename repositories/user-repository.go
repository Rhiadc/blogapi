package repositories

import (
	"github.com/rhiadc/blogapi/auth"
	"github.com/rhiadc/blogapi/database"
	"github.com/rhiadc/blogapi/models"
)

type UserRepo interface {
	NewUser(user models.User) error
}

func NewUser(user models.User) error {
	db := database.Connect()
	defer db.Close()
	var err error
	user.Password, err = auth.Hash(user.Password)

	if err != nil {
		return err
	}
	return db.Create(&user).Error
}

func GetUsers() []models.User {
	db := database.Connect()
	defer db.Close()

	var user []models.User

	db.Find(&user)

	for i, _ := range user {
		db.Model(user[i]).Related(&user[i].Posts)
	}

	return user
}

func GetUser(id uint64) (*models.User, error) {
	db := database.Connect()
	defer db.Close()

	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(id uint64) error {
	db := database.Connect()
	defer db.Close()

	var user models.User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *models.User) (int64, error) {
	db := database.Connect()
	defer db.Close()

	rs := db.Model(&user).Where("id = ?", user.ID).UpdateColumns(
		map[string]interface{}{
			"nickname": user.Nickname,
			"email":    user.Email,
		},
	)
	return rs.RowsAffected, rs.Error
}

func GetByEmail(email string) (*models.User, error) {
	db := database.Connect()
	defer db.Close()

	var user models.User

	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
