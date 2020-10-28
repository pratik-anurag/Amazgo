package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
func getScrape(response http.ResponseWriter, request *http.Request){
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occurred",
	}
	url:= request.URL.Query().Get("url")
	if url == "" {
		httpError.Message = "url can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		jsonResponse := callColly(url)
		if jsonResponse == nil {
			returnErrorResponse(response, request, httpError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
			insertManyProduct(jsonResponse,url,response, request)
		}
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

func insertProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product ProductResponse
	err1 := json.NewDecoder(r.Body).Decode(&product)
	if err1 != nil {
		fmt.Println("error in build body: ", err1)
	}
	product.Timestamp = time.Now()
	url:= r.URL.Query().Get("url")
	product.Url = url
	collection := ConnectDB()
	result, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func insertManyProduct(jsonResponse []byte,url string,w http.ResponseWriter, r *http.Request){
	var product []ProductResponse
	err1:= json.Unmarshal(jsonResponse, &product)
	if err1 != nil {
		fmt.Println("error in build body: ", err1)
	}
	for i := range product {
		doc := product[i]
		doc.Timestamp = time.Now()
		doc.Url = url
		collection := ConnectDB()
		result, insertErr := collection.InsertOne(context.TODO(), doc)
		if insertErr != nil {
			GetError(insertErr, w)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}