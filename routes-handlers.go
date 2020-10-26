package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)
func getScrape(response http.ResponseWriter, request *http.Request){
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occurred",
	}
	url := mux.Vars(request)["url"]
	if url == "" {
		httpError.Message = "url can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		//call logic on url
	}
}



func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Message: errorMesage.Message,Code: errorMesage.Code}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}