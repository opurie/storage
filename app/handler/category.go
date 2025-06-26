package handler

import (
	"main/app/models"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve categories")
		return 
	}
	respondJson(w, http.StatusOK, categories)
}

func GetCategoryByName(name string, db *gorm.DB) *model.Category {
	category := model.Category{}
	if err := db.First(&category, "name = ?", name).Error; err != nil {
		return nil
	}
	return &category
}

func GetCategoryById(id string, db *gorm.DB) *model.Category {
	category := model.Category{}
	if err := db.First(&category, id).Error; err != nil {
		return nil
	}
	return &category
}

func AddCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if category := GetCategoryByName(name, db); category != nil {
		respondError(w, http.StatusConflict, "Category already exists")
		return
	}
	category := model.Category{Name: name}

	if err := db.Save(&category).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save category")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, category)
}