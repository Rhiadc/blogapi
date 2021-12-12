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

	if err := models.ValidatePost(post); err != nil {
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

func GetPost(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)
	post, err := repositories.GetPost(id)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.ToJson(w, post, http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	err := repositories.DeletePost(id)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusNotFound)
		return
	}
	utils.ToJson(w, "Post successfully deleted", http.StatusNoContent)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	body := utils.BodyParser(r)
	var post models.Post
	err := json.Unmarshal(body, &post)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	post.ID = uint32(id)

	_, err = repositories.UpdatePost(post)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "post successfully updated", http.StatusOK)
}
