package main

import "time"

type ErrorResponse struct {
	Message string
	Code    int
}

type CollyResponse struct {
	Name string
	ImageUrl string
	TotalReview string
	Description string
	Price string
}

type ProductResponse struct {
	Url string 			`json:"url" bson:"url,omitempty"`
	Product *Prod		`json:"product" bson:"product,omitempty"`
	Timestamp time.Time	`json:"timestamp" bson:"timestamp,omitempty"`
}

type Prod struct {
	Name string
	ImageUrl string
	TotalReview string
	Description string
	Price string
}