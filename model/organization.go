package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                         uuid.UUID `json:"id" gorm:"primaryKey"`
	PackageID                  uuid.UUID `json:"package_id" binding:"required" gorm:"not null"`
	Name                       string    `json:"name" binding:"required" gorm:"not null"`
	Description                string    `json:"description"`
	PackageType                string    `json:"package_type" binding:"required" gorm:"not null"`
	Email                      string    `json:"email" binding:"required" gorm:"not null"`
	Password                   string    `json:"password" binding:"required" gorm:"not null" `
	AvailablePoints            int       `json:"available_points"`
	AvailableNumberOfRooms     int       `json:"available_number_of_rooms"`
	AvailableNumberOfFloors    int       `json:"available_number_of_floors"`
	AvailableNumberOfHouses    int       `json:"available_number_of_houses"`
	AvailableNumberOfUsers     int       `json:"available_number_of_users"`
	AvailableNumberOfActuators int       `json:"available_number_of_actuators"`
	AvailableNumberOfSensors   int       `json:"available_number_of_sensors"`
	CreatedBy                  string    `json:"created_by" binding:"required" gorm:"not null"`
	CreatedAt                  time.Time `json:"created_at"  gorm:"not null"`
	UpdatedAt                  time.Time `json:"updated_at"`
	UpdatedBy                  string    `json:"updated_by"`
}