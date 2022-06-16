package entities

// "id": 1,
// "name": "Chuveiro Elétrico",
// "category_id": 1,
// "available": true,
// "image_filename": "chuveiro",
// "quantity": 1

type Gift struct {
	Base
	Name          string `json:"name"`
	CategoryID    int64  `json:"category_id"`
	Available     bool   `json:"available"`
	ImageFilename string `json:"image_filename"`
	Quantity      int64  `json:"quantity"`
}
