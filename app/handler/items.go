package handler

import (
	"main/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category_id  := vars["category_id"]
	location_id := vars["location_id"]
	owner_id := vars["owner_id"]
	tags := vars["tags"]
	description := vars["description"]


	item := model.Item{
		Name: name, 
		Description: description,
		CategoryID: uint(stringToInt(category_id)),
		LocationID: uint(stringToInt(location_id)),
		UserID: uint(stringToInt(owner_id)),
		Tags: tags,
	}
	if err := itemsService.CreateItem(&item); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not create item: "+err.Error())
		return
	}
	defer r.Body.Close()
	respondJson(w, http.StatusCreated, "Item created successfully")

}

func GetItems(w http.ResponseWriter, r *http.Request){
	
	items, err := itemsService.GetAllItems()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve items: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := itemsService.GetItemById(stringToInt(id))
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Could not retrieve item: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := itemsService.DeleteItem(stringToInt(id)); err != nil {
		respondError(w, http.StatusInternalServerError, "Could not delete item: "+err.Error())
		return
	}
	respondJson(w, http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}
