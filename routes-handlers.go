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
	if url == "https://www.amazon.in/s?k=ps2&ref=nb_sb_noss_2" {
		httpError.Message = "url can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		callColly(url)
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMsg ErrorResponse) {
	httpResponse := &ErrorResponse{Message: errorMsg.Message,Code: errorMsg.Code}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMsg.Code)
	response.Write(jsonResponse)
}