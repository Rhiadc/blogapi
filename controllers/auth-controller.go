package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rhiadc/blogapi/auth"
	"github.com/rhiadc/blogapi/models"
	"github.com/rhiadc/blogapi/repositories"
	"github.com/rhiadc/blogapi/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	givenUser, err := repositories.GetByEmail(user.Email)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err := auth.VerifyPassword(givenUser.Password, user.Password); err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(givenUser.ID)

	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	saveErr := auth.CreateAuth(givenUser.ID, token)
	if saveErr != nil {
		fmt.Println("asdasd")
		return
	}
	resp := utils.TokenResponse{Token: token.AccessToken, RefreshToken: token.RefreshToken}

	utils.ToJson(w, resp, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	au, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	deleted, delErr := auth.DeleteAuth(au.AccessUuid)

	if delErr != nil || deleted == 0 {
		utils.ToJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.ToJson(w, "Successfully logged out", http.StatusOK)
}
