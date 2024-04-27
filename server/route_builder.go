package server

import (
	"maps"

	"github.com/gofiber/fiber/v2"
)

type RouteBuilder[BReq, BResp, HReq, HResp, P, Q any] interface {
	AddRoute(endpoint, method string) *Route[BReq, BResp, HReq, HResp, P, Q]

	EnableValidatorRequestBody() *Route[BReq, BResp, HReq, HResp, P, Q]
	EnableValidatorRequestHeaders() *Route[BReq, BResp, HReq, HResp, P, Q]
	EnableValidatorPath() *Route[BReq, BResp, HReq, HResp, P, Q]
	EnableValidatorQueries() *Route[BReq, BResp, HReq, HResp, P, Q]
	DisableValidatorRequestBody() *Route[BReq, BResp, HReq, HResp, P, Q]
	DisableValidatorRequestHeaders() *Route[BReq, BResp, HReq, HResp, P, Q]
	DisableValidatorPath() *Route[BReq, BResp, HReq, HResp, P, Q]
	DisableValidatorQueries() *Route[BReq, BResp, HReq, HResp, P, Q]

	EnableDocumentation(summary, description, operationId string) *Route[BReq, BResp, HReq, HResp, P, Q]
	DisableDocumentation() *Route[BReq, BResp, HReq, HResp, P, Q]

	ExceptionHandleValidation(err error, callback func(report any)) *Route[BReq, BResp, HReq, HResp, P, Q]

	ListenRoute() *Route[BReq, BResp, HReq, HResp, P, Q]
	ListenRawRoute(callback func(c *fiber.Ctx) error) *Route[BReq, BResp, HReq, HResp, P, Q]
	AndServer() *Server
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) AddRoute(endpoint, method string) *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Endpoint = endpoint
	r.Method = method

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) EnableValidatorRequestBody() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.RequestBody.Enable = true

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) EnableValidatorRequestHeaders() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.RequestHeaders.Enable = true

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) EnableValidatorPath() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.Path.Enable = true

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) EnableValidatorQueries() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.Queries.Enable = true

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DisableValidatorRequestBody() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.RequestBody.Enable = false

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DisableValidatorRequestHeaders() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.RequestHeaders.Enable = false

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DisableValidatorPath() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.Path.Enable = false

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DisableValidatorQueries() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Validator.Queries.Enable = false

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) EnableDocumentation(summary, description, operationId string) *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Documentation.Enable = true
	r.Documentation.Summary = summary
	r.Documentation.Description = description
	r.Documentation.OperationId = operationId

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DisableDocumentation() *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.Documentation.Enable = false
	r.Documentation.Summary = ""
	r.Documentation.Description = ""
	r.Documentation.OperationId = ""

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) AddExceptionHandle(err error, callback func(report error) (RestResponse[any, any], error)) *Route[BReq, BResp, HReq, HResp, P, Q] {
	r.ExceptionHandler[err] = callback

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) AddExceptionsHandle(handles map[error]func(report error) (RestResponse[any, any], error)) *Route[BReq, BResp, HReq, HResp, P, Q] {
	maps.Copy(handles, r.ExceptionHandler)

	return r
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) ListenRoute(callback func(request RestRequest[BReq, HReq, P, Q]) (RestResponse[BResp, HResp], error)) {
	r.Server.ctx.Add(r.Method, r.Endpoint, r.DefaultCallbackFiber)
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) ListenRawRoute(callback func(c *fiber.Ctx) error) {
	r.Server.ctx.Add(r.Method, r.Endpoint, callback)
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) AndServer() *Server {
	return r.Server
}
