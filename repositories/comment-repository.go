package repositories

import (
	"github.com/rhiadc/blogapi/database"
	"github.com/rhiadc/blogapi/models"
)

func NewComment(postId uint64, comment models.Comment) error {
	db := database.Connect()
	defer db.Close()
	comment.PostId = postId
	return db.Create(&comment).Error
}

func GetComments(postId uint64) []models.Comment {
	db := database.Connect()
	defer db.Close()
	var comments []models.Comment
	db.Where("post_id = ?", postId).Find(&comments)

	for i, _ := range comments {
		db.Model(&comments[i]).Related(&comments[i].Post)
	}

	return comments
}

func DeleteComment(postId uint64, commentId uint64) error {
	db := database.Connect()
	defer db.Close()
	var comment models.Comment

	err := db.Where("post_id = ? and id =?", postId, commentId).Delete(&comment).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateComment(comment models.Comment) (int64, error) {
	db := database.Connect()
	defer db.Close()

	rs := db.Model(&comment).Where("id = ?", comment.ID).UpdateColumns(
		map[string]interface{}{
			"comment": comment.Comment,
		},
	)

	return rs.RowsAffected, rs.Error
}
