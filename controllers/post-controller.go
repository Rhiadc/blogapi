package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
	"github.com/rhiadc/blogapi/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.NewPost(post); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "Post successfully created!", http.StatusCreated)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := repositories.GetPosts()
	utils.ToJson(w, posts, http.StatusOK)
}
