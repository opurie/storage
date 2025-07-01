package handler

import (
	"main/app/services"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

var userService *services.UserService
var categoryService *services.CategoryService
var itemsService *services.ItemsService
var locationService *services.LocationService

func InitServices(db *gorm.DB) {
	userService = services.NewUserService(db)
	categoryService = services.NewCategoryService(db)
	itemsService = services.NewItemsService(db)
	locationService = services.NewLocationService(db)
}


func respondJson(w http.ResponseWriter, status int, payload interface{}){
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJson(w, status, map[string]string{"error": message})
}


//no need for error handling here, as this is a utility function
// it will return -1 if the string is empty or cannot be converted to an int
func stringToInt(s string) int {
	if s == "" {
		return -1
		}
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}