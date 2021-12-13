package database

import "github.com/rhiadc/blogapi/models"

var (
	DBNAME = "blogapi"
)

func AutoMigrations() {
	db := Connect()

	defer db.Close()
	db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")
	db.Exec("USE " + DBNAME)

	db.Debug().DropTableIfExists(&models.Comment{}, &models.Post{}, &models.User{})
	db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	db.Debug().Model(&models.Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
