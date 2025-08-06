package main

import (
	"fmt"
	"log"
	"net/http"

	dbfunctions "example/MerchAPI/DbFunctions"

	"github.com/gin-gonic/gin"
)

type Merch struct {
	Name     string  `json: "name"`
	Price    float64 `json: "price"`
	Type     string  `json: "type"`
	Size     string  `json: "size"`
	Quantity int64   `json: quantity`
}

var Data = []Merch{
	{Name: "Rooted in christ", Price: 499, Type: "Round Neck"},
	{Name: "One Way", Price: 699, Type: "Polo"},
}

func GetAllMerch(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Data)
}

func GetAllMerchFromDb(c *gin.Context) {
	data, err := dbfunctions.GetAllMerchWithQuantity()
	if err != nil {
		log.Fatal(err)
		// handle for 0 rows
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
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

func PostMerchtoDb(c *gin.Context) {
	var newMerch dbfunctions.Merch

	err := c.BindJSON(&newMerch)

	if err == nil {
		stockId, err := dbfunctions.AddMerchToDb(newMerch)
		if err != nil {
			log.Fatal(err)
		}
		if stockId != 0 {
			c.IndentedJSON(http.StatusCreated, stockId)
		}
	} else {
		return
	}
}

func main() {
	router := gin.Default()

	router.GET("/allmerch", GetAllMerch)
	router.GET("/merch/:name", GetMerchByName)

	isDbConnect, err := dbfunctions.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	if isDbConnect {
		fmt.Println("Databse Connected")
	}

	router.GET("/merchwithquantity", GetAllMerchFromDb)
	router.POST("/addmerchtodb", PostMerchtoDb)
	router.Run("localhost:8080")
}
