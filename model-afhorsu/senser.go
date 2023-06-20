package model

import ( "time"
uuid "github.com/google/uuid"
)

type Senser struct {
	ID                 uuid.UUID `json: "id"gorm:"primarykey" `
	RoomsID            uuid.UUID `json: "rooms_id" binding:"requird" gorm:"not null"`
	HouseID            uuid.UUID `json:"House_id" binding:"requird" gorm:"not null"`
	FloorID            uuid.UUID `json:"floor_id" binding :"requird" gorm: "not null"`
	Name               string    `json:"name" binding: "requird" gorm: "not null"`
	Description        string    `json:"description"`
	Area               string    `json :"area"`
	WattConsumption    string    `json :"watt_consumption"`
	Occupancy          string    `json :"occupancy"`
	HeatingEnabled     bool      `json :"heating_enabled"`
	HeatingType        string    `json :"heating_type"`
	CoolingEnabled     bool      `json :"cooling_enabled"`
	cooling_type       string    `json :"cooling_type"`
	VantilationEnabled bool      `json :"vantilation_enabled"`
	VantilationType    string    `json :"vantilation_type"`
	Teamperature       string    `json :"teamperature"`
	Humidity           string    `json :"humidity"`
	LightLavel         string    `json :"light_lavel"`
	Co2Lavel           string    `json :"co2_lavel"`
	CreatedAt          time.Time `json :"created_at" gorm: "not null"`
	CreatedBy     string `json:"created_by"inding:"requird" gorm: "not null"`
}