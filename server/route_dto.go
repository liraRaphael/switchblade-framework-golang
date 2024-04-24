package server

type BaseRequest[B, H any] struct {
	Body    B
	Headers H
}

type Response[B, H any] struct {
	BaseRequest[B, H]
}

type Request[B, H, P, Q any] struct {
	BaseRequest[B, H]

	Path    P
	Queries Q
}

type Validator struct {
	RequestBody    bool
	RequestHeaders bool
	Path           bool
	Queries        bool
}

type Documentation struct {
	IsEnable bool

	Summary     string
	Description string
	OperationId string
}

type Route struct {
	Server *Server

	Endpoint string
	Method   string

	ExceptionHandler map[error]func(report error)

	Request  Request[any, any, any, any]
	Response Response[any, any]

	Validator Validator

	Documentation Documentation
}
