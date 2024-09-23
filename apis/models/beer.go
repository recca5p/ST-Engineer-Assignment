package models

type Rating struct {
	Average float64 `json:"average"`
	Reviews int     `json:"reviews"`
}

type Beer struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Price         string  `json:"price"`
	Name          string  `json:"name"`
	AverageRating float64 `json:"average_rating"`
	Reviews       int     `json:"reviews"`
	Image         string  `json:"image"`
}
