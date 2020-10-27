package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addAppRoutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/scrapes", getScrape).Methods("POST")
	route.HandleFunc("/products", insertProduct).Methods("POST")
	fmt.Println("Routes loading is completed")
}
