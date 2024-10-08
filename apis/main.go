package main

import (
	"apis/controllers"
	"apis/database"
	"apis/services"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	beerService := services.NewBeerService(database.DB)
	beerController := controllers.NewBeerController(beerService)

	r.GET("/beers", beerController.GetBeers)
	r.GET("/beers/:id", beerController.GetBeerByID)

	r.Run(":8081")
}
