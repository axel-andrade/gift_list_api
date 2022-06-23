package entities

type Gift struct {
	Base
	Name          string   `json:"name"`
	CategoryID    int64    `json:"category_id"`
	Category      Category `json:"category"`
	Available     bool     `json:"available"`
	ImageFilename string   `json:"image_filename"`
	Quantity      int64    `json:"quantity"`
	PriceGrade    int64    `json:"price_grade"`
}
