package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	anddressFormat = ":%s"
)

type Server struct {
	router *fiber.App

	profile string
	port    string

	healthCheckActive bool
	healthCheckUrl    string
}

func Init() *Server {
	server := &Server{
		router: fiber.New(),
	}

	return server
}

func (s *Server) Run() {
	s.InitRoute()
	s.InitEnvs()

	s.router.Listen(fmt.Sprintf(anddressFormat, s.port))
}
