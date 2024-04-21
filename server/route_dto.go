package server

type BaseRequest struct {
	Body    interface{}
	Headers interface{}
}

type Route struct {
	Server *Server

	Endpoint string
	Method   string

	Callback func(request RestRequest[any, any, any, any]) (RestResponse[any, any], error)

	ExceptionHandler map[interface{}]func(report interface{})

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
