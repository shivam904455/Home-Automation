package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Actuator struct {
	ID                 uuid.UUID `json:"id" gorm:"primarykey"`
	RoomsID            uuid.UUID `json:"rooms_id" binding:"requird"  gorm:"not null"`
	HouseID            uuid.UUID `json:"house_id" binding:"requird"  gorm:"not null"`
	FloorID            uuid.UUID `json:"floor_id" binding:"requird"  gorm:"not null"`
	Name               string    `json:"name" binding:"requird"  gorm:"not null"`
	Description        string    `josn:"description"`
	Area               string    `json:"area"`
	WattConsumption    string    `json:"watt_consumption"`
	Occupancy          string    `json:"Occupancy"`
	HeatingEnabled     bool      `json:"heating_enabled"`
	HeatingType        string    `json:"heating-type"`
	CoolingEnabled     bool      `json:"cooling_enabled"`
	CoolingType        string    `json:"cooling_type"`
	VantilationEnabled string    `json:"vantilation_enabled"`
	VantilationType    bool      `json:"vantilation_type"`
	Teamperature       string    `json:"teamperature"`
	Humidity           string    `json:"humidity"`
	LightLavel         string    `json:"light_lavel"`
	Co2Lavel           string    `json:"co2lavel"`
	CreatedAt          time.Time `json:"created_at"  gorm:"not null"`
	CreatedBy          string    `json:"created_by"  binding:"requird" gorm:"not null"`
	UpdatedAt          time.Time `json:"updated_at"`
	Upadatedby         string    `json:"updated_by"`

}
