package api

import (
	"github.com/rhiadc/blogapi/database"
	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
)

func Run() {
	database.AutoMigrations()
	repositories.NewUser(models.User{Nickname: "Fsparza", Email: "f@esparza.com", Password: "esparzinha.com"})
	// todo:
	// 	- testes
	// 	- jwt e autenticação
	// 	- env
	// 	- docker e kubernetes
	// 	- deploy na aws
}
