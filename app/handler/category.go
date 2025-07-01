package handler

import (
	model "main/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	if categories, err := categoryService.GetAllCategories(); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve categories: "+err.Error())
		return
	} else {
		respondJson(w, http.StatusOK, categories)
	}
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if category, err := categoryService.GetCategory(stringToInt(id)); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve category: "+err.Error())
		return
	} else {
		respondJson(w, http.StatusOK, category)
		return
	}
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	category := model.Category{Name: name}

	if c, err := categoryService.CreateCategory(&category); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not create category: "+err.Error())
		return
	} else {
		respondJson(w, http.StatusCreated, c)
		return
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := categoryService.DeleteCategory(stringToInt(id)); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete category: "+err.Error())
		return
	}

	respondJson(w, http.StatusOK, "Category deleted successfully")
}
