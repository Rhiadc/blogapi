package repositories

import "github.com/rhiadc/blogapi/models"

func NewPost(post models.Post) error {
	db := models.Connect()
	defer db.Close()

	return db.Create(&post).Error
}

func GetPosts() []models.Post {
	db := models.Connect()
	defer db.Close()
	var posts []models.Post
	db.Find(&posts)

	for i, _ := range posts {
		db.Model(posts[i]).Related(&posts[i].User)
	}
	return posts
}
