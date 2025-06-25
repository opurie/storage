package handlers

import (
	"main/app/handlers"
	"main/app/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AddItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category := vars["category"]
	location := vars["location"]
	owner := vars["owner"]
	tags := vars["tags"]
	description := vars["description"]

	

	if i := GetItemOrNil(name, db); i != nil {
		respondError(w, http.StatusConflict, "Item already exists")
		return
	}

	item := model.Item{Name: name, Description: description}
	if err := db.Save(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save item")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, item)
}

func GetItemOrNil(name string, db *gorm.DB) *model.Item {
	item := model.Item{}
	if err := db.First(&item, "name = ?", name).Error; err != nil {
		return nil
	}
	return &item
}

