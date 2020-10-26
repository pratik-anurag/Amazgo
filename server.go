package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:8000/")

	ConnectDB()

	route := mux.NewRouter()

	addAppRoutes(route)

	log.Fatal(http.ListenAndServe(":8000", route))
}
