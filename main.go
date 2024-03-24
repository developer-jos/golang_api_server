package main

import (
	"net/http"

	"server/api"
	//"github.com/your-username/your-project/api" // Adjust the import path as per your project structure
)

func main() {
	server := api.NewServer(&api.InMemoryDatabase{})

	http.HandleFunc("/items", api.HandlerFunc(server.GetItems))
	http.HandleFunc("/items/create", api.HandlerFunc(server.CreateItem))
	http.HandleFunc("/items/get", api.HandlerFunc(server.GetItem))
	http.HandleFunc("/items/delete", api.HandlerFunc(server.DeleteItem))

	http.ListenAndServe(":8080", nil)
}
