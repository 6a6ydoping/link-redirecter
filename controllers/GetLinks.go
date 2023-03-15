package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"technodom_test/db"
	"technodom_test/middlewares"
	"technodom_test/models"
)

func GetLinks(w http.ResponseWriter, r *http.Request) {
	paramPage := r.URL.Query().Get("page")
	paramPageSize := r.URL.Query().Get("size")

	if paramPage == "" {
		paramPage = "1"
	}
	if paramPageSize == "" {
		paramPageSize = "100"
	}
	page, err := strconv.Atoi(paramPage)
	if err != nil {
		fmt.Println("Error during conversion")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(paramPageSize)
	if err != nil {
		fmt.Println("Error during conversion")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var links []models.Link
	links, err = middlewares.GetAllLinks(db.DB, page, pageSize)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for _, link := range links {
		fmt.Fprintf(w, "ID: %d\n", link.ID)
		fmt.Fprintf(w, "Active Link: %s\n", link.ActiveLink)
		fmt.Fprintf(w, "History Link: %s\n", link.HistoryLink)
		fmt.Fprintf(w, "\n")
	}
}
