package main

import (
	"fmt"
	"log"
	"net/http"

	dbfunctions "example/MerchAPI/DbFunctions"

	"github.com/gin-gonic/gin"
)

type Merch struct {
	Name  string  `json: "name"`
	Price float64 `json: "price"`
	Type  string  `json: "type"`
}

var Data = []Merch{
	{Name: "Rooted in christ", Price: 499, Type: "Round Neck"},
	{Name: "One Way", Price: 699, Type: "Polo"},
}

func GetAllMerch(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data)
}

func GetMerchByName(c *gin.Context) {
	name := c.Param("name")

	for _, m := range Data {
		if m.Name == name {
			c.IndentedJSON(http.StatusFound, m)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product does not exist"})
}

func main() {
	router := gin.Default()

	router.GET("/allmerch", GetAllMerch)
	router.GET("/merch/:name", GetMerchByName)

	//Using anonymous function
	/*{router.POST("/merch", func(c *gin.Context) {
		var newMerch Merch

		err := c.BindJSON(&newMerch)

		if err == nil {
			Data = append(Data, newMerch)
			c.IndentedJSON(http.StatusCreated, Data)
		} else {
			return
		}
	}) */

	isDbConnect, err := dbfunctions.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	if isDbConnect {
		fmt.Println("Databse Connected")
	}
	router.Run("localhost:8080")
}
