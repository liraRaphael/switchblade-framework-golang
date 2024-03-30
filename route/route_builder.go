package route

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

	BuildRoute() *Route
}

func (r *Route) AddRoute(endpoint, method string) *Route {

	return r
}

func (r *Route) RequestBody(body interface{}) *Route {

	return r
}

func (r *Route) RequestHeaders(headers interface{}) *Route {

	return r
}

func (r *Route) Path(path interface{}) *Route {

	return r
}

func (r *Route) Queries(queries interface{}) *Route {

	return r
}

func (r *Route) ResponseBody(body interface{}) *Route {

	return r
}

func (r *Route) ResponseHeaders(headers interface{}) *Route {

	return r
}

func (r *Route) EnableValidatorRequestBody() *Route {

	return r
}

func (r *Route) EnableValidatorRequestHeaders() *Route {

	return r
}

func (r *Route) EnableValidatorPath() *Route {

	return r
}

func (r *Route) EnableValidatorQueries() *Route {

	return r
}

func (r *Route) DisableValidatorRequestBody() *Route {

	return r
}

func (r *Route) DisableValidatorRequestHeaders() *Route {

	return r
}

func (r *Route) DisableValidatorPath() *Route {

	return r
}

func (r *Route) DisableValidatorQueries() *Route {

	return r
}

func (r *Route) EnableDocumentation(summary, description, operationId string) *Route {

	return r
}

func (r *Route) DisableDocumentation() *Route {

	return r
}

func (r *Route) BuildRoute() *Route {

	return r
}
