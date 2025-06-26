package handler

import (
	"log"
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

func GetLocation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respondError(w, http.StatusBadRequest, "Location ID is required")
		return
	}
	
	location := GetLocationById(id, db)
	if location == nil {
		respondError(w, http.StatusNotFound, "Location not found")
		return
	}
	respondJson(w, http.StatusOK, location)
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
	if name == "" {
		respondError(w, http.StatusBadRequest, "Location name is required")
		return
	}
	log.Print("Adding location with name:", name)
	if l := db.Find(&model.Location{}, "name = ?", name); l.RowsAffected > 0 {
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

func DeleteLocation(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respondError(w, http.StatusBadRequest, "Location ID is required")
		return
	}

	location := GetLocationById(id, db)
	if location == nil {
		respondError(w, http.StatusNotFound, "Location not found")
		return
	}

	if err := db.Delete(&location).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete location")
		return
	}

	respondJson(w, http.StatusOK, map[string]string{"message": "Location deleted successfully"})
}