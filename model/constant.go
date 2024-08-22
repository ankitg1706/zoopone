package model

import "time"

var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelError   = "error"
	LegLevelDebug   = "debug"
	LogLevelWarning = "warn"
)

var (
	ApiPackage        = "api"
	StorePackage      = "store"
	ControllerPackage = "controller"
	ModelPackage      = "model"
	UtilPackage       = "util"
	MainPackage       = "main"
)

var (
	TokenExpiration = time.Hour * 24
)

var SecretKey = []byte("managment-secreat-key")

var (
	Controller = "controller"
	Store      = "store"
	Api        = "api"
	Main       = "main"
)

var (
	NewServer = "new-server"
	NewStore  = "new-store"

	CreateUser      = "create-user"
	GetUser         = "get-user"
	SignUP          = "sign-up"
	SignIn          = "sign-in"
	GetUsers        = "get-users"
	GetUserByFilter = "get-user-by-filter"
	UpdateUser      = "update-user"
	DeleteUser      = "delete-user"

	AuthMiddleware         = "AuthMiddleware"
	AuthMiddlewareComplete = "AuthMiddlewareComplete"
	SetLimitAndPage        = "setLimitAndPag e"
	SetDateRangeFilter     = "setDateRangeFilter"


)

// General
var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "userID"
	Expire   = "exp"

	Authorization = "X-Token"

	DNS = "host=localhost user=ankit password=password dbname=manage port=5432 sslmode=disable"

	DataPerPage = "limit"
	PageNumber  = "page"
	StartDate   = "start_date"
	EndDate     = "end_date"
	TimeLayout  = "2006-01-02 15:04:05.000 -0700"
)

// user type
var (
	HomeAutomationOwner = "HomeAutomationOwner"
	SuperAdminUser      = "superAdmin"
	AdminUser           = "Admin"
	NormalUser          = "User"
	UserTypes           = []string{"HomeAutomationOwner", "superAdmin", "Admin", "User"}
)
