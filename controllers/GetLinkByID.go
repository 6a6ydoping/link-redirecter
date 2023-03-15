package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"technodom_test/cache"
	"technodom_test/db"
	"technodom_test/middlewares"
	"technodom_test/models"
)

func GetLinkByID(w http.ResponseWriter, r *http.Request) {
	fmt.Print(cache.UserCache.Get("ASD"))
	params := mux.Vars(r)
	paramID := params["id"]
	id, err := strconv.Atoi(paramID)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	var link *models.Link
	link, err = middlewares.GetLinkByID(db.DB, id)
	fmt.Fprintf(w, "ID: %d\n", link.ID)
	fmt.Fprintf(w, "Active Link: %s\n", link.ActiveLink)
	fmt.Fprintf(w, "History Link: %s\n", link.HistoryLink)
	fmt.Fprintf(w, "\n")
}
