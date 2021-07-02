package router

import (
	"ToDo/views"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", views.CreateTask).Methods("POST")

	return router
}
