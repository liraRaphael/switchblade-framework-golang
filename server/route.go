package server

type BaseRequest struct {
	Body    interface{}
	Headers interface{}
}

type Route struct {
	Server *Server

	HttpSuccess int

	Endpoint string
	Method   string

	Request struct {
		BaseRequest

		Path    interface{}
		Queries interface{}
	}

	Response BaseRequest

	EnableValidator struct {
		RequestBody    bool
		RequestHeaders bool
		Path           bool
		Queries        bool
	}

	Documentation struct {
		IsEnable bool

		Summary     string
		Description string
		OperationId string
	}
}

func (s *Server) InitRoute() *Route {
	// if s.healthCheckActive {
	// 	s.healthCheck()
	// }

	return &Route{Server: s}
}

// ToDo: Criar interface para os builder
