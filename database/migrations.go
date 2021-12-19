package database

import "github.com/rhiadc/blogapi/models"

var (
	name = "blogapi"
)

func AutoMigrations() {
	db := Connect()

	defer db.Close()
	db.Exec("CREATE DATABASE " + name)
	db.Exec("USE " + name)

	db.Debug().DropTableIfExists(&models.Comment{}, &models.Post{}, &models.User{})
	db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	db.Debug().Model(&models.Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
