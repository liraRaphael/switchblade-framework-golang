package server

import (
	"maps"

	"github.com/gofiber/fiber/v2"
)

type RouteBuilder[BReq, BResp, HReq, HResp, P, Q any] interface {
	AddRoute(endpoint, method string) *Route

	RequestBody(body BReq) *Route
	RequestHeaders(headers HReq) *Route
	Path(path P) *Route
	Queries(queries Q) *Route

	ResponseBody(body BResp) *Route
	ResponseHeaders(headers HResp) *Route

	EnableValidatorRequestBody() *Route
	EnableValidatorRequestHeaders() *Route
	EnableValidatorPath() *Route
	EnableValidatorQueries() *Route
	DisableValidatorRequestBody() *Route
	DisableValidatorRequestHeaders() *Route
	DisableValidatorPath() *Route
	DisableValidatorQueries() *Route

	EnableDocumentation(summary, description, operationId string) *Route
	DisableDocumentation() *Route

	ExceptionHandleValidation(err error, callback func(report any)) *Route

	ListenRoute() *Route
	AndServer() *Server
}

func (r *Route) AddRoute(endpoint, method string) *Route {
	r.Endpoint = endpoint
	r.Method = method

	r.Validator.RequestBody = true
	r.Validator.RequestHeaders = true
	r.Validator.Path = true
	r.Validator.Queries = true

	r.Documentation.IsEnable = false
	r.Documentation.Summary = ""
	r.Documentation.Description = ""
	r.Documentation.OperationId = ""

	return r
}

func (r *Route) RequestBody(body any) *Route {
	r.Request.Body = body

	return r
}

func (r *Route) RequestHeaders(headers any) *Route {
	r.Request.Headers = headers

	return r
}

func (r *Route) Path(path any) *Route {
	r.Request.Path = path

	return r
}

func (r *Route) Queries(queries any) *Route {
	r.Request.Queries = queries

	return r
}

func (r *Route) ResponseBody(body any) *Route {
	r.Response.Body = body

	return r
}

func (r *Route) ResponseHeaders(headers any) *Route {
	r.Response.Headers = headers

	return r
}

func (r *Route) EnableValidatorRequestBody() *Route {
	r.Validator.RequestBody = true

	return r
}

func (r *Route) EnableValidatorRequestHeaders() *Route {
	r.Validator.RequestHeaders = true

	return r
}

func (r *Route) EnableValidatorPath() *Route {
	r.Validator.Path = true

	return r
}

func (r *Route) EnableValidatorQueries() *Route {
	r.Validator.Queries = true

	return r
}

func (r *Route) DisableValidatorRequestBody() *Route {
	r.Validator.RequestBody = false

	return r
}

func (r *Route) DisableValidatorRequestHeaders() *Route {
	r.Validator.RequestHeaders = false

	return r
}

func (r *Route) DisableValidatorPath() *Route {
	r.Validator.Path = false

	return r
}

func (r *Route) DisableValidatorQueries() *Route {
	r.Validator.Queries = false

	return r
}

func (r *Route) EnableDocumentation(summary, description, operationId string) *Route {
	r.Documentation.IsEnable = true
	r.Documentation.Summary = summary
	r.Documentation.Description = description
	r.Documentation.OperationId = operationId

	return r
}

func (r *Route) DisableDocumentation() *Route {
	r.Documentation.IsEnable = false
	r.Documentation.Summary = ""
	r.Documentation.Description = ""
	r.Documentation.OperationId = ""

	return r
}

func (r *Route) AddExceptionHandle(err error, callback func(report any)) *Route {
	r.ExceptionHandler[err] = callback

	return r
}

func (r *Route) AddExceptionsHandle(handles map[error]func(report any)) *Route {
	maps.Copy(handles, r.ExceptionHandler)

	return r
}

func (r *Route) ListenRoute() *Route {
	r.Server.ctx.Add(r.Method, r.Endpoint, func(c *fiber.Ctx) error {
		return nil
	})

	return r.Server.NewRoute()
}

func (r *Route) AndServer() *Server {
	return r.Server
}
