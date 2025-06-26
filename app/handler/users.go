package handler

import (
	"io"
	"main/app/models"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	email := vars["email"]
	
	description, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not read request body")
		return
	}
	if username == "" || email == "" {
		respondError(w, http.StatusBadRequest, "Username and email are required")
		return
	}
	if u := GetUserOrNil(username, db); u != nil {
		respondError(w, http.StatusConflict, "User already exists")
		return
	}

	user := model.User{Username: username, Email: email, Description: string(description)}
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save user")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		respondError(w, http.StatusBadRequest, "User ID is required")
		return
	}
	username := vars["username"]
	email := vars["email"]
	description, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not read request body")
		return
	}

	user := GetUserById(id, db)
	if user == nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
	if description != nil && len(description) > 0 {
		user.Description = string(description)
	}
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not update user")
		return
	}

	defer r.Body.Close()
	respondJson(w, http.StatusOK, user)
}

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve Users")
		return
	}
	respondJson(w, http.StatusOK, users)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		respondError(w, http.StatusBadRequest, "User ID is required")
		return
	}
	user := GetUserById(id, db)
	if user == nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}


	respondJson(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		respondError(w, http.StatusBadRequest, "User ID is required")
		return
	}
	user := GetUserById(id, db)
	if user == nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := db.Unscoped().Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete user")
		return
	}
	respondJson(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func GetUserOrNil(name string, db *gorm.DB) *model.User {
	user := model.User{}
	if err := db.First(&user, "username = ?", name).Error; err != nil {
		return nil
	}
	return &user

}

func GetUserById(id string, db *gorm.DB) *model.User {
	user := model.User{}
	if err := db.First(&user, id).Error; err != nil {
		return nil
	}
	return &user
}
