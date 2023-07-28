package api

import "github.com/gin-gonic/gin"

func (api APIRoutes) OrganizationRouts(router *gin.Engine) {
	// Define routes
	organizationapi := router.Group("/organization")
	{
		organizationapi.GET("/all", api.GetAllOrganizations)
		organizationapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetOrganizationsByFilter)
		organizationapi.GET("/:id", api.GetOrganization)
		organizationapi.POST("/create", api.CreateOrganization)
		organizationapi.PUT("/update/:id", api.UpdateOrganization)
		organizationapi.DELETE("/delete/:id", api.DeleteOrganization)
	}

}

// Handler to get all organizations
// @router /organization/all [get]
// @summary Get all organizations
// @tags organizations
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Organization
func (api APIRoutes) GetAllOrganizations(c *gin.Context) {
	api.Server.GetOrganizations(c)
}

// Handler to get all organizations based on filter
// @router /organization/filter [get]
// @summary Get all organizations based on given filters
// @tags organizations
// @produce json
// @param id query string false "ID"
// @param package_id query string true "Package ID"
// @param name query string true "Name"
// @param description query string false "Description"
// @param package_type query string true "Package Type"
// @param email query string true "Email"
// @param password query string true "Password"
// @param available_points query integer false "Available Points"
// @param available_number_of_rooms query integer false "Available Number of Rooms"
// @param available_number_of_floors query integer false "Available Number of Floors"
// @param available_number_of_houses query integer false "Available Number of Houses"
// @param available_number_of_users query integer false "Available Number of Users"
// @param available_number_of_actuators query integer false "Available Number of Actuators"
// @param available_number_of_sensors query integer false "Available Number of Sensors"
// @param created_by query string true "Created By"
// @param created_at query string true "Created At"
// @param updated_at query string false "Updated At"
// @param updated_by query string false "Updated By"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Organization
// @Security ApiKeyAuth
func (api APIRoutes) GetOrganizationsByFilter(c *gin.Context) {
	api.Server.GetOrganizationsByFilter(c)
}

// Handler to get a organization by ID
// @router /organization/{id} [get]
// @summary Get a organization by ID
// @tags organizations
// @produce json
// @param id path string true "Organization ID"
// @success 200 {object} model.Organization
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetOrganization(c *gin.Context) {
	api.Server.GetOrganization(c)
}

// Handler to create a organization
// @router /organization/create [post]
// @summary Create a organization
// @tags organizations
// @accept json
// @produce json
// @param organization body model.Organization true "Organization object"
// @success 201 {object} model.Organization
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateOrganization(c *gin.Context) {
	api.Server.CreateOrganization(c)
}

// Handler to update a organization
// @router /organization/update/{id} [put]
// @summary Update a organization
// @tags organizations
// @accept json
// @produce json
// @param id path string true "Organization ID"
// @param organization body model.Organization true "Organization object"
// @success 200 {object} model.Organization
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateOrganization(c *gin.Context) {
	api.Server.UpdateOrganization(c)
}

// Handler to delete a organization
// @router  /organization/delete/{id} [delete]
// @summary Delete a organization
// @tags organizations
// @param id path string true "Organization ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteOrganization(c *gin.Context) {
	api.Server.DeleteOrganization(c)
}
