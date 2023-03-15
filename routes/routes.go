package routes

import (
	"github.com/gorilla/mux"
	"technodom_test/controllers"
)

var Router *mux.Router

func RegisterRoutes() {
	Router = mux.NewRouter()
	Router.HandleFunc("/admin/redirects", controllers.GetLinks).Methods("GET")
	Router.HandleFunc("/admin/redirects/{id}", controllers.GetLinkByID).Methods("GET")
	Router.HandleFunc("/admin/redirects", controllers.CreateLink).Methods("POST")
	Router.HandleFunc("/admin/redirects/{id}", controllers.ChangeActiveLink).Methods("PATCH")
	Router.HandleFunc("/admin/redirects/{id}", controllers.DeleteLinkById).Methods("DELETE")
	Router.HandleFunc("/redirects", controllers.GetActiveLink).Methods("GET")

}
