package server

import (
	"go_gift_list_api/src/infra/http/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	s := gin.New()
	s.Use(cors.Default())

	return Server{
		port:   os.Getenv("PORT"),
		server: s,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.SetTrustedProxies(nil)

	log.Fatal(router.Run(":" + s.port))
}
