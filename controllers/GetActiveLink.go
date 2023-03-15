package controllers

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"technodom_test/cache"
	"technodom_test/db"
	"technodom_test/middlewares"
	"time"
)

func GetActiveLink(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	cacheActiveLink, ok := cache.UserCache.Get(link)
	// Если ссылка присутствует в кэше, функция не дойдет до обращения к бд
	if ok {
		w.WriteHeader(http.StatusMovedPermanently)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Moved Permanently"
		resp["new_link"] = cacheActiveLink
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("error happened in JSON marshalling. Err: %s\n", err)
			return
		}
		w.Write(jsonResp)
		return
	}
	activeLink, err := middlewares.FindActiveLink(db.DB, link)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cache.UserCache.Add(link, activeLink, 24*time.Hour)

	w.WriteHeader(http.StatusMovedPermanently)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Moved Permanently"
	resp["new_link"] = activeLink
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(jsonResp)
}
