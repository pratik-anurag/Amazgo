package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server will start at http://localhost:3000/")

	route := mux.NewRouter()

	addAppRoutes(route)

	log.Fatal(http.ListenAndServe(":3000", route))
}
