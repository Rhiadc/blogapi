package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
	"github.com/rhiadc/blogapi/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := repositories.NewUser(user); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, "User successfully created!", http.StatusCreated)
}
