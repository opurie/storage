package handler

import (
	"io"
	model "main/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	email := vars["email"]
	description, _ := io.ReadAll(r.Body)

	if user, err := userService.CreateUser(&model.User{
		Username:    username,
		Email:       email,
		Description: string(description),
	}); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not Create user: "+err.Error())
		return
	} else {
		respondJson(w, http.StatusCreated, user)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	username := vars["username"]
	email := vars["email"]
	description, _ := io.ReadAll(r.Body)

	user := model.User{
		Username:    username,
		Email:       email,
		Description: string(description),
	}

	_, err := userService.UpdateUser(stringToInt(id), &user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not update user: "+err.Error())
		return
	}

	defer r.Body.Close()
	respondJson(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetAllUsers()

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve users: "+err.Error())
		return
	}

	respondJson(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := userService.GetUser(stringToInt(id))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve user: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := userService.DeleteUser(stringToInt(id))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete user: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
