package server

import "strings"

func (s *Server) NewRoute() *Route {
	return &Route{Server: s}
}

func (r *Route) DefineDefaultStatusCode() int {
	if r.Response.Body == nil {
		return 204
	}

	if strings.EqualFold(r.Method, "POST") {
		return 201
	}

	return 200
}
