package route

import (
	"github.com/liraRaphael/golang-api-lib/core"
)

type BaseRequest struct {
	Body    interface{}
	Headers interface{}
}

type Route struct {
	Server *core.Server

	HttpSuccess int

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

// ToDo: Criar interface para os builder
