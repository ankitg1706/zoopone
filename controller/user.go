package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ankitg1706/zoopone/model"
	"github.com/ankitg1706/zoopone/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (server Server) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "creating new user", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgressDb.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUser(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "fetching user by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "invalid user ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := server.PostgressDb.GetUser(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUserByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUserByFilter, "fetching user by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	user, err := server.PostgressDb.GetUserByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUserByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, "fetching all user", nil)

	users, err := server.PostgressDb.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "error while fetching users record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, users)

}

func (server *Server) UpdateUser(c *gin.Context) error {

	var user model.User
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateUser,
		"unmarshaling user data", nil)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateUser,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	user.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateUser,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers,
		"successfully updated user record and setting response", user)
	c.JSON(http.StatusOK, user)
	return nil

}

func (server *Server) DeleteUser(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteUser,
		"reading user id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteUser,
			"missing user id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteUser(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteUser,
			"error while deleting user record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteUser,
		"successfully deleted user record ", nil)
	return nil

}

// Signup API handler
func (server *Server) SignUp(c *gin.Context) {
	var user model.User

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.SignUP,
		"unmarshaling user data", nil)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	err := server.PostgressDb.SignUp(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.SignUP,
			"error in saving user record", user)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to SignUp User"})
		return
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   user.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(), // Token expiration time
		// Additional data can be added here
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// SignIn API handler
func (server *Server) SignIn(c *gin.Context) {
	var user model.UserSignIn
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser,
			"error while unmarshaling payload", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user Data from payload"})
		return
	}

	userResp, err := server.PostgressDb.SingIn(user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.SignIn,
			"error in getting user data from pgress for emailId", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user Data for given user"})
		return
	}
	if userResp.Email != user.Email || userResp.Password != user.Password {
		util.Log(model.LogLevelInfo, model.ControllerPackage, model.SignIn,
			"user data not matched , database response", userResp)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user data"})
		return
	}

	// Create a new token
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   userResp.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(), // Token expiration time
		// Additional data can be added here
	})

	// Sign the newtoken with the secret key
	tokenString, err := newtoken.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
