package handler
import (
	"encoding/json"
	"net/http"
)

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