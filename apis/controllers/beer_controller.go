package controllers

import (
	"apis/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BeerController struct {
	Service *services.BeerService
}

func NewBeerController(service *services.BeerService) *BeerController {
	return &BeerController{Service: service}
}

func (c *BeerController) GetBeers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	beers, err := c.Service.GetBeers(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, beers)
}
