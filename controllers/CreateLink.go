package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"technodom_test/db"
	"technodom_test/middlewares"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.ParseForm()
	activeLink := fmt.Sprintf("%v", requestBody["active_link"])
	historyLink := fmt.Sprintf("%v", requestBody["history_link"])
	if err = middlewares.CreateLink(db.DB, activeLink, historyLink); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
