package database

import "github.com/rhiadc/blogapi/models"

func AutoMigrations() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&models.Comment{}, &models.Post{}, &models.User{})
	db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	db.Debug().Model(&models.Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&models.Comment{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
