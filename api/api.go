package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shivam904455/Home-Automation/server"
	"github.com/shivam904455/Home-Automation/store/pgress"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type APIRoutes struct {
	Server server.ServerOperation
}

func (api APIRoutes) StartApp(router *gin.Engine, server server.Server) {

	//Setup Server
	fmt.Println("Creating new server .....")
	api.Server = &server
	api.Server.NewServer(pgress.PgressStore{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user routs
	api.ActuatorRouts(router)
	api.FloorRouts(router)
	api.HouseRouts(router)
	api.OrganizationRouts(router)
	api.PackageRouts(router)
	api.PointRateRouts(router)
	api.RoomRouts(router)
	api.SensorRouts(router)
	api.SensorReadingRouts(router)
	api.UserRouts(router)

}
