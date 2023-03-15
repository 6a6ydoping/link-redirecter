package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"technodom_test/db"
	"technodom_test/middlewares"
)

func DeleteLinkById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramID := params["id"]
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		fmt.Println("Error during conversion")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	middlewares.DeleteLinkById(db.DB, uint(id))
	w.WriteHeader(http.StatusOK)
}
