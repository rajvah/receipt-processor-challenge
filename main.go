package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", CreateReciept())

	router.GET("/receipts/:id/points", GetPoints())

	// starts the server at port 8080x
	err := router.Run(":8080")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
}
