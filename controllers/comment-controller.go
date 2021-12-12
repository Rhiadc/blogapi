package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
	"github.com/rhiadc/blogapi/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var comment models.Comment
	id := utils.GetID(r)
	if err := json.Unmarshal(body, &comment); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err := models.ValidateComment(comment); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.NewComment(id, comment); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "Comment successfully created!", http.StatusCreated)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)
	comments := repositories.GetComments(id)
	utils.ToJson(w, comments, http.StatusOK)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := strconv.ParseUint(vars["id"], 10, 64)
	commentId, _ := strconv.ParseUint(vars["comment_id"], 10, 64)
	err := repositories.DeleteComment(postId, commentId)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "Comment successfully deleted", http.StatusNoContent)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, _ := strconv.ParseUint(vars["id"], 10, 64)
	commentId, _ := strconv.ParseUint(vars["comment_id"], 10, 64)

	body := utils.BodyParser(r)

	var comment models.Comment

	err := json.Unmarshal(body, &comment)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	comment.PostId = uint64(postId)
	comment.ID = uint32(commentId)

	_, err = repositories.UpdateComment(comment)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "Comment successfully updated", http.StatusOK)
}
