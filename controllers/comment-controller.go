package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
	"github.com/rhiadc/blogapi/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var comment models.Comment
	if err := json.Unmarshal(body, &comment); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.NewComment(comment); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "Comment successfully created!", http.StatusCreated)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	comments := repositories.GetComments()
	utils.ToJson(w, comments, http.StatusOK)
}
