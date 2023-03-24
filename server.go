package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/liraRaphael/golang-api-lib/util"
)

const (
	envDefault     = ".env"
	anddressFormat = ":%s"
	envFormat      = "%s.env"
)

type Server struct {
	Router *gin.Engine

	Profile string
	Port    string

	HealthCheckActive bool
	HealthCheckUrl    string
}

func Init() *Server {
	server := &Server{
		Router: gin.Default(),
	}

	var envs []string
	envs = append(envs, envDefault)

	server.loadEnvs(envDefault)
	if len(strings.TrimSpace(server.Profile)) > 0 {
		server.loadEnvs(envDefault, fmt.Sprintf(envFormat, server.Profile))
	}

	return server
}

func (s *Server) Run() {
	if s.HealthCheckActive {
		s.healthCheck()
	}

	s.Router.Run(fmt.Sprintf(anddressFormat, s.Port))
}

func (s *Server) loadEnvs(paths ...string) {
	godotenv.Load(paths...)

	s.Profile = strings.ToLower(util.GetEnvValue("application.profile"))
	s.Port = util.GetEnvValueOrDefault("application.port", "8080")

	s.HealthCheckActive = util.StringToBool(util.GetEnvValueOrDefault("application.healthcheck.active", "true"))
	s.HealthCheckUrl = util.GetEnvValueOrDefault("application.healthcheck.url", "/health/check")
}

func (s *Server) healthCheck() {
	s.Router.GET(s.HealthCheckUrl, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
