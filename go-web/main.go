package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	// internal 
	"golang-web-module/internal/services"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todo", services.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id}", services.GetTodoById).Methods(http.MethodGet) 

	log.Fatal(http.ListenAndServe(":8080", r))
}