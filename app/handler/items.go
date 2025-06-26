package handler

import (
	"main/app/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func AddItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category_id  := vars["category_id"]
	location_id := vars["location_id"]
	owner_id := vars["owner_id"]
	tags := vars["tags"]
	description := vars["description"]

	if i := GetItemOrNil(name, db); i != nil {
		respondError(w, http.StatusConflict, "Item already exists")
		return
	}
	category := GetCategoryById(category_id, db)
	location := GetLocationById(location_id, db)
	owner := GetUserById(owner_id, db)

	if category == nil || location == nil || owner == nil {
		respondError(w, http.StatusBadRequest, "Invalid category, location, or owner ID")
		return
	}

	item := model.Item{
		Name: name, 
		Description: description,
		Category: *category,
		Location: *location,
		Owner: *owner,
		Tags: tags,
	}
	if err := db.Save(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not save item")
		return
	}
	defer r.Body.Close()

	respondJson(w, http.StatusCreated, item)
}

func GetItems(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	var items []model.Item
	if err := db.Find(&items).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve Items")
		return
	}
	respondJson(w, http.StatusOK, items)
}

func GetItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respondError(w, http.StatusBadRequest, "Item ID is required")
		return
	}
	
	item := GetItemById(id, db)
	if item == nil {
		respondError(w, http.StatusNotFound, "Item not found")
		return
	}
	respondJson(w, http.StatusOK, item)
}

func DeleteItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respondError(w, http.StatusBadRequest, "Item ID is required")
		return
	}

	item := GetItemById(id, db)
	if item == nil {
		respondError(w, http.StatusNotFound, "Item not found")
		return
	}

	if err := db.Delete(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete item")
		return
	}

	respondJson(w, http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}

func GetItemOrNil(name string, db *gorm.DB) *model.Item {
	item := model.Item{}
	if err := db.First(&item, "name = ?", name).Error; err != nil {
		return nil
	}
	return &item
}

func GetItemById(id string, db *gorm.DB) *model.Item {
	item := model.Item{}
	if err := db.First(&item, id).Error; err != nil {
		return nil
	}
	return &item
}

func SplitTags(tags string) []string {
	if tags == "" {
		return []string{}
	}
	tagList := strings.Split(tags, ",")
	for i, tag := range tagList {
		tagList[i] = strings.TrimSpace(tag)
	}
	return tagList
}
