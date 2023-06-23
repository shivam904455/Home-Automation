package model

import (
	"time"
uuid "github.com/google/uuid"
)

type Organisation struct {
	ID                         uuid.UUID `json:"id" gorm:"primarykey "`
	PakcageId                  uuid.UUID `json:"pakcage_id" binding:"required" gorm:"not null"`
	Name                       string    `json:"name" binding:"required" gorm:"not null"`
	Description                string    `json:"description"`
	PackageType                string    `json:"package_type" binding:"required" gorm:"not null"`
	Email                      string    `json:"email" binding:"required" gorm:"not null"`
	Password                   string    `json:"password" binding:"required" gorm:"not null"`
	AvailablePoints            int       `json:"available_points"`
	AvailableNumberofRooms     int       `json:"available_numer_of_rooms"`
	AvailableNUmberoffloor     int       `json:"availabel_number_of_floor"`
	AvailableNumberofHouses    int       `json:"available_number_of_houses"`
	AvailableNumberofUsers     int       `json:"available_number_of_users"`
	AvailableNUmberofActuators int       `json:"available_number_of_actuators"`
	AvailableNumberofSensors   int       `json:"available_number_of_sensor"`
	CreatedBy                  string    `json:"Created_by" binding:"required" gorm:"not null"`
	CreatedAt                  time.Time `json:"created_at" gorm:"not null"`
	UpdatedBy                  string    `json:"updated_by" `
	UpdatedAt                  time.Time `json:"updated_at"`
}
 