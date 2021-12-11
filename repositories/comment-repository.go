package repositories

import "github.com/rhiadc/blogapi/models"

func NewComment(comment models.Comment) error {
	db := models.Connect()
	defer db.Close()

	return db.Create(&comment).Error
}

func GetComments() []models.Comment {
	db := models.Connect()
	defer db.Close()
	var comments []models.Comment
	db.Find(&comments)

	return comments
}
