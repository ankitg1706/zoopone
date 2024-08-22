package controller

import (

	"github.com/ankitg1706/zoopone/model"
	"github.com/ankitg1706/zoopone/store"
	"github.com/ankitg1706/zoopone/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	PostgressDb store.SoteOperations
}

func (s *Server) NewServer(pgstore store.Postgress) {
	util.SetLogger()
	util.Logger.Infof("Logger setup Done....\n")
	s.PostgressDb = &pgstore
	err := s.PostgressDb.NewStore()
	if err != nil {
		util.Logger.Errorf("error while creating store connection, err = %v\n", err)
		util.Log(model.LogLevelError, model.Controller, "NewServer", "error while creating store connection", err)
	} else {
		util.Logger.Infof("Connected with store\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected with store", nil)
	}
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)

	//User controllers
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByFilter(ctx *gin.Context)
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}
