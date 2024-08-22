package api

import (
	"github.com/ankitg1706/zoopone/model"
	"github.com/ankitg1706/zoopone/util"
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) UserRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("user")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateUser)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetUser)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetUsers)
		group.POST("/signup", api.SignUp)
		group.POST("/signin", api.SignIn)
		group.GET("/filter", api.GetUsersByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateUser)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteUser)
	}

}

// Handler to create a user
// @router /user/create [post]
// @summary Create a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @success 201 {object} model.User
func (api APIRoutes) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get a user by ID
// @router /user/{id} [get]
// @summary Get a user by ID
// @tags users
// @produce json
// @param id path string true "User ID"
// @success 200 {object} model.User
// @Security ApiKeyAuth
func (api APIRoutes) GetUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "fetching  user", nil)
	api.Server.GetUser(ctx)
}

// Handler to get all users
// @router /user/all [get]
// @summary Get all users
// @tags users
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api APIRoutes) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "fetching users", nil)
	api.Server.GetUsers(ctx)
}

// Handler to get all users based on filter
// @router /user/filter [get]
// @summary Get all users based on given filters
// @tags users
// @produce json
// @param email query string false "email"
// @param id query string false "id"
// @param password query string false "password"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param first_name query string false "first_name"
// @param middle_name query string false "middle_name"
// @param last_name query string false "last_name"
// @param lane query string false "lane"
// @param village query string false "village"
// @param city query string false "city"
// @param district query string false "district"
// @param pincode query int false "pincode"
// @param state query string false "state"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api APIRoutes) GetUsersByFilter(c *gin.Context) {
	api.Server.GetUserByFilter(c)
}

// Handler to update a user
// @router /user/update/{id} [put]
// @summary Update a user
// @tags users
// @accept json
// @produce json
// @param id path string true "User ID"
// @param user body model.User true "User object"
// @success 200 {object} model.User
// @Security ApiKeyAuth
func (api APIRoutes) UpdateUser(c *gin.Context) {
	api.Server.UpdateUser(c)
}

// Handler to delete a user
// @router  /user/delete/{id} [delete]
// @summary Delete a user
// @tags users
// @param id path string true "User ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteUser(c *gin.Context) {
	api.Server.DeleteUser(c)
}

// Handler to SignUp a user
// @router /user/signup [post]
// @summary SignUp a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @Success 200 {string} string "Successful SignUp"
// @failure 400 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api APIRoutes) SignUp(c *gin.Context) {
	api.Server.SignUp(c)
}

// Handler to signIn a user by email and password
// @router /user/signin [post]
// @summary SighIn user
// @tags users
// @produce json
// @param user body model.UserSignIn true "User object"
// @Success 200 {string} string "Successful SignIn"
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api APIRoutes) SignIn(c *gin.Context) {
	api.Server.SignIn(c)
}
