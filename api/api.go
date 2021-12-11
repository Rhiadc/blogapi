package api

import (
	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
)

func Run() {
	//models.AutoMigrations()
	repositories.NewUser(models.User{Nickname: "Fsparza", Email: "f@esparza.com", Password: "esparzinha.com"})

}
