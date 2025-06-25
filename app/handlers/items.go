package handlers

import (
	"main/app/handlers"
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

	category := handlers.GetCategoryById(category_id, db)
	location := handlers.GetLocationById(location_id, db)
	owner := handlers.GetUserOrNil(owner_id, db)

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
		Tags: SplitTags(tags),
	}
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
