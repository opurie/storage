package handler

import (
	"main/app/models"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetLocations(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var locations []model.Location
	if err := db.Find(&locations).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve locations")
		return 
	}
	respondJson(w, http.StatusOK, locations)
}

func GetLocationById(id string, db *gorm.DB) *model.Location {
	location := model.Location{}
	if err := db.First(&location, id).Error; err != nil {
		return nil
	}
	return &location
}

func AddLocation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if location := db.First(&model.Location{}, "name = ?", name ); location != nil {
		respondError(w, http.StatusConflict, "Location already exists")
		return
	}
	location := model.Location{Name: name}

	if err := db.Save(&location).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save location")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, location)
}