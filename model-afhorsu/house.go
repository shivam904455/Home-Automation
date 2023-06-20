package model

import (
	"time"
 uuid "github.com/google/uuid"
)




type house struct {
	ID             uuid.UUID `json:"id" gorm:"primarykey" `
	Name           string    `json:"name" binding:"required" gorm:"not null"`
	Description    string    `json:"description"`
	Area           string    `json:"area"`
	Occupancy      string    `json:"occupancy"`
	NumberOfFloors int       `json:"number_of_floors"`
	CreatedAt      time.Time `json:"created_at"gorm:"not null"`
	CreatedBy      string    `json:"created_by"binding:"required" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
}
