package models

func AutoMigrations() {
	db := Connect()
	defer db.Close()

	db.Debug().DropTableIfExists(&Comment{}, &Post{}, &User{})
	db.Debug().AutoMigrate(&User{}, &Post{}, &Comment{})
	db.Debug().Model(&Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&Comment{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	db.Debug().Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
