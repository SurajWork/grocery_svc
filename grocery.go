package grocery_svc

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity uint32 `json:"quantity"`
}
