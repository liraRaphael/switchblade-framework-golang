package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/liraRaphael/golang-api-lib/util"
)

const (
	envDefault     = ".env"
	anddressFormat = ":%s"
	envFormat      = "%s.env"
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

	var envs []string
	envs = append(envs, envDefault)

	server.loadEnvs(envDefault)
	if len(strings.TrimSpace(server.profile)) > 0 {
		server.loadEnvs(envDefault, fmt.Sprintf(envFormat, server.profile))
	}

	return server
}

func (s *Server) Run() {
	if s.healthCheckActive {
		s.healthCheck()
	}

	s.router.Listen(fmt.Sprintf(anddressFormat, s.port))
}

func (s *Server) loadEnvs(paths ...string) {
	godotenv.Load(paths...)

	s.profile = strings.ToLower(util.GetEnvValue("application.profile"))
	s.port = util.GetEnvValueOrDefault("application.port", "8080")

	s.healthCheckActive = util.StringToBool(util.GetEnvValueOrDefault("application.healthcheck.active", "true"))
	s.healthCheckUrl = util.GetEnvValueOrDefault("application.healthcheck.url", "/health/check")
}

func (s *Server) healthCheck() {
	s.router.Get(s.healthCheckUrl, func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
