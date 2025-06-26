package handler

import (
	"main/app/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AddUser(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	email := vars["email"]
	description := vars["description"]
	
	if u := GetUserOrNil(username, db); u != nil {
		respondError(w, http.StatusConflict, "User already exists")
		return
	}

	user := model.User{Username: username, Email: email, Description: description}
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save user")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	email := vars["email"]
	description := vars["description"]

	user := model.User{}
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	user.Email = email
	user.Description = description
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not update user")
		return
	}

	defer r.Body.Close()
	respondJson(w, http.StatusOK, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]

	user := model.User{}
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	respondJson(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]

	user := GetUserOrNil(username, db)
	if user == nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete user")
		return
	}
	respondJson(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func GetUserOrNil(name string, db *gorm.DB) *model.User{
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