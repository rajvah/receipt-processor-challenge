package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const parameter = "id"

type Point struct {
	Points int64 `json:"points"`
}

type Receipt struct {
	ID           string `json:"ID"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var reciepts = map[string]Receipt{}

var points = map[string]Point{}

func CreateReciept() gin.HandlerFunc {
	return func(context *gin.Context) {
		var newReceipt Receipt
		if err := context.BindJSON(&newReceipt); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		//Generate a student ID
		newReceipt.ID = uuid.New().String()
		point, err := calculatePoints(&newReceipt)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		reciepts[newReceipt.ID] = newReceipt
		points[newReceipt.ID] = point
		context.JSON(http.StatusCreated, gin.H{"id": newReceipt.ID})
	}
}

func GetPoints() gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param(parameter)

		val, ok := points[id]
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": fmt.Sprintf("ID %s not found ", id)}})
			return
		}
		context.JSON(http.StatusCreated, val)
	}
}
