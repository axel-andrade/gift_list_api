package server

import (
	cors_mdw "go_gift_list_api/src/infra/http/middlewares/cors"
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
	s := gin.New()
	s.Use(cors_mdw.CORSMiddleware())

	return Server{
		port:   os.Getenv("PORT"),
		server: s,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	log.Fatal(router.Run(":" + s.port))
}
