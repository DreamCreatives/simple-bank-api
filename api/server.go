package api

import (
	db "github.com/DreamCreatives/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := engine.RegisterValidation("currency", validCurrency, false); err != nil {
			log.Fatalf("cannot create validator. Error: %v", err)
		}
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}
