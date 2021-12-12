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

func GetPost(id uint64) (*models.Post, error) {
	db := models.Connect()
	defer db.Close()

	var post models.Post

	if err := db.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func DeletePost(id uint64) error {
	db := models.Connect()
	defer db.Close()

	var post models.Post

	if err := db.Where("id = ?", id).Delete(&post).Error; err != nil {
		return err
	}

	return nil
}
