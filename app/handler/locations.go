package handler

import (
	model "main/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := locationService.GetAllLocations()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve locations: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	location, err := locationService.GetLocation(stringToInt(id))

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve location: "+err.Error())
		return
	}

	respondJson(w, http.StatusOK, location)
}

func AddLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	location := &model.Location{
		Name: name,
	}
	location, err := locationService.CreateLocation(location)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not create location: "+err.Error())
		return
	}

	respondJson(w, http.StatusCreated, location)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := locationService.DeleteLocation(stringToInt(id))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete location: "+err.Error())
		return
	}

	respondJson(w, http.StatusOK, map[string]string{"message": "Location deleted successfully"})
}
