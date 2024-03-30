package core

import (
	"github.com/liraRaphael/golang-api-lib/route"
)

func (s *Server) InitRoute() *route.Route {
	if s.healthCheckActive {
		s.healthCheck()
	}

	return &route.Route{Server: s}
}
