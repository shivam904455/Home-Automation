package model

import (
	"time"

	uuid "github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                    uuid.UUID `json:"id" gorm:"primarykey"`
	OrgId                 uuid.UUID `json:"org_id"`
	Name                  string    `json:"name" binding:"required" gorm:"unique;not null"`
	Publisher             string    `json:"publisher"`
	Email                 string    `json:"email" binding:"required" gorm:"unique;not null"`
	Password              string    `json:"passwoed" binding:"required" gorm:"not null"`
	PrimaryMobileNumber   int       `json:"primary_mobile_number" binding:"required" gorm:"not null"`
	SecondaryMobileNumber int       `json:"secondary_mobile_number"`
	LandLineNumber        int       `json:"land_line_number"`
	ActiveStatus          bool      `json:"active_status" binding:"required" gorm:"not null"`
	Type                  string    `json:"type" gorm:"not null"`
	AddressType           string    `json:"address_type"`
	HouseNo               int       `json:"house_no"`
	HouseName             string    `json:"house_name"`
	LaneNumber            int       `json:"lane_number"`
	LaneName              string    `json:"lane_name"`
	Landmark              string    `json:"landmark"`
	District              string    `json:"district" binding:"required" gorm:"mot null"`
	Post                  string    `json:"post"`
	City                  string    `json:"city"`
	VIllage               string    `json:"village"`
	State                 string    `json:"state" binding:"required" gorm:"not null"`
	Nation                string    `json:"nation" binding:"required" gorm:"not null"`
	CreatedBy             string    `json:"created_by" binding:"required" gorm:"not null"`
	CreatedAt             time.Time `json:"created_at" gorm:"not null"`
	UpdatedBy             string    `json:"updated_by"`
	UpdatedAt             time.Time `json:"updated_at"`
}
type UserSingnIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
