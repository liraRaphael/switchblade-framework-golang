package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/liraRaphael/golang-api-lib/environment"
	"github.com/liraRaphael/golang-api-lib/log"
	"github.com/liraRaphael/golang-api-lib/util"
)

const (
	envDefault     = ".env"
	envFormat      = "%s.env"
	anddressFormat = ":%s"
)

type Server struct {
	ctx *gin.Engine

	profile string
	port    string

	healthCheckActive bool
	healthCheckUrl    string

	defaultHandleUnknowException func(error)

	defaultExceptionHandler map[error]func(error)

	log log.Logging
}

func Init() *Server {
	server := &Server{
		ctx: gin.Default(),
	}

	return server
}

func (s *Server) Run() {
	s.InitEnvs()
	s.ctx.Use(gin.Recovery())

	s.ctx.Run()
}

func (s *Server) loadEnvs(paths ...string) {
	godotenv.Load(paths...)
}

func (s *Server) SetDefaultHandleUnknowException(callback func(error)) *Server {
	s.defaultHandleUnknowException = callback

	return s
}

func (s *Server) loadCoreEnvs() {
	s.profile = strings.ToLower(environment.GetEnvValueOrDefault("ENVRIONMENT", ""))
	s.port = environment.GetEnvValueOrDefault("PORT", "8080")

	s.healthCheckActive = util.StringToBool(environment.GetEnvValueOrDefault("HEALTHCHECK_ACTIVE", "true"))
	s.healthCheckUrl = environment.GetEnvValueOrDefault("HEALTHCHECK_URL", "/health/check")
}

func (s *Server) AddDefaultExceptionHandle(err error, callback func(error)) *Server {
	s.defaultExceptionHandler[err] = callback

	return s
}

func (s *Server) GetContext() *gin.Engine {
	return s.ctx
}

func (s *Server) InitEnvs() *Server {
	s.loadEnvs(envDefault)

	if len(strings.TrimSpace(s.profile)) > 0 {
		s.loadEnvs(envDefault, fmt.Sprintf(envFormat, s.profile))
	}

	s.loadCoreEnvs()

	return s
}

func (s *Server) SetLog(log log.Logging) *Server {
	s.log = log

	return s
}
