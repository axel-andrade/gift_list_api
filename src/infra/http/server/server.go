package server

import (
	"go_gift_list_api/src/infra/http/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   os.Getenv("PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	log.Fatal(router.Run(":" + s.port))
}
