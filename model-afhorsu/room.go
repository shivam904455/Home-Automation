package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Room struct {
	ID                 uuid.UUID `json:"id" gorm:"not null"`
	HouseID            uuid.UUID `json:"house_id" binding:"required" gorm:"not null"`
	FloorID            uuid.UUID `json:"foolr_id" binding:"required" gorm:"not null"`
	Name               string    `json:"name" binding:"required" gorm:"not null"`
	Description        string    `json:"description"`
	Area               string    `json:"area"`
	Occupancy          string    `json:"occuoancy"`
	HeatingEnabled     string    `json:"heating_enabled"`
	HeatingType        string    `json:"heating_type"`
	CoolingEnabled     string    `json:"cooling_enabled"`
	CoolingType        string    `json:"cooling_type"`
	VantilationEnabled string    `json:"vantilation_enabled"`
	VantilationType    string    `json:"vantilation_type"`
	Teamperature       string    `json:"teamperature"`
	Humidity           string    `json:"humidity"`
	LightLavel         string    `json:"light_lavel"`
	Co2Lavel           string    `json:"co2_lavel"`
	CreatedAt          time.Time `json:"created_at" gorm:"not null"`
	CreatedBy          string    `json:"created_by" binding:"required" gorm:"not null"`
	UpdatedAt         time.Time `json:"upadated_at"`
	UpdatedBy          string    `json:"updated_by"`
}
