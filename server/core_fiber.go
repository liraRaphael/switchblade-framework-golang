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
	ctx *fiber.App

	profile string
	port    string

	healthCheckActive bool
	healthCheckUrl    string

	defaultHandleUnknowException func(report any)

	defaultExceptionHandler map[error]func(report any)
}

func Init() *Server {
	server := &Server{
		ctx: fiber.New(),
	}

	return server
}

func (s *Server) Run() {
	s.InitEnvs()

	s.ctx.Listen(fmt.Sprintf(anddressFormat, s.port))
}

func (s *Server) loadEnvs(paths ...string) {
	godotenv.Load(paths...)
}

func (s *Server) SetDefaultHandleUnknowException(callback func(report interface{})) {
	s.defaultHandleUnknowException = callback
}

func (s *Server) loadCoreEnvs() {
	s.profile = strings.ToLower(environment.GetEnvValueOrDefault("ENVRIONMENT", ""))
	s.port = environment.GetEnvValueOrDefault("PORT", "8080")

	s.healthCheckActive = util.StringToBool(environment.GetEnvValueOrDefault("HEALTHCHECK_ACTIVE", "true"))
	s.healthCheckUrl = environment.GetEnvValueOrDefault("HEALTHCHECK_URL", "/health/check")
}

func (s *Server) AddDefaultExceptionHandle(err error, callback func(report any)) *Server {
	s.defaultExceptionHandler[err] = callback

	return s
}

func (s *Server) GetContext() *fiber.App {
	return s.ctx
}

func (s *Server) InitEnvs() {
	s.loadEnvs(envDefault)

	if len(strings.TrimSpace(s.profile)) > 0 {
		s.loadEnvs(envDefault, fmt.Sprintf(envFormat, s.profile))
	}

	s.loadCoreEnvs()
}
