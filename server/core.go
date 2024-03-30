package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/liraRaphael/golang-api-lib/environment"
	"github.com/liraRaphael/golang-api-lib/util"
)

const (
	envDefault     = ".env"
	envFormat      = "%s.env"
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
	s.InitEnvs()

	s.router.Listen(fmt.Sprintf(anddressFormat, s.port))
}

func (s *Server) loadEnvs(paths ...string) {
	godotenv.Load(paths...)
}

func (s *Server) loadCoreEnvs() {
	s.profile = strings.ToLower(environment.GetEnvValueOrDefault("application.profile", ""))
	s.port = environment.GetEnvValueOrDefault("application.port", "8080")

	s.healthCheckActive = util.StringToBool(environment.GetEnvValueOrDefault("application.healthcheck.active", "true"))
	s.healthCheckUrl = environment.GetEnvValueOrDefault("application.healthcheck.url", "/health/check")
}

func (s *Server) InitEnvs() {
	s.loadEnvs(envDefault)

	if len(strings.TrimSpace(s.profile)) > 0 {
		s.loadEnvs(envDefault, fmt.Sprintf(envFormat, s.profile))
	}

	s.loadCoreEnvs()
}
