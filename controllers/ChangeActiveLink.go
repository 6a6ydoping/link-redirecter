package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"technodom_test/db"
	"technodom_test/middlewares"
)

func ChangeActiveLink(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramID := params["id"]
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	var requestBody map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.ParseForm()
	newActiveLink := fmt.Sprintf("%v", requestBody["active_link"])
	if err = middlewares.ChangeActiveLink(db.DB, uint(id), newActiveLink); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
