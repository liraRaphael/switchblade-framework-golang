package core

import "github.com/gofiber/fiber/v2"

func (s *Server) InitRoute() {
	if s.healthCheckActive {
		s.healthCheck()
	}
}

func (s *Server) healthCheck() {
	s.router.Get(s.healthCheckUrl, func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
