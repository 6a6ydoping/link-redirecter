package main

import (
	"net/http"
	"technodom_test/cache"
	"technodom_test/db"
	"technodom_test/routes"
)

func init() {
	db.ConnectToDataBase()
	db.SyncDB()
	routes.RegisterRoutes()
}

func main() {
	cache.UserCache = cache.NewCache(1000)
	go cache.UserCache.WarmUpCache()
	http.ListenAndServe("localhost:8080", routes.Router)
}
