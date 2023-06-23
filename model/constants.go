package model

import "time"

// Loglavels
var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelDebug   = "debug"
	LogLevelError   = "error"
	LogLevelWarning = "warn"
)

// PAckageLevel
var (
	PgressPackageLevel = "pgress"
	SarverPackageLevel = "server"
	ApiPackageLevel    = "api"
	MainPackageLevel   = "main"
	UtilPackageLevel   = "util"
)

// Function
var (
	// Server
	NewServer = "newServer"

	// Store
	NewStore = "newStore"

	// Actuator methods
	GetActuators        = "GetActuators"
	GetActuatorByFilter = "GetActuatorByFilter"
	GetActuator         = "GetActuator"
	CreatActuator       = "CreateActuator"
	UpdateActuator      = "UpdateActuator"
	DeleteActuator      = "DeleteActuator"

	//  House Method
	GetHouses        = "GetHouses"
	GEtHouseByFilter = "GetHouseByFilter"
	GetHouse         = "GetHouse"
	CreatHouse       = "CreatHouse"
	UpdateHouse      = "UpdateHouse"
	DeletFloor       = "DeleteHouse"

	// Floor Method
	GetFloors        = "GetFloors"
	GetFloorByFilter = "GetFloorByFilter"
	GetFloor         = "GetFloor"
	CreateFloor      = "CreateFloor"
	UpdateFloor      = "UpdateFloor"
	DeleteFloor      = "DeleteFloor"

	// Organization Methods
	GetOrganizations        = "GetOrfanizations"
	GetOrganizationByFilter = "GetOrganizationByFilter"
	GetOrganization         = "GetOrganization"
	CreateOrganization      = "CreateOrganization"
	UpdateORganization      = "UpdateOrganization"
	DeleteOrganization      = "DeleteOrganization"

	// Package Methods

	GetPackages        = "GetPackages"
	GetPackageByFilter = "GetPackageByFilter"
	GetPackage         = "GetPackage"
	CreatePackage      = "CreatePackage"
	UpdatedPackage     = "UpdatedPackage"
	DeletePackage      = "DeletePackage"

	// PoinRate Methods

	GetPointRates        = "GetPointRates"
	GetPointRateByFilter = "GetPointRateByFilter"
	GetPointRate         = "GetPointRate"
	CreatePoinRate       = "CreatePoinRate"
	UpdatedPointRate     = "UpdatedPointRate"
	DeletePointRate      = "DeletePointRate"

	// Room Methods

	GetRooms         = "GetRooms"
	GetRoomeByFilter = "GetRoomeByFilter"
	GetRoom          = "GetRoom"
	CreatedRoom      = "CreatedRoom"
	UpdatedRoom      = "UpdatedRoom"
	DeletedRoom      = "DeletedRoom"

	// Sensor Methods
	GetSensors         = "GetSensors"
	GetSensorsBYFilter = "GetSensorsByFilter"
	GetSensor          = "GetSensor"
	CreateSensor       = "CreateSensor"
	UpdatedSensor      = "UpdatedSenSor"
	DeleteSensor       = "DeleteSensor"

	// SensorReading Methods
	GetSensorReadings         = "GetSensorsReadings"
	GetSensorsReadingBYFilter = "GetSensorsReadingByFilter"
	GetSensorReading          = "GetSensorReading"
	CreatedSensorReading      = "CreateSensorReading"
	UpdatedSensorReading      = "UpdatedSenSorReading"
	DeleteSensorReading       = "DeleteSensorReading"

	// User Methods
	GetUsers        = "GetUsers"
	GetUserByFilter = "GetUserByFilter"
	GetUser         = "GetUser"
	CreatedUser     = "CreatedUser"
	UpdatedUser     = "UpdatedUser"
	DeletedUser     = "DeletedUser"
	SingUp			= "SingUp"
	SingIn			= "SingIn"

	Init			="init"
	Log				="Log"
	AuthNiddleWare ="AuthMiddleWare"
	AuthNiddleWareComplete ="AuthMiddleWareComplete" 
	SateLimitAndPage="SetLimitAndPage"
	SetDateRangeFilter="SetDateRangeFilter"

	

	// User Type
	HomeAutomationOwner ="HomeAutomationOwner"
	SuperAdminUser="SuperAdmin"
	AdminUser ="Admin"
	NormalUser ="User"
	
)
var(
	TokenExpiration=time.Hour*24

)
var Secretkey =[]byte("HomeAutomation-Secret-key")

// General
var (
	Value="value"
	Email="email"
	Password="password"
	UserId="userid"
	Expire ="exp"

	Autorization = "X-token"
	DSN ="host=localhost user=iot dbname=homeautomation port=5432 sslmode=disable"
	DataPerPage="limit"
	PageNumber="page"
	Startdate ="start_date"
	EndDate ="end_date"
	TimeLayout = "2023-08-01 12:05:09.000 -0700"
)