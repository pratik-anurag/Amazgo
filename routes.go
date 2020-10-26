package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}

func addAppRoutes(route *mux.Router) {

	setStaticFolder(route)

	route.HandleFunc("/scrapes/{url}", getScrape).Methods("POST")
	fmt.Println("Routes loading is completed")
}
