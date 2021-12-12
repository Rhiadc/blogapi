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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repositories.GetUsers()
	utils.ToJson(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)
	user, err := repositories.GetUser(id)

	if err != nil {
		utils.ToJson(w, "Not found", http.StatusNotFound)
		return
	}

	utils.ToJson(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	err := repositories.DeleteUser(id)

	if err != nil {
		utils.ToJson(w, "Not found", http.StatusNotFound)
		return
	}

	utils.ToJson(w, "User successfully deleted", http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, user)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user.ID = uint32(id)

	rows, err := repositories.UpdateUser(user)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, rows, http.StatusOK)
}
