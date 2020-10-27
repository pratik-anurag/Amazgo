package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func callColly(url string)[]byte{
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	var (
		outputDetail CollyResponse
		outputDetails []CollyResponse
		)
	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName, imageUrl, description, price, totalReviews string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}

			imageUrl = e.ChildAttr("img.s-image","src")

			totalReviews = e.ChildText("span.a-size-base")

			description = e.ChildText("span.a-size-base.a-color-secondary")

			description = strings.ReplaceAll(description,totalReviews,"")

			price = e.ChildText("span.a-price > span.a-offscreen")

			FormatPrice(&price)

			outputDetail = CollyResponse{
				Name:productName,ImageUrl:imageUrl,TotalReview:totalReviews,Description:description,Price:price,
			}

			//var newProduct ProductResponse
			//var newProd Prod
			//newProduct.Url = url
			//newProd.Name = outputDetail.Name
			//newProd.ImageUrl = outputDetail.ImageUrl
			//newProd.TotalReview = outputDetail.TotalReview
			//newProd.Description = outputDetail.Description
			//newProd.Price = outputDetail.Price
			//newProduct.Product = &newProd
			//newProduct.Timestamp = time.Now()
			//collection := ConnectDB()
			//_, err := collection.InsertOne(context.TODO(), newProduct)
			//if err != nil {
			//	return
			//}
			outputDetails = append(outputDetails,outputDetail)
		})
	})
	c.Visit(url)

	jsonResponse, jsonError := json.Marshal(outputDetails)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func FormatPrice(price *string) {
	r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)

	newPrices := r.FindStringSubmatch(*price)

	if len(newPrices) > 1 {
		*price = newPrices[1]
	} else {
		*price = "Unknown"
	}
}