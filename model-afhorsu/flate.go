package model

import ( "time"
uuid "github.com/google/uuid"
)
type flate struct {
	ID                 uuid.UUID `json:"id" gorm:"primaryKey"`
	HouseID            uuid.UUID `json:"house_id" binding:"required" gorm:"not null"`
	Name               string    `json:"name" binding:"requird" gorm:"not null"`
	Description        string    `json:"descripiton"`
	Area               string    `json:"area"`
	Occupancy          string    `json :"occupancy"`
	HeatingEnabled     bool      `json :"heating_enabled"`
	HeatingType        string    `json :"heating_type"`
	CoolingEnabled     string    `json :"cooling_enabled"`
	CoolingType        string    `json :"cooling_type"`
	VantilationEnabled bool      `json :"vantilation_enabled"`
	VantilationType    string    `json :"vantilation_type"`
	Teamperature       string    `json:"teamperature"`
	Humidity           string    `json: "humidity"`
	LightLavel         string    `json :"light_lavel"`
	Co2Lavel           string    `json : "co2_lavel"`
	CreatedAt          time.Time `json:"created_at" gorm: "not null" `
	CreatedBy          string    `json:"creted_by" binding:"requird" gorm:"not null"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"not null"`
	UpdatedBy          string    `json:"updated_by" binding:"requird" gorm:"not null"`
}
