package services

import (
	"apis/models"
	"gorm.io/gorm"
)

type BeerService struct {
	DB *gorm.DB
}

func NewBeerService(db *gorm.DB) *BeerService {
	return &BeerService{DB: db}
}

func (s *BeerService) GetBeers(page, limit int) ([]models.Beer, error) {
	var beers []models.Beer
	offset := (page - 1) * limit
	err := s.DB.Offset(offset).Limit(limit).Find(&beers).Error
	return beers, err
}
