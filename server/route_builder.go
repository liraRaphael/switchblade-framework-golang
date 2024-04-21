package server

import "maps"

type RouteBuilder interface {
	AddRoute(endpoint, method string) *Route

	RequestBody(body interface{}) *Route
	RequestHeaders(headers interface{}) *Route
	Path(path interface{}) *Route
	Queries(queries interface{}) *Route

	ResponseBody(body interface{}) *Route
	ResponseHeaders(headers interface{}) *Route

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

	ExceptionHandleValidation(err interface{}, callback func(report interface{})) *Route

	ListenRoute() *Route
	AndServer() *Server
}

func (r *Route) AddRoute(endpoint, method string) *Route {
	r.Endpoint = endpoint
	r.Method = method

	r.EnableValidator.RequestBody = true
	r.EnableValidator.RequestHeaders = true
	r.EnableValidator.Path = true
	r.EnableValidator.Queries = true

	r.Documentation.IsEnable = false
	r.Documentation.Summary = ""
	r.Documentation.Description = ""
	r.Documentation.OperationId = ""

	return r
}

func (r *Route) RequestBody(body interface{}) *Route {
	r.Request.Body = body

	return r
}

func (r *Route) RequestHeaders(headers interface{}) *Route {
	r.Request.Headers = headers

	return r
}

func (r *Route) Path(path interface{}) *Route {
	r.Request.Path = path

	return r
}

func (r *Route) Queries(queries interface{}) *Route {
	r.Request.Queries = queries

	return r
}

func (r *Route) ResponseBody(body interface{}) *Route {
	r.Response.Body = body

	return r
}

func (r *Route) ResponseHeaders(headers interface{}) *Route {
	r.Response.Headers = headers

	return r
}

func (r *Route) EnableValidatorRequestBody() *Route {
	r.EnableValidator.RequestBody = true

	return r
}

func (r *Route) EnableValidatorRequestHeaders() *Route {
	r.EnableValidator.RequestHeaders = true

	return r
}

func (r *Route) EnableValidatorPath() *Route {
	r.EnableValidator.Path = true

	return r
}

func (r *Route) EnableValidatorQueries() *Route {
	r.EnableValidator.Queries = true

	return r
}

func (r *Route) DisableValidatorRequestBody() *Route {
	r.EnableValidator.RequestBody = false

	return r
}

func (r *Route) DisableValidatorRequestHeaders() *Route {
	r.EnableValidator.RequestHeaders = false

	return r
}

func (r *Route) DisableValidatorPath() *Route {
	r.EnableValidator.Path = false

	return r
}

func (r *Route) DisableValidatorQueries() *Route {
	r.EnableValidator.Queries = false

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

func (r *Route) AddExceptionHandle(err interface{}, callback func(report interface{})) *Route {
	r.ExceptionHandler[err] = callback

	return r
}

func (r *Route) AddExceptionsHandle(handles map[interface{}]func(report interface{})) *Route {
	maps.Copy(handles, r.ExceptionHandler)

	return r
}

func (r *Route) ListenRoute() *Route {
	//ToDo: Colocar as rotas para serem ouvidas pelo fiber
	return r.Server.NewRoute()
}

func (r *Route) AndServer() *Server {
	return r.Server
}
